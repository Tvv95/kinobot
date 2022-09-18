package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (tb *tgBot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, "Используй меню ⬇")
	switch message.Text {
	case best:
		msg.Text, msg.ReplyMarkup = tb.buildTopAnswer(BestKey, "1")
	case popular:
		msg.Text, msg.ReplyMarkup = tb.buildTopAnswer(PopularKey, "1")
	case genres:
		msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{ResizeKeyboard: true, Keyboard: keyboardButtonsGenres}
		msg.Text = genres
	case countries:
		msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{ResizeKeyboard: true, Keyboard: keyboardButtonsCountries}
		msg.Text = countries
	case menu:
		msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{ResizeKeyboard: true, Keyboard: keyboardButtonsMain}
		msg.Text = menu
	case thriller,
		drama,
		crime,
		melodrama,
		detective,
		sciFi,
		adventure,
		biography,
		noir,
		western,
		action,
		fantasy,
		comedy,
		war,
		history,
		music,
		horror,
		cartoon,
		family,
		musical,
		sport,
		documentary,
		shortFilm,
		anime:
		msg.Text, msg.ReplyMarkup = tb.buildCategoryAnswer(GenresKey, genreToId[message.Text], "1")
	case ussr,
		russia,
		usa,
		france,
		uk,
		germany,
		italy,
		denmark,
		india,
		japan,
		kr:
		msg.Text, msg.ReplyMarkup = tb.buildCategoryAnswer(CountriesKey, countryToId[message.Text], "1")
	}
	if _, err := tb.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (tb *tgBot) handleCommand(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Command())
	msg := tgbotapi.NewMessage(message.Chat.ID, "Некорректная команда")
	if message.Command() == cmdStart {
		msg.Text = "Добро пожаловать!"
		keyboard := tgbotapi.ReplyKeyboardMarkup{ResizeKeyboard: true, Keyboard: keyboardButtonsMain}
		msg.ReplyMarkup = keyboard
		msg.ReplyToMessageID = message.MessageID
	} else if strings.HasPrefix(message.Command(), "id") {
		msg.ParseMode = "HTML"
		msg.Text = tb.buildFilmAnswer(message.Command()[2:])
		if trailerReplyMarkup := tb.buildTrailerAnswer(message.Command()[2:]); trailerReplyMarkup != nil {
			msg.ReplyMarkup = trailerReplyMarkup
		}
	}
	if _, err := tb.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (tb *tgBot) handleCallback(callback *tgbotapi.CallbackQuery) error {
	mainType, typeId, curPage := formatCallback(callback)
	var text string
	var replyMarkup tgbotapi.InlineKeyboardMarkup
	if mainType == BestKey || mainType == PopularKey {
		text, replyMarkup = tb.buildTopAnswer(mainType, curPage)
	} else if mainType == GenresKey || mainType == CountriesKey {
		text, replyMarkup = tb.buildCategoryAnswer(mainType, typeId, curPage)
	}
	edit := tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      callback.Message.Chat.ID,
			MessageID:   callback.Message.MessageID,
			ReplyMarkup: &replyMarkup,
		},
		Text: text,
	}

	if _, err := tb.bot.Send(edit); err != nil {
		log.Println(err)
	}
	return nil
}

func formatCallback(callback *tgbotapi.CallbackQuery) (string, string, string) {
	splitData := strings.Split(callback.Data, ":")
	curPage, err := strconv.Atoi(splitData[3])
	if err != nil {
		log.Println("Error format: ", err)
	}
	if splitData[0] == "next" {
		curPage++
	} else {
		curPage--
	}
	return splitData[1], splitData[2], strconv.Itoa(curPage)
}
