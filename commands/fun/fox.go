package fun

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/arithefirst/whisker/helpers"
	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

var DefineFox = &discordgo.ApplicationCommand{
	Name:        "fox",
	Description: "Get a random fox image",
}

type RandomFoxCaResponse struct {
	Image string `json:"image"`
	Link  string `json:"link"`
}

func Fox(s *discordgo.Session, i *discordgo.InteractionCreate) {
	resp, err := http.Get("https://randomfox.ca/floof/")
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("querying randomfox.ca", err))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data RandomFoxCaResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("parsing JSON response", err))
		return
	}

	helpers.IntRespondEmbed(s, i, []*discordgo.MessageEmbed{
		helpers.
			CreateEmbed().
			SetTitle("Check out this fox!").
			SetImage(data.Image).
			SetColor(colors.Primary).MessageEmbed,
	})
}
