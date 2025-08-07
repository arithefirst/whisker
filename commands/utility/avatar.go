package utility

import (
	"fmt"

	"github.com/arithefirst/whisker/helpers"
	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

var DefineAvatar = &discordgo.ApplicationCommand{
	Name:        "avatar",
	Description: "Gets the avatar of a user or yourself",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "user",
			Description: "The user to get the avatar of",
			Type:        discordgo.ApplicationCommandOptionUser,
			Required:    false,
		},
	},
}

func Avatar(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var user *discordgo.User

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		if option.Name == "user" {
			user = option.UserValue(s)
			break
		}
	}

	if user == nil {
		user = i.Member.User
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				helpers.
					CreateEmbed().
					SetTitle(fmt.Sprintf("%s's Avatar", user.Username)).
					SetImage(user.AvatarURL("2048")).
					SetColor(colors.Primary).MessageEmbed,
			},
		},
	})
}
