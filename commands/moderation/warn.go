package moderation

import (
	"fmt"

	"github.com/arithefirst/whisker/helpers"
	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

var DefineWarn = &discordgo.ApplicationCommand{
	Name:        "warn",
	Description: "Warns a user in the server",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "user",
			Description: "The user to warn",
			Type:        discordgo.ApplicationCommandOptionUser,
			Required:    true,
		},
		{
			Name:        "reason",
			Description: "The reason for warning the user",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

func Warn(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// is kick appropriate for this?
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

	channel, err := s.UserChannelCreate(user.ID)
	if err != nil {
		helpers.IntRespondEmbedEph(s,i,helpers.ErrorEmbed("creating DM channel", err))
		return
	}

	// TODO: log the warning in db

	guild, err := s.State.Guild(i.GuildID)
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("fetching guild", err))
		return
	}

	_, err = s.ChannelMessageSend(channel.ID, fmt.Sprintf("You have been warned in %s for: `%s`", guild.Name, reason))
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("sending DM", err))
		return
	}

	helpers.IntRespondEmbed(s, i, []*discordgo.MessageEmbed{
		helpers.
			CreateEmbed().
			SetTitle("User Warned").
			SetDescription(fmt.Sprintf("%s has been warned by %s", user.Mention(), i.Member.Mention())).
			AddField("Reason", fmt.Sprintf("`%s`", reason)).
			SetThumbnail(user.AvatarURL("2048")).
			SetColor(colors.Primary).MessageEmbed,
	})
}
