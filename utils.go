package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func isNotMentioned(s *discordgo.Session, msg string) bool {
	if app.Config.BotID == s.State.User.ID {
		return true
	}

	var botID string

	if strings.Contains(msg, " ") {
		botID = msg[strings.LastIndex(msg, " "):]
	} else {
		botID = msg
	}

	if strings.Trim(botID, " ") != app.Config.BotID {
		return true
	}

	return false
}

func extractMsg(msg string) (string, string) {
	log.Println(msg)
	return msg[:strings.Index(msg, " ")], strings.Trim(msg[strings.Index(msg, " "):strings.LastIndex(msg, " ")], " ")
}

func getAppConfig() Config {

	config := Config{}

	err := json.Unmarshal(decrypt(os.Getenv("APP_CONFIG_K"), os.Getenv("APP_CONFIG")), &config)
	if err != nil {
		log.Fatal("Error parsing config file:", err.Error())
	}

	return config
}

// decrypt from base64 to decrypted string
func decrypt(keyString string, stringToDecrypt string) []byte {
	key, _ := hex.DecodeString(keyString)
	ciphertext, _ := base64.URLEncoding.DecodeString(stringToDecrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}
