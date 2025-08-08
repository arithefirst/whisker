package events

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/arithefirst/whisker/helpers"
	"github.com/bwmarrin/discordgo"
)

type InvocationData struct {
	Timestamp  time.Time
	ReplyMsgID string
	ChannelID  string
}

var previousInvocations = make(map[string]InvocationData)
var mutex sync.RWMutex
var preamble string = `
#set page(fill: black, height: auto, width: auto, margin: 10pt)
#set text(fill: white, size: 18pt)
#show math.equation: eq => [
  #text(fill: white, [ #eq ])
]

`

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "meow!ty") {
		code := preamble + strings.TrimSpace(m.Content[8:])
		file, err := helpers.RenderTypst(code)

		if err != nil {
			s.ChannelMessageSendEmbed(m.ChannelID, helpers.ErrorEmbed("rendering Typst", err)[0])
			return
		}

		msg, err := s.ChannelFileSend(m.ChannelID, file.Name, file.Reader)

		if err != nil {
			s.ChannelMessageSendEmbed(m.ChannelID, helpers.ErrorEmbed("sending attachment", err)[0])
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		previousInvocations[m.ChannelID] = InvocationData{
			Timestamp:  time.Now(),
			ReplyMsgID: msg.ID,
			ChannelID:  msg.ChannelID,
		}
	}
}

// clean up invocations after 5 minutes
func init() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			mutex.Lock()
			for channelID, data := range previousInvocations {
				if time.Since(data.Timestamp) > 5*time.Minute {
					delete(previousInvocations, channelID)
					log.Printf("cleaned up expired invocation for channel: %s", channelID)
				}
			}

			mutex.Unlock()
		}
	}()
}
