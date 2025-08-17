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
- [x] **Mute** - Mute users in channels
- [x] **Warn** - Warn users
- [ ] **Warn Log** - View user warnings
- [ ] **Snipe** - Show last deleted message
- [ ] **Edit Snipe** - Show last edited message
- [x] **Purge** - Mass delete messages
- [ ] **Message Logging** - Log message edits and deletions

#### General Utilities
- [ ] **Remind Me** - Set personal reminders
- [X] **Urban Dictionary** - Look up definitions
- [ ] **LaTeX** - Render LaTeX expressions
- [x] **Typst** - Render Typst documents
- [ ] **Wolfram Alpha** - Query Wolfram Alpha
- [ ] **Calculator** - Mathematical calculations
- [ ] **Time Zone** - Display time zones
- [x] **Avatar** - Get user avatars
- [X] **Banner** - Get user banners
- [ ] **Image Processing** - Rotate, crop, zoom, flip images
- [ ] **Meme Generator** - Create memes with captions and templates
- [ ] **Tags system**
    - /tag create <name> <content>
    - /tag delete <name>
    - /tag <name>
    - Restrict create/delete to users with Level-Up role

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
- [X] **Fox** - Random fox images
- [X] **Cat** - Random cat images
- [X] **Dog** - Random dog images

## Contributing

### Rules for adding a command

For the following, this assumes the name of your command is `examplecommand`

- The command definition must be at the top of the file, and be formatted as:
  File `examplecommand.go`

  ```go
    var DefineExamplecommnd = &discordgo.ApplicationCommand{
    Name:        "examplecommand",
    Description: "An example command",
    Options: []*discordgo.ApplicationCommandOption{
      // put command options here
    },
  }

  // The command implementation must be below the definition, and be formatted as

  // or func Ping(s *discrodgo.Session, i *discordgo.InteractionCreate, db *pgxpool.Pool)
  // if it requires a database connection

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

- Every time a command is added, you must add it's definition and impl to [definitions.go](/commands/definitions.go).
  For example:

  ```go
  var commandRegistry = []Command{
    // other definitions
    {
      Definition: DefineExamplecommnd,
      Handler:    Ping,
    }
    // ...
  }
  ```
