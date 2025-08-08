package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/arithefirst/whisker/commands"
	"github.com/arithefirst/whisker/events"
	"github.com/bwmarrin/discordgo"
)

var (
	BotToken = flag.String("token", "", "Bot access token")
	GuildID  = flag.String("guild", "", "Guild ID (for testing, optional)")
)

func init() {
	flag.Parse()
}

func main() {
	if *BotToken == "" {
		log.Fatal("No bot token provided. Use -token flag.")
	}

	client, err := discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot token: %v", err)
	}

	client.Identify.Intents |= discordgo.IntentsGuilds |
		discordgo.IntentsGuildMembers |
		discordgo.IntentsGuildMessages |
		discordgo.IntentMessageContent

	commandHandler, commandDefs, err := commands.GetCommandSetupComponents()
	if err != nil {
		log.Fatalf("Error barreling commands: %v", err)
	}

	events.RegisterEvents(client)
	// Register the interaction handler
	client.AddHandler(commandHandler)

	err = client.Open()
	if err != nil {
		log.Fatalf("Error opening client: %v", err)
	}
	defer client.Close()

	// Register commands
	var commandCounter uint16 = 0
	for _, v := range commandDefs {
		_, err := client.ApplicationCommandCreate(client.State.User.ID, *GuildID, v)
		if err != nil {
			log.Fatalf("Error creating command '%s': %v", v.Name, err)
		}
		commandCounter += 1
	}
	fmt.Printf("Registered %d commands.\n", commandCounter)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("Bot is running. Press Ctrl+C to exit.")
	<-stop

	log.Println("Bot shut down.")
}
