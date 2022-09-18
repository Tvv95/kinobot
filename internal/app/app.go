package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"kinobot/internal/config"
	"kinobot/internal/kinopoisk"
	"kinobot/internal/telegram"
	"log"
)

type app struct {
	cfg *config.Config
	bot *tgbotapi.BotAPI
}

func NewApp(cfg *config.Config) (*app, error) {
	return &app{
		cfg: cfg,
	}, nil
}

func (a *app) Run() {
	tgBot, err := a.createTgBot()
	if err != nil {
		log.Println(err)
	}
	a.bot = tgBot
	if err = a.start(); err != nil {
		log.Println(err)
	}
}

func (a *app) createTgBot() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(a.cfg.TgToken)
	if err != nil {
		return nil, err
	}
	return bot, nil
}

func (a *app) start() error {
	log.Printf("Authorized on account %s", a.bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := a.initUpdateChannel()
	if err != nil {
		return err
	}

	clientKp := kinopoisk.NewClientKinopoisk(a.cfg.KPConsumerKey)
	tgBot := telegram.NewTgBot(a.bot, clientKp)
	if err = tgBot.HandleUpdate(updates); err != nil {
		return err
	}
	return nil
}

func (a *app) initUpdateChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return a.bot.GetUpdatesChan(u)
}
