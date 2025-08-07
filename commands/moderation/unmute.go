package moderation

import (
	"fmt"

	"github.com/arithefirst/whisker/helpers"
	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

var DefineUnmute = &discordgo.ApplicationCommand{
	Name:        "unmute",
	Description: "Unmute a member",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "user",
			Description: "The member to unmute",
			Type:        discordgo.ApplicationCommandOptionUser,
			Required:    true,
		},
	},
}

func Unmute(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !hasPermission(s, i, discordgo.PermissionModerateMembers, "Unmute Members") {
		return
	}

	var targetUser *discordgo.User

	for _, option := range i.ApplicationCommandData().Options {
		switch option.Name {
		case "user":
			targetUser = option.UserValue(s)
		}
	}

	if targetUser == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "User not found or not specified.",
			},
		})
		return
	}

	err := s.GuildMemberTimeout(i.GuildID, targetUser.ID, nil)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:  discordgo.MessageFlagsEphemeral,
				Embeds: helpers.ErrorEmbed("Unmute failed", err),
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				helpers.
					CreateEmbed().
					SetTitle("Unmute Successful").
					SetDescription(fmt.Sprintf("%s has been unmuted by %s", targetUser.Mention(), i.Member.Mention())).
					SetThumbnail(targetUser.AvatarURL("2048")).
					SetColor(colors.Primary).MessageEmbed,
			},
		},
	})
}
