package moderation

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// hasPermission checks if the user who triggered the interaction has the required permission.
// it sends an ephemeral error message and returns false if they do not
func hasPermission(s *discordgo.Session, i *discordgo.InteractionCreate, permission int64, permissionName string) bool {
	if i.Member.Permissions&discordgo.PermissionAdministrator != discordgo.PermissionAdministrator && i.Member.Permissions&permission != permission {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: fmt.Sprintf("you do not have the required permissions to use this command: `%s`", permissionName),
			},
		})

		return false
	}

	return true
}
