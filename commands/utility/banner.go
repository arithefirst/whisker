package utility

import (
	"fmt"

	"github.com/arithefirst/whisker/helpers"
	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

var DefineBanner = &discordgo.ApplicationCommand{
	Name:        "banner",
	Description: "Gets the banner of a user or yourself",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "user",
			Description: "The user to get the banner of",
			Type:        discordgo.ApplicationCommandOptionUser,
			Required:    false,
		},
	},
}

func Banner(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	userResponse, err := helpers.ManualGetUser(user.ID, s)
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("getting user info", err))
		return
	}

	var bannerURL string
	if len(user.BannerURL("256")) > 1 {
		bannerURL = user.BannerURL("256")
	} else if userResponse.AccentColor != nil {
		bannerURL = fmt.Sprintf("https://singlecolorimage.com/get/%06X/600x240", *userResponse.AccentColor)
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "This user does not have a banner or accent color!",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}

	helpers.IntRespondEmbed(s, i, []*discordgo.MessageEmbed{
		helpers.
			CreateEmbed().
			SetTitle(fmt.Sprintf("%s's Banner", user.Username)).
			SetImage(bannerURL).
			SetColor(colors.Primary).MessageEmbed,
	})
}
