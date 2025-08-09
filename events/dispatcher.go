package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	DB *pgxpool.Pool
}

func (h *Handler) RegisterEvents(s *discordgo.Session) {
	s.AddHandler(h.messageCreate)
	s.AddHandler(messageEdit)
	s.AddHandler(messageDelete)
}
