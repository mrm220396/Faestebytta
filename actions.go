package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func DiceRoll(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	roll := strings.TrimPrefix(update.Message.Text, "/roll ")
	parts := strings.FieldsFunc(roll, func(r rune) bool {
		return r == 'd' || r == ' '
	})
	if len(parts) != 2 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid roll format. Usage: /roll <number of dice>d<sides of dice> or /roll <number of dice> <sides of dice>")
		bot.Send(msg)
		return
	}
	numDice, err := strconv.Atoi(parts[0])
	if err != nil || numDice < 1 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid number of dice.")
		bot.Send(msg)
		return
	}
	numSides, err := strconv.Atoi(parts[1])
	if err != nil || numSides < 2 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid number of sides.")
		bot.Send(msg)
		return
	}
	result := 0
	for i := 0; i < numDice; i++ {
		result += rand.Intn(numSides) + 1
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, strconv.Itoa(result))
	bot.Send(msg)
}

func kick(bot *tgbotapi.BotAPI, chatID int64, userID int) {
	kickConfig := tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: chatID,
			UserID: userID,
		},
		UntilDate: time.Now().Add(1 * time.Minute).Unix(), // Ban for 5 minutes
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

func ban(bot *tgbotapi.BotAPI, chatID int64, userID int) {
	banConfig := tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: chatID,
			UserID: userID,
		},
		UntilDate: time.Now().Add(701280 * time.Hour).Unix(),
	}

	_, err := bot.KickChatMember(banConfig)

	if err != nil {
		fmt.Errorf("Errors were found %s", err)
	}

	msg := tgbotapi.NewMessage(chatID, "__Won't be back in 80 years__")

	msg.ParseMode = "Markdown"
	bot.Send(msg)
}

func help(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Available commands:\n/start - Start the bot\n/help - Display available commands\n/echo <message> - Echo back a message\n/roll <number of dice>d<sides of dice> - Roll some dice\n/kick (only admins) Reply someone to kick\n/ban Ban users for 80 years.")
	bot.Send(msg)
}

func echo(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	echo := strings.TrimPrefix(update.Message.Text, "/echo ")
	response := "*" + echo + "*" + " - By " + update.Message.From.LastName + update.Message.From.FirstName

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	msg.ParseMode = "markdown"
	bot.Send(msg)
}
