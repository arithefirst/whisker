package helpers

import (
	"fmt"

	colors "github.com/arithefirst/whisker/helpers/embedColors"
	"github.com/bwmarrin/discordgo"
)

func ErrorEmbed(errorContext string, err error) []*discordgo.MessageEmbed {
	return []*discordgo.MessageEmbed{
		CreateEmbed().
			SetTitle("Unexpected error!").
			SetDescription(fmt.Sprintf("An unexpected error occured when %s. See details below.", errorContext)).
			AddField("Full Error Details", fmt.Sprintf("```text\n%v\n```", err)). // Use colors.Error
			SetColor(colors.Error).MessageEmbed,
	}
}
