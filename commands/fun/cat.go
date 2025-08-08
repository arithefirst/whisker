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

var DefineCat = &discordgo.ApplicationCommand{
	Name:        "cat",
	Description: "Get a random cat image",
}

/*
[
  {
    "id": "EPF2ejNS0",
    "url": "https://cdn2.thecatapi.com/images/EPF2ejNS0.jpg",
    "width": 850,
    "height": 1008
  }
]
*/

type TheCatApiResponse []struct {
	Id     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func Cat(s *discordgo.Session, i *discordgo.InteractionCreate) {
	resp, err := http.Get("https://api.thecatapi.com/v1/images/search?size=med&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1")
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("querying thecatapi", err))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data TheCatApiResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("parsing JSON response", err))
		return
	}

	helpers.IntRespondEmbed(s, i, []*discordgo.MessageEmbed{
		helpers.
			CreateEmbed().
			SetTitle("Check out this cat!").
			SetImage(data[0].URL).
			SetColor(colors.Primary).MessageEmbed,
	})
}
