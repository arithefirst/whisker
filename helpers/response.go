package helpers

import "github.com/bwmarrin/discordgo"

func IntRespondEmbedEph(s *discordgo.Session, i *discordgo.InteractionCreate, embeds []*discordgo.MessageEmbed) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "",
			Flags:   discordgo.MessageFlagsEphemeral,
			Embeds:  embeds,
		},
	})
}

func IntRespondEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, embeds []*discordgo.MessageEmbed) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "",
			Embeds:  embeds,
		},
	})
}
