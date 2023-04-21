package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Config struct {
	BotToken string `json:"bot_token"`
}

func readConfigFile(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)

	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func kick(bot *tgbotapi.BotAPI, chatID int64, userID int) {
	kickConfig := tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: chatID,
			UserID: userID,
		},
		UntilDate: time.Now().Add(5 * time.Minute).Unix(), // Ban for 5 minutes
	}

	_, err := bot.KickChatMember(kickConfig)
	if err != nil {
		log.Fatalf("Invalid argument %v", err)
	}

	resp := fmt.Sprintf("User %d has been kicked!", userID)

	msg := tgbotapi.NewMessage(chatID, resp)
	bot.Send(msg)

	// Sassy sentence
	resp = "_This dude might have done shit_"
	msg = tgbotapi.NewMessage(chatID, resp)
	msg.ParseMode = "markdown"
	bot.Send(msg)
}

func main() {

	config, err := readConfigFile("settings/config.json")
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		go handleMessage(bot, update)
	}
}
