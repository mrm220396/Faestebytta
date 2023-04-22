package main

import (
	"bots/pkg/quotes"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	isBotInitialized bool
	isNotAdminTold   bool
)

func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	switch {
	case strings.HasPrefix(update.Message.Text, "/start"):
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! I am your friendly neighborhood bot. How can I assist you?")
		bot.Send(msg)
		break

	case strings.HasPrefix(update.Message.Text, "/help"):
		help(bot, update)
		break

	case strings.HasPrefix(update.Message.Text, "/echo "):
		echo(bot, update)
		break

	case strings.HasPrefix(update.Message.Text, "/roll "):
		DiceRoll(bot, update)
		break

	case strings.HasPrefix(update.Message.Text, "/kick"):
		if update.Message.ReplyToMessage == nil {
			break
		}
		userID := update.Message.ReplyToMessage.From.ID
		kick(bot, update.Message.Chat.ID, userID)
		break

	case strings.HasPrefix(update.Message.Text, "/ban"):

		if update.Message.ReplyToMessage == nil {
			break
		}
		userID := update.Message.ReplyToMessage.From.ID

		ban(bot, update.Message.Chat.ID, userID)

	case strings.HasPrefix(update.Message.Text, "/quote"):
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, quotes.GetQuote())
		bot.Send(msg)

	case strings.HasPrefix(update.Message.Text, "/pin"):

		if update.Message.ReplyToMessage == nil {
			break
		}
		pinMessage(bot, update.Message.Chat.ID, update.Message.ReplyToMessage.MessageID, update.Message.ReplyToMessage.From.ID)

	case strings.HasPrefix(update.Message.Text, "/unpin"):
		unpinMessage(bot, update.Message.Chat.ID)

	default:
		break
	}
}
