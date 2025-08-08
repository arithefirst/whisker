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

var DefineDog = &discordgo.ApplicationCommand{
	Name:        "dog",
	Description: "Get a random dog image",
}

type DogApiResponse struct {
	Image string `json:"message"`
}

func Dog(s *discordgo.Session, i *discordgo.InteractionCreate) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("querying dog.ceo", err))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data DogApiResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		helpers.IntRespondEmbedEph(s, i, helpers.ErrorEmbed("parsing JSON response", err))
		return
	}

	helpers.IntRespondEmbed(s, i, []*discordgo.MessageEmbed{
		helpers.
			CreateEmbed().
			SetTitle("Check out this dog!").
			SetImage(data.Image).
			SetColor(colors.Primary).MessageEmbed,
	})
}
