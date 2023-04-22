package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

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
