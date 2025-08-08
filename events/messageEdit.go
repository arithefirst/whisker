package events

import (
	"strings"
	"time"

	"github.com/arithefirst/whisker/helpers"
	"github.com/bwmarrin/discordgo"
)

func messageEdit(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if m.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "meow!ty") {
		mutex.Lock()
		invocation, found := previousInvocations[m.ID]
		mutex.Unlock()

		// update timestamp
		invocation.Timestamp = time.Now()

		s.ChannelMessageDelete(invocation.ChannelID, invocation.ReplyMsgID)
		if invocation.ErrorMsgID != "" {
			s.ChannelMessageDelete(invocation.ChannelID, invocation.ErrorMsgID)
		}

		if !found {
			return
		}

		code := preamble + strings.TrimSpace(m.Content[8:])
		file, err := helpers.RenderTypst(code)

		if err != nil {
			errorMsg, _ := s.ChannelMessageSendEmbed(m.ChannelID, helpers.ErrorEmbed("rendering Typst", err)[0])

			mutex.Lock()
			defer mutex.Unlock()

			invocation.ErrorMsgID = errorMsg.ID
			previousInvocations[m.ID] = invocation

			return
		}

		msg, err := s.ChannelFileSend(m.ChannelID, file.Name, file.Reader)
		if err != nil {
			s.ChannelMessageSendEmbed(m.ChannelID, helpers.ErrorEmbed("sending attachment", err)[0])
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		invocation.ReplyMsgID = msg.ID
		previousInvocations[m.ID] = invocation
	}
}
