package moderation

import (
	"fmt"

	"github.com/arithefirst/whisker/helpers"
	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

var DefineKick = &discordgo.ApplicationCommand{
	Name:        "kick",
	Description: "Kicks a user from the server",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "user",
			Description: "The user to kick",
			Type:        discordgo.ApplicationCommandOptionUser,
			Required:    true,
		},
		{
			Name:        "reason",
			Description: "The reason for kicking the user",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
		},
	},
}

func Kick(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !hasPermission(s, i, discordgo.PermissionKickMembers, "Kick Members") {
		return
	}

	var user *discordgo.User
	var reason string

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch option.Name {
		case "user":
			user = option.UserValue(s)
		case "reason":
			reason = option.StringValue()
		}
	}

	var err error
	if reason == "" {
		err = s.GuildMemberDelete(i.GuildID, user.ID)
		reason = "*No reason provided*"
	} else {
		err = s.GuildMemberDeleteWithReason(i.GuildID, user.ID, reason)
		reason = fmt.Sprintf("`%s`", reason)
	}

	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("kicking user", err))
		return
	}

	helpers.IntRespondEmbed(s, i, []*discordgo.MessageEmbed{
		helpers.
			CreateEmbed().
			SetTitle("User Kicked").
			SetDescription(fmt.Sprintf("%s has been kicked from the server by %s", user.Mention(), i.Member.Mention())).
			AddField("Reason", reason).
			SetThumbnail(user.AvatarURL("2048")).
			SetColor(colors.Primary).MessageEmbed,
	})
}
