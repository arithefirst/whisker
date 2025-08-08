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
	},
}

func Purge(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !hasPermission(s, i, discordgo.PermissionManageMessages, "Manage Messages") {
		return
	}

	options := i.ApplicationCommandData().Options
	count := int(options[0].IntValue())

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

	descriptionText := fmt.Sprintf("Successfully deleted %d messages.", count)
	for user, count := range msgDeletedPerUser {
		descriptionText += fmt.Sprintf("\n%s %d", user, count)
	}

	helpers.IntRespondEmbed(s, i, []*discordgo.MessageEmbed{
		helpers.
			CreateEmbed().
			SetDescription(descriptionText).
			SetTitle("Messages Purged").
			SetColor(colors.Primary).MessageEmbed,
	})

	// delete the reply after 3 secs
	go func() {
		time.Sleep(3 * time.Second)
		s.InteractionResponseDelete(i.Interaction)
	}()

}
