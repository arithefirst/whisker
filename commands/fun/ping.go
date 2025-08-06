package fun

import "github.com/bwmarrin/discordgo"

var DefinePing = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Replies with Pong!",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "ephemeral",
			Description: "Wether the response should be ephermal or not",
			Type:        discordgo.ApplicationCommandOptionBoolean,
			Required:    false,
		},
	},
}

func Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ephemeral := false

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		if option.Name == "ephemeral" {
			ephemeral = option.BoolValue()
			break
		}
	}

	var flags discordgo.MessageFlags
	if ephemeral {
		flags = discordgo.MessageFlagsEphemeral
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
			Flags:   flags,
		},
	})
}
