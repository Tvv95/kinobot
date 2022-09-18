package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"kinobot/internal/kinopoisk"
)

type tgBot struct {
	bot             *tgbotapi.BotAPI
	clientKinopoisk *kinopoisk.ClientKinopoisk
}

func NewTgBot(bot *tgbotapi.BotAPI, clientKinopoisk *kinopoisk.ClientKinopoisk) *tgBot {
	return &tgBot{bot: bot, clientKinopoisk: clientKinopoisk}
}
func (tb *tgBot) HandleUpdate(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.CallbackQuery != nil {
			if err := tb.handleCallback(update.CallbackQuery); err != nil {
				return err
			}
			continue
		}

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := tb.handleCommand(update.Message); err != nil {
				return err
			}
			continue
		}

		if err := tb.handleMessage(update.Message); err != nil {
			return err
		}
	}
	return nil
}
