package commands

import (
	"github.com/arithefirst/whisker/commands/fun"
	"github.com/bwmarrin/discordgo"
)

func GetCommandSetupComponents() (func(*discordgo.Session, *discordgo.InteractionCreate), []*discordgo.ApplicationCommand, error) {
	// TODO: Automate fetching of command defs so we don't have to update this for every new command
	return handleInteraction, []*discordgo.ApplicationCommand{fun.DefinePing}, nil
}

// handleInteractions routes an interaction to its handler
func handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// TODO: Automate this so we don't have to update it for every new command
	switch i.ApplicationCommandData().Name {
	case "ping":
		fun.Ping(s, i)
	}
}
