package main

import (
	"bytes"
	"log"

	"github.com/bwmarrin/discordgo"
)

func promptsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isNotMentioned(s, m.Content) {
		return
	}

	prompt, params := extractMsg(m.Content)
	log.Println("Prompt:", prompt)
	log.Println("Params:", params)

	switch prompt {
	case "ping":
		handlePing(s, m)
	case "define":
		handleDefine(s, m)
	default:
		_, err := s.ChannelMessageSend(m.ChannelID, "I don't know what you mean. Type /help to see all available prompts.")
		if err != nil {
			log.Println("Error sending message, ", err)
		}
	}
}

func commandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "help":
		var buffer bytes.Buffer

		buffer.WriteString("Here are all the available prompts:")
		buffer.WriteString("\n\n`ping` - Pong!")
		buffer.WriteString("\n\n`define [word]` - Get a definition for a word.")
		buffer.WriteString("\n\n`random [amount]` - Get random words and their definitions.")
		buffer.WriteString("\n\n`save [word]` - Get word definition and save it to database.")
		buffer.WriteString("\n\n`symnonym [word]` - Find words with similar meaning.")
		buffer.WriteString("\n\n`atonnym [word]` - Find words with opposite meaning.")
		buffer.WriteString("\n\nTag me at the end of each prompt to get a response.")

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: buffer.String(),
			},
		})
		if err != nil {
			log.Println("Error sending message, ", err)
		}
	}
}

func handlePing(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
	if err != nil {
		log.Println("Error sending message, ", err)
	}
}

func handleDefine(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "coming soon!")
	if err != nil {
		log.Println("Error sending message, ", err)
	}
}
