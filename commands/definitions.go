package commands

import (
	"github.com/arithefirst/whisker/commands/fun"
	"github.com/arithefirst/whisker/commands/moderation"
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
		Definition: moderation.DefineClear,
		Handler:    moderation.Clear,
	},
	{
		Definition: moderation.DefineMute,
		Handler:    moderation.Mute,
	}, {
		Definition: moderation.DefineUnmute,
		Handler:    moderation.Unmute,
	},
}
