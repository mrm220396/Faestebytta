package main

import (
	"math/rand"
	"strconv"
	"strings"

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

func help(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Available commands:\n/start - Start the bot\n/help - Display available commands\n/echo <message> - Echo back a message\n/roll <number of dice>d<sides of dice> - Roll some dice\n/kick (only admins) Reply someone to kick")
	bot.Send(msg)
}

func echo(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	echo := strings.TrimPrefix(update.Message.Text, "/echo ")
	response := "*" + echo + "*" + " - By " + update.Message.From.LastName + update.Message.From.FirstName

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	msg.ParseMode = "markdown"
	bot.Send(msg)
}
