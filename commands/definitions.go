package commands

import (
	"github.com/arithefirst/whisker/commands/fun"
	"github.com/arithefirst/whisker/commands/utility"
	"github.com/bwmarrin/discordgo"
)

// Add to this list every time you create a new command
var commandDefinitions = []*discordgo.ApplicationCommand{fun.DefinePing, utility.DefineUrban, utility.DefineAvatar}
