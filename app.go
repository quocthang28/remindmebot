package main

import (
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/bwmarrin/discordgo"
)

type Config struct {
	BotPrefix string `json:"BotPrefix"`
	AppID     string `json:"AppID"`
	GuildID   string `json:"GuildID"`
	BotID     string `json:"BotID"`
}

type App struct {
	Config         Config
	DiscordSession *discordgo.Session
	Firestore      *firestore.Client
}

func NewApp() *App {
	app := App{
		Config: getAppConfig(),
	}

	ds, err := discordgo.New("Bot " + string(decrypt(os.Getenv("BOT_TOKEN_K"), os.Getenv("BOT_TOKEN"))))
	if err != nil {
		log.Fatal("Error creating Discord session, ", err)
	}

	app.DiscordSession = ds

	app.DiscordSession.Identify.Intents = discordgo.IntentsGuildMessages
	app.DiscordSession.AddHandler(promptsHandler)
	app.DiscordSession.AddHandler(commandsHandler)
	app.registerDiscordCmds()

	app.Firestore, err = NewFirebaseClient()
	if err != nil {
		log.Fatal("Error creating Firestore client, ", err)
	}

	return &app
}

func (a *App) registerDiscordCmds() {
	_, err := a.DiscordSession.ApplicationCommandBulkOverwrite(a.Config.AppID, a.Config.GuildID, []*discordgo.ApplicationCommand{
		{
			Name:        "help",
			Description: "List all available prompts.",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
