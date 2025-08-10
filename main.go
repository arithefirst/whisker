package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/arithefirst/whisker/commands"
	"github.com/arithefirst/whisker/events"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	BotToken = flag.String("token", "", "Bot access token")
	// hardcoded for now
	// GuildID  = flag.String("guild", "", "Guild ID (for testing, optional)")
)

func init() {
	flag.Parse()
}

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		databaseUrl = "postgres://botuser:botpassword@localhost:5432/botdb"
	}

	dbpool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	log.Println("Connected to databse")

	defer dbpool.Close()

	// dependency injection
	eventHandler := &events.Handler {
		DB: dbpool,
	}

	cmdHandler := &commands.Handler {
        DB: dbpool,
	}

	GuildID := "1402745840220635187"

	if *BotToken == "" {
		*BotToken = os.Getenv("DISCORD_TOKEN")
		if *BotToken == "" {
			log.Fatal("No bot token provided. Use -token flag.")
		}
	}

	client, err := discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot token: %v", err)
	}

	client.Identify.Intents |= discordgo.IntentsGuilds |
		discordgo.IntentsGuildMembers |
		discordgo.IntentsGuildMessages |
		discordgo.IntentMessageContent

	commandHandler, commandDefs, err := cmdHandler.GetCommandSetupComponents()
	if err != nil {
		log.Fatalf("Error barreling commands: %v", err)
	}

	eventHandler.RegisterEvents(client)
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
		_, err := client.ApplicationCommandCreate(client.State.User.ID, GuildID, v)
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
