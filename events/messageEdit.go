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
		data, found := previousInvocations[m.ChannelID]
		mutex.Unlock()

		if !found {
			return
		}

		code := preamble + strings.TrimSpace(m.Content[8:])
		file, err := helpers.RenderTypst(code)

		if err != nil {
			s.ChannelMessageSendEmbed(m.ChannelID, helpers.ErrorEmbed("rendering Typst", err)[0])
			return
		}

		err = s.ChannelMessageDelete(data.ChannelID, data.ReplyMsgID)
		if err != nil {
			s.ChannelMessageSendEmbed(m.ChannelID, helpers.ErrorEmbed("deleting previous message", err)[0])
			return
		}

		msg, err := s.ChannelFileSend(m.ChannelID, file.Name, file.Reader)

		if err != nil {
			s.ChannelMessageSendEmbed(m.ChannelID, helpers.ErrorEmbed("sending attachment", err)[0])
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		// update timestamp and message id
		previousInvocations[m.ChannelID] = InvocationData{
			Timestamp: time.Now(),
			ReplyMsgID: msg.ID,
			ChannelID: msg.ChannelID,
		}
	}
}
