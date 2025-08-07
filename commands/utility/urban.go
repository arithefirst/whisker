package utility

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/arithefirst/whisker/helpers"
	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

type UrbanDictionaryResponse struct {
	List []struct {
		Author     string `json:"author"`
		Definition string `json:"definition"`
		Example    string `json:"example"`
		Permalink  string `json:"permalink"`
		Word       string `json:"word"`
		WrittenOn  string `json:"written_on"`
	} `json:"list"`
}

var DefineUrban = &discordgo.ApplicationCommand{
	Name:        "urban",
	Description: "Searches Urban Dictionary for a term",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "term",
			Description: "The term to search Urban Dictionary for",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

func injectSearchLinks(input string) string {
	query := regexp.MustCompile(`\[(.*?)\]`)
	return query.ReplaceAllStringFunc(input, func(match string) string {
		content := query.FindStringSubmatch(match)[1]
		return fmt.Sprintf("[%s](https://www.urbandictionary.com/define.php?term=%s)", content, url.QueryEscape(content))
	})
}

func Urbandictionary(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var searchTerm string

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		if option.Name == "term" {
			searchTerm = option.StringValue()
			break
		}
	}

	resp, err := http.Get(fmt.Sprintf("https://api.urbandictionary.com/v0/define?term=%s", url.QueryEscape(searchTerm)))
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:  discordgo.MessageFlagsEphemeral,
				Embeds: helpers.ErrorEmbed("querying Urban Dictionary", err),
			},
		})

		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data UrbanDictionaryResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:  discordgo.MessageFlagsEphemeral,
				Embeds: helpers.ErrorEmbed("parsing JSON response", err),
			},
		})

		return
	}

	if len(data.List) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					helpers.
						CreateEmbed().
						SetTitle("No results found").
						SetDescription(fmt.Sprintf("No definitions found for \"%s\"", searchTerm)).
						SetThumbnail("https://www.urbandictionary.com/apple-touch-icon.png").
						SetColor(colors.Error).MessageEmbed,
				},
			},
		})
		return
	}

	formattedDate, err := helpers.RFC3339toDateString(data.List[0].WrittenOn)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:  discordgo.MessageFlagsEphemeral,
				Embeds: helpers.ErrorEmbed("converting RFC339 String to date object", err),
			},
		})

		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				helpers.
					CreateEmbed().
					SetTitle(fmt.Sprintf("Definition of \"%s\"", data.List[0].Word)).
					AddField("Definition", injectSearchLinks(data.List[0].Definition)).
					AddField("Example", injectSearchLinks(data.List[0].Example)).
					SetURL(data.List[0].Permalink).
					SetThumbnail("https://www.urbandictionary.com/apple-touch-icon.png").
					SetAuthor(fmt.Sprintf("By %s • %s", data.List[0].Author, formattedDate)).
					SetColor(colors.Primary).MessageEmbed,
			},
		},
	})
}
