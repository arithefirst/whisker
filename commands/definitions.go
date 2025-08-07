package commands

import (
	"github.com/arithefirst/whisker/commands/fun"
	"github.com/arithefirst/whisker/commands/utility"
)

// Add to this list every time you create a new command
var commandRegistry = []Command{
	{
		Definition: fun.DefinePing,
		Handler:    fun.Ping,
	},
	{
		Definition: utility.DefineUrban,
		Handler:    utility.Urbandictionary,
	},
	{
		Definition: utility.DefineAvatar,
		Handler:    utility.Avatar,
	},
}
