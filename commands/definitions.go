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
		Handler:    SimpleHandlerFn(fun.Ping),
	},
	{
		Definition: fun.DefineCat,
		Handler:    SimpleHandlerFn(fun.Cat),
	},
	{
		Definition: fun.DefineFox,
		Handler:    SimpleHandlerFn(fun.Fox),
	},
	{
		Definition: fun.DefineDog,
		Handler:    SimpleHandlerFn(fun.Dog),
	},
	// Utility
	{
		Definition: utility.DefineUrban,
		Handler:    SimpleHandlerFn(utility.Urbandictionary),
	},
	{
		Definition: utility.DefineAvatar,
		Handler:    SimpleHandlerFn(utility.Avatar),
	},
	{
		Definition: utility.DefineBanner,
		Handler:    SimpleHandlerFn(utility.Banner),
	},
	// Moderation
	{
		Definition: moderation.DefineKick,
		Handler:    SimpleHandlerFn(moderation.Kick),
	},
	{
		Definition: moderation.DefineBan,
		Handler:    SimpleHandlerFn(moderation.Ban),
	},
	{
		Definition: moderation.DefineWarn,
		Handler:    SimpleHandlerFn(moderation.Warn),
	},
	{
		Definition: moderation.DefinePurge,
		Handler:    SimpleHandlerFn(moderation.Purge),
	},
	{
		Definition: moderation.DefineMute,
		Handler:    SimpleHandlerFn(moderation.Mute),
	},
	{
		Definition: moderation.DefineUnmute,
		Handler:    SimpleHandlerFn(moderation.Unmute),
	},
}
