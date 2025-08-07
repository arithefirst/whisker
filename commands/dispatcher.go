package commands

import (
	"github.com/arithefirst/whisker/commands/fun"
	"github.com/arithefirst/whisker/commands/utility"
	"github.com/bwmarrin/discordgo"
)

func GetCommandSetupComponents() (func(*discordgo.Session, *discordgo.InteractionCreate), []*discordgo.ApplicationCommand, error) {
	return handleInteraction, commandDefinitions, nil
}

// handleInteractions routes an interaction to its handler
func handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Add to this switch case every time you define a new command
	switch i.ApplicationCommandData().Name {
	case "ping":
		fun.Ping(s, i)
	case "urban":
		utility.Urbandictionary(s, i)
	case "avatar":
		utility.Avatar(s, i)
	}
}
