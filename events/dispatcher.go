package events

import "github.com/bwmarrin/discordgo"

func RegisterEvents(s *discordgo.Session) {
	s.AddHandler(messageCreate)
	s.AddHandler(messageEdit)
	s.AddHandler(messageDelete)
}
