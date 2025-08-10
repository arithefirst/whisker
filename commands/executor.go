package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    CommandExecutor
}

type CommandExecutor interface {
	Execute(s *discordgo.Session, i *discordgo.InteractionCreate, db *pgxpool.Pool)
}

// slash command handler function which does not require a db connection
type SimpleHandlerFn func(*discordgo.Session, *discordgo.InteractionCreate)

// slash command handler function which requires a db connection
type DBAwareHandlerFn func(*discordgo.Session, *discordgo.InteractionCreate, *pgxpool.Pool)

func (f SimpleHandlerFn) Execute(s *discordgo.Session, i *discordgo.InteractionCreate, db *pgxpool.Pool) {
	f(s, i)
}

func (f DBAwareHandlerFn) Execute(s *discordgo.Session, i *discordgo.InteractionCreate, db *pgxpool.Pool) {
	f(s, i, db)
}
