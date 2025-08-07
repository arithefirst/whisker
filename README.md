# Whisker

A multi-purpose discord bot built by [thatmagicalcat](https://thatmagicalcat.pages.dev) and [arithefirst](https://arithefirst.com)

## Features

Whisker is a versatile Discord bot designed to enhance your server experience with fun commands, utility tools, and moderation features. A checkbox next to a command means it is implemented. An unchecked box means that we're currently working on implementing that command.

### 🧰 Utility Commands

#### XP System
- [ ] **XP System** - Non-intrusive XP tracking with opt-in/out notifications
- [ ] **Leaderboard** - View server XP leaderboard
- [ ] **Rank** - Check your XP rank

#### Moderation
- [x] **Kick** - Kick users from server
- [x] **Ban** - Ban users from server
- [ ] **Mute** - Mute users in channels
- [x] **Warn** - Warn users
- [ ] **Warn Log** - View user warnings
- [ ] **Snipe** - Show last deleted message
- [ ] **Edit Snipe** - Show last edited message
- [x] **Clear** - Mass delete messages
- [ ] **Message Logging** - Log message edits and deletions

#### General Utilities
- [ ] **Remind Me** - Set personal reminders
- [X] **Urban Dictionary** - Look up definitions
- [ ] **LaTeX** - Render LaTeX expressions
- [ ] **Typst** - Render Typst documents
- [ ] **Wolfram Alpha** - Query Wolfram Alpha
- [ ] **Calculator** - Mathematical calculations
- [ ] **Time Zone** - Display time zones
- [x] **Avatar** - Get user avatars
- [ ] **Banner** - Get user banners
- [ ] **Image Processing** - Rotate, crop, zoom, flip images
- [ ] **Meme Generator** - Create memes with captions and templates

### 🎉 Fun Commands

#### Social/Relationship
- [ ] **Ship** - Calculate compatibility between users
- [ ] **Marry** - Marry another user
- [ ] **Divorce** - End a marriage
- [ ] **Adopt** - Adopt another user
- [ ] **Exes** - View relationship history

#### Games & Randomness
- [ ] **8Ball** - Ask the magic 8-ball a question
- [ ] **Trivia** - Play trivia questions (with XP rewards)
- [ ] **Fortune** - Get fortune cookie wisdom
- [ ] **Coinflip** - Flip a coin (XP gambling)
- [ ] **Slot Machine** - Play slots (XP gambling)
- [ ] **Roulette** - Play roulette (XP gambling)
- [ ] **Blackjack** - Single-player blackjack (XP gambling)
- [ ] **Guess the Number** - Number guessing game (XP gambling)
- [ ] **Rock Paper Scissors** - Play RPS (XP gambling)

#### Text Tools
- [ ] **Leet Speak** - Convert text to leet speak
- [ ] **ASCII Art** - Generate ASCII art
- [ ] **Mad Libs** - Generate Mad Libs stories

#### Animal Images
- [ ] **Fox** - Random fox images
- [ ] **Cat** - Random cat images
- [ ] **Dog** - Random dog images

## Contributing

### Rules for adding a command

For the following, this assumes the name of your command is `examplecommand`

- The command definition must be at the top of the file, and be formatted as
  ```go
    var DefineExamplecommnd = &discordgo.ApplicationCommand{
    Name:        "examplecommand",
    Description: "An example command",
    Options: []*discordgo.ApplicationCommandOption{
      // put command options here
    },
  }
  ```

- The command implementation must be below the definition, and be formatted as
  ```go
  func Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {

    // Command Logic & Option parsing goes here

    // Stuff for the response goes below
    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
      Type: discordgo.InteractionResponseChannelMessageWithSource,
      Data: &discordgo.InteractionResponseData{
        Content: "Pong!",
      },
    })
  }
  ```
  

- Every time a command is added, you must add it's definition to [definitions.go](/commands/definitions.go), and you must add it's name and impl to the switch in [dispatcher.go](/commands/dispatcher.go) (The case in the switch should be the same as the name in the command definition)