package events

import "github.com/bwmarrin/discordgo"

func messageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	mutex.Lock()
	invocation, found := previousInvocations[m.ID]
	mutex.Unlock()

	if !found {
		return
	}

	// delete everything
	delete(previousInvocations, m.ID)
	s.ChannelMessageDelete(invocation.ChannelID, invocation.ReplyMsgID)
	if invocation.ErrorMsgID != "" {
		s.ChannelMessageDelete(invocation.ChannelID, invocation.ErrorMsgID)
	}
}
