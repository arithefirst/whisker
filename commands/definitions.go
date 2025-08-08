package commands

import (
	"github.com/arithefirst/whisker/commands/fun"
	"github.com/arithefirst/whisker/commands/moderation"
	"github.com/arithefirst/whisker/commands/utility"
)

// Add to this list every time you create a new command
var commandRegistry = []Command{
	// Fun
	{
		Definition: fun.DefinePing,
		Handler:    fun.Ping,
	},
	{
		Definition: fun.DefineCat,
		Handler:    fun.Cat,
	},
	{
		Definition: fun.DefineFox,
		Handler:    fun.Fox,
	},
	{
		Definition: fun.DefineDog,
		Handler:    fun.Dog,
	},
	// Utility
	{
		Definition: utility.DefineUrban,
		Handler:    utility.Urbandictionary,
	},
	{
		Definition: utility.DefineAvatar,
		Handler:    utility.Avatar,
	},
	{
		Definition: utility.DefineBanner,
		Handler:    utility.Banner,
	},
	// Moderation
	{
		Definition: moderation.DefineKick,
		Handler:    moderation.Kick,
	},
	{
		Definition: moderation.DefineBan,
		Handler:    moderation.Ban,
	},
	{
		Definition: moderation.DefineWarn,
		Handler:    moderation.Warn,
	},
	{
		Definition: moderation.DefinePurge,
		Handler:    moderation.Purge,
	},
}
