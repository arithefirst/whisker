package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    any
}

type Handler struct {
	DB *pgxpool.Pool
}

var (
	commandDefinitions []*discordgo.ApplicationCommand
	commandHandlers    = make(map[string]any)
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
	if f, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		execute(s, i, h.DB, f)
	}
}

func execute(s *discordgo.Session, i *discordgo.InteractionCreate, db *pgxpool.Pool, f any) {
	switch v := f.(type) {
	case func(*discordgo.Session, *discordgo.InteractionCreate):
		v(s, i)
	case func(*discordgo.Session, *discordgo.InteractionCreate, *pgxpool.Pool):
		v(s, i, db)
	default:
		log.Printf("Invalid handler type, handler will never be called")
	}
}
