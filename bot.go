package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

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

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
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

	// http.ListenAndServe("0.0.0.0:", nil)
}
