package moderation

import (
	"fmt"

	"github.com/arithefirst/whisker/helpers"
	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

// TODO: add an option to bulk delete messages sent by the user
var DefineBan = &discordgo.ApplicationCommand{
	Name:        "ban",
	Description: "Bans a user from the server",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "user",
			Description: "The user to ban",
			Type:        discordgo.ApplicationCommandOptionUser,
			Required:    true,
		},
		{
			Name:        "reason",
			Description: "The reason for banning the user",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
		},
	},
}

func Ban(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !hasPermission(s, i, discordgo.PermissionBanMembers, "Ban Members") {
		return
	}

	var user *discordgo.User
	var reason string
	var deleteMessageDays int = 0

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch option.Name {
		case "user":
			user = option.UserValue(s)
		case "reason":
			reason = option.StringValue()
		case "days":
			deleteMessageDays = int(option.IntValue())
		}
	}

	var err error
	if reason == "" {
		err = s.GuildBanCreate(i.GuildID, user.ID, deleteMessageDays)
		reason = "*No reason provided*"
	} else {
		err = s.GuildBanCreateWithReason(i.GuildID, user.ID, reason, deleteMessageDays)
		reason = fmt.Sprintf("`%s`", reason)
	}

	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("banning user", err))
		return
	}

	helpers.IntRespondEmbed(s, i, []*discordgo.MessageEmbed{
		helpers.
			CreateEmbed().
			SetTitle("User Banned").
			SetDescription(fmt.Sprintf("%s has been kicked from the server by %s", user.Mention(), i.Member.Mention())).
			AddField("Reason", reason).
			SetThumbnail(user.AvatarURL("2048")).
			SetColor(colors.Primary).MessageEmbed,
	})
}
