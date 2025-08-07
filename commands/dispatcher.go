package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	commandDefinitions []*discordgo.ApplicationCommand
	commandHandlers    = make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate))
)

func init() {
	for _, cmd := range commandRegistry {
		commandDefinitions = append(commandDefinitions, cmd.Definition)
		commandHandlers[cmd.Definition.Name] = cmd.Handler
	}
}

func GetCommandSetupComponents() (func(*discordgo.Session, *discordgo.InteractionCreate), []*discordgo.ApplicationCommand, error) {
	return handleInteraction, commandDefinitions, nil
}

// handleInteractions dispatches the appropriate command handlers based on the incoming interaction
func handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		handler(s, i)
	}
}
