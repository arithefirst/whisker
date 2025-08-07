package moderation

import (
	"fmt"
	"time"

	"github.com/arithefirst/whisker/helpers"
	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

var DefineMute = &discordgo.ApplicationCommand{
	Name:        "mute",
	Description: "Mute a member",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "user",
			Description: "The member to mute",
			Type:        discordgo.ApplicationCommandOptionUser,
			Required:    true,
		},
		{
			Name:        "duration",
			Description: "Duration (e.g. 5m, 1h, 2h30m)",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
		{
			Name:        "reason",
			Description: "Reason for Mute",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
		},
	},
}

func Mute(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !hasPermission(s, i, discordgo.PermissionModerateMembers, "Mute Members") {
		return
	}

	var targetUser *discordgo.User
	var reason, durationStr string

	for _, option := range i.ApplicationCommandData().Options {
		switch option.Name {
		case "user":
			targetUser = option.UserValue(s)
		case "duration":
			durationStr = option.StringValue()
		case "reason":
			reason = option.StringValue()
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

	// Parse duration
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Invalid duration format. Use like: 5m, 1h, 2h30m",
			},
		})
		return
	}

	until := time.Now().Add(duration)

	err = s.GuildMemberTimeout(i.GuildID, targetUser.ID, &until)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:  discordgo.MessageFlagsEphemeral,
				Embeds: helpers.ErrorEmbed("Mute failed", err),
			},
		})
		return
	}

	if reason == "" {
		reason = "*No reason provided*"
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				helpers.
					CreateEmbed().
					SetTitle("Muted Successful").
					SetDescription(fmt.Sprintf(
						"%s has been muted by %s for %s",
						targetUser.Mention(),
						i.Member.Mention(),
						durationStr)).
					AddField("Duration", durationStr).
					AddField("Reason", reason).
					SetThumbnail(targetUser.AvatarURL("2048")).
					SetColor(colors.Primary).MessageEmbed,
			},
		},
	})
}
