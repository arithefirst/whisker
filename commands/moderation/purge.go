package moderation

import (
	"fmt"
	"time"

	"github.com/arithefirst/whisker/helpers"
	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

var DefinePurge = &discordgo.ApplicationCommand{
	Name:        "purge",
	Description: "Purges a specified number of messages from the current channel",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "count",
			Description: "The number of messages to purge (1-100)",
			Type:        discordgo.ApplicationCommandOptionInteger,
			Required:    true,
			MinValue:    &[]float64{1}[0], // wth?
			MaxValue:    100,
		},
		{
			Name:        "autodelete",
			Description: "Wether to delete the bot response automatically (Default: True)",
			Type:        discordgo.ApplicationCommandOptionBoolean,
			Required:    false,
		},
	},
}

func Purge(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !hasPermission(s, i, discordgo.PermissionManageMessages, "Manage Messages") {
		return
	}

	options := i.ApplicationCommandData().Options
	count := int(options[0].IntValue())
	audodelete := true
	for _, opt := range options {
		if opt.Name == "autodelete" {
			audodelete = opt.BoolValue()
			break
		}
	}

	messages, err := s.ChannelMessages(i.ChannelID, count, "", "", "")
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("fetching messages", err))
		return
	}

	if len(messages) == 0 {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("purging messages", fmt.Errorf("no messages found to clear")))
		return
	}

	messageIDs := make([]string, len(messages))
	msgDeletedPerUser := make(map[string]int)

	for idx, msg := range messages {
		msgDeletedPerUser[msg.Author.Mention()]++
		messageIDs[idx] = msg.ID
	}

	err = s.ChannelMessagesBulkDelete(i.ChannelID, messageIDs)
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("deleting messages", err))
		return
	}

	var messageBreakdown string
	for user, count := range msgDeletedPerUser {
		messageBreakdown += fmt.Sprintf("\n- %s: %d Deletions", user, count)
	}

	helpers.IntRespondEmbed(s, i, []*discordgo.MessageEmbed{
		helpers.
			CreateEmbed().
			SetDescription(fmt.Sprintf("Successfully deleted %d messages.", count)).
			SetTitle("Messages Purged").
			AddField("Message Breakdown", messageBreakdown).
			SetColor(colors.Primary).MessageEmbed,
	})

	// delete the reply after 3 secs if autodelete on
	if audodelete {
		go func() {
			time.Sleep(3 * time.Second)
			s.InteractionResponseDelete(i.Interaction)
		}()
	}

}
