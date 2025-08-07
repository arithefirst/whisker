package commands

import "github.com/bwmarrin/discordgo"

// Command holds the definition and handler for a slash command.
type Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    func(s *discordgo.Session, i *discordgo.InteractionCreate)
}
