package moderation

import (
	"fmt"
	"time"

	"github.com/arithefirst/whisker/helpers"
	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

var DefineClear = &discordgo.ApplicationCommand{
	Name:        "clear",
	Description: "Clears a specified number of messages from the current channel",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "count",
			Description: "The number of messages to clear (1-100)",
			Type:        discordgo.ApplicationCommandOptionInteger,
			Required:    true,
			MinValue:    &[]float64{1}[0], // wth?
			MaxValue:    100,
		},
	},
}

func Clear(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !hasPermission(s, i, discordgo.PermissionManageMessages, "Manage Messages") {
		return
	}

	options := i.ApplicationCommandData().Options
	count := int(options[0].IntValue())

	messages, err := s.ChannelMessages(i.ChannelID, count, "", "", "")
	if err != nil {
		helpers.IntRespondEmbedEph(s,i, helpers.ErrorEmbed("fetching messages", err))
		return
	}

	if len(messages) == 0 {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("clearing messages", fmt.Errorf("no messages found to clear")))
		return
	}

	messageIDs := make([]string, len(messages))
	for idx, msg := range messages {
		messageIDs[idx] = msg.ID
	}

	err = s.ChannelMessagesBulkDelete(i.ChannelID, messageIDs)
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("deleting messages", err))
		return
	}

	helpers.IntRespondEmbed(s, i, []*discordgo.MessageEmbed{
		helpers.
			CreateEmbed().
			SetTitle("Messages Cleared").
			SetDescription(fmt.Sprintf("Successfully deleted %d messages.", count)).
			SetColor(colors.Primary).MessageEmbed,
	})

	// delete the reply after 3 secs
	go func() {
		time.Sleep(3 * time.Second)
		s.InteractionResponseDelete(i.Interaction)
	}()

}
