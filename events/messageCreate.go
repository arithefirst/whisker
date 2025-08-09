package events

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/arithefirst/whisker/helpers"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5"
)

type MemberEntry struct {
	UserID               string
	GuildID              string
	XP                   uint64
	Level                uint
	LastMessageTimestamp time.Time
	LastMessage          string
}

type InvocationData struct {
	Timestamp  time.Time
	ReplyMsgID string
	ErrorMsgID string
	ChannelID  string
}

// map[parent message id] = invocation data
var previousInvocations = make(map[string]InvocationData)
var mutex sync.RWMutex
var preamble string = `
#set page(fill: black, height: auto, width: auto, margin: 10pt)
#set text(fill: white, size: 18pt)
#show math.equation: eq => [
  #text(fill: white, [ #eq ])
]

`

// TODO:
//  criteria for message xp (prevent spam)
// 	create a command for xp
//  leveling
//  guild leaderboard
//  level up announcement

func (h *Handler) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "meow!xp" {
		var xp uint64
		query := "SELECT xp FROM bot.guild_members WHERE user_id = $1 AND guild_id = $2"

		err := h.DB.QueryRow(context.Background(), query, m.Author.ID, m.GuildID).Scan(&xp)
		if err != nil && err != pgx.ErrNoRows {
			log.Printf("Error getting user xp: %v", err)
			return
		}

		s.ChannelMessageSendEmbed(m.ChannelID, helpers.CreateEmbed().
			SetDescription(fmt.Sprintf("you got %v xp", xp)).
			MessageEmbed,
		)
	}

	if strings.HasPrefix(m.Content, "meow!ty") {
		invocation := InvocationData{
			Timestamp: time.Now(),
			ChannelID: m.ChannelID,
		}

		code := preamble + strings.TrimSpace(m.Content[7:])
		file, err := helpers.RenderTypst(code)

		if err != nil {
			errorMsg, _ := s.ChannelMessageSendEmbed(m.ChannelID, helpers.ErrorEmbed("rendering Typst", err)[0])

			mutex.Lock()
			defer mutex.Unlock()

			invocation.ErrorMsgID = errorMsg.ID
			previousInvocations[m.ID] = invocation

			return
		}

		msg, err := s.ChannelFileSend(m.ChannelID, file.Name, file.Reader)
		if err != nil {
			s.ChannelMessageSendEmbed(m.ChannelID, helpers.ErrorEmbed("sending attachment", err)[0])
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		invocation.ReplyMsgID = msg.ID
		previousInvocations[m.ID] = invocation

		return
	}

	h.handleXP(m)
}

func countUniqueChars(s string) int {
	charSet := make(map[rune]bool)
	for _, char := range s {
		charSet[char] = true
	}
	return len(charSet)
}

func (h *Handler) handleXP(m *discordgo.MessageCreate) {
	var lastMessageTimestamp time.Time

	query := "SELECT last_message_timestamp FROM bot.guild_members WHERE user_id = $1 AND guild_id = $2"
	err := h.DB.QueryRow(context.Background(), query, m.Author.ID, m.GuildID).Scan(&lastMessageTimestamp)

	// user not in the db yet, so they can get XP
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Error checking user timestamp: %v", err)
		return
	}

	// 60s cooldown
	if time.Since(lastMessageTimestamp) < 5*time.Second {
		log.Printf("%s is on cooldown", m.Author.DisplayName())
		return
	}

	uniqueChars := float64(countUniqueChars(m.Content))

	minXP := uniqueChars / 7
	maxXP := uniqueChars / 4

	if minXP > maxXP {
		minXP = maxXP
	}

	xpToAdd := int(minXP + rand.Float64()*(maxXP-minXP))

	if xpToAdd == 0 {
		xpToAdd = 1
	}

	if xpToAdd > 35 {
		xpToAdd = 35
	}

	upsertQuery := `
        INSERT INTO bot.guild_members (user_id, guild_id, xp, last_message_timestamp)
        VALUES ($1, $2, $3, NOW())
        ON CONFLICT (user_id, guild_id) DO UPDATE
        SET xp = guild_members.xp + $3,
            last_message_timestamp = NOW();
    `

	_, err = h.DB.Exec(context.Background(), upsertQuery, m.Author.ID, m.GuildID, xpToAdd)
	if err != nil {
		log.Printf("Failed to update XP for user %s: %v", m.Author.ID, err)
	} else {
		log.Printf("Awarded %d XP to user %s in guild %s", xpToAdd, m.Author.ID, m.GuildID)
	}
}

// clean up invocations after 5 minutes
func init() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			mutex.Lock()
			for msgID, data := range previousInvocations {
				if time.Since(data.Timestamp) > 5*time.Minute {
					delete(previousInvocations, msgID)
					log.Printf("cleaned up expired invocation for channel: %s", msgID)
				}
			}

			mutex.Unlock()
		}
	}()
}
