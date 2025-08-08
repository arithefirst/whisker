package events

import (
	"strings"

	"github.com/arithefirst/whisker/helpers"
	"github.com/bwmarrin/discordgo"
)

var preamble string = `
#set page(fill: black, height: auto, width: auto, margin: 10pt)
#set text(fill: white, size: 18pt)
#show math.equation: eq => [
  #text(fill: white, [ #eq ])
]

`

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "meow!ty") {
		code := preamble + strings.TrimSpace(m.Content[8:])
		file, err := helpers.RenderTypst(code)

		if err != nil {
			s.ChannelMessageSendEmbed(m.ChannelID, helpers.ErrorEmbed("rendering Typst", err)[0]);
			return
		}

		_, err = s.ChannelFileSend(m.ChannelID, file.Name, file.Reader)

		if err != nil {
			s.ChannelMessageSendEmbed(m.ChannelID, helpers.ErrorEmbed("sending attachment", err)[0]);
			return
		}
	}
}
