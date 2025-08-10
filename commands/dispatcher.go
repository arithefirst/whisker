package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	DB *pgxpool.Pool
}

var (
	commandDefinitions []*discordgo.ApplicationCommand
	commandHandlers    = make(map[string]CommandExecutor)
)

func init() {
	for _, cmd := range commandRegistry {
		commandDefinitions = append(commandDefinitions, cmd.Definition)
		commandHandlers[cmd.Definition.Name] = cmd.Handler
	}
}

func (h *Handler) GetCommandSetupComponents() (func(*discordgo.Session, *discordgo.InteractionCreate), []*discordgo.ApplicationCommand, error) {
	return h.handleInteraction, commandDefinitions, nil
}

// handleInteractions dispatches the appropriate command handlers based on the incoming interaction
func (h *Handler) handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if executor, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		executor.Execute(s, i, h.DB)
	}
}
