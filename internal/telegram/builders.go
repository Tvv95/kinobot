package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
	"time"
)

func (tb *tgBot) buildTopAnswer(mainType, curPage string) (string, tgbotapi.InlineKeyboardMarkup) {
	filmsData, err := tb.clientKinopoisk.GetRequestTopList(mainType, curPage)
	if err != nil {
		return "", tgbotapi.InlineKeyboardMarkup{}
	}
	var sb strings.Builder
	for _, v := range filmsData.Films {
		name := v.NameRu
		if name == "" {
			name = v.NameEn
		}
		sb.WriteString(fmt.Sprintf("%s (%s) %s⭐ /id%d\n", v.NameRu, v.Year, v.Rating, v.FilmId))
	}
	return sb.String(), NewInlinePaginator(mainType, "", curPage, filmsData.PagesCount)
}

func (tb *tgBot) buildCategoryAnswer(mainType, typeId, curPage string) (string, tgbotapi.InlineKeyboardMarkup) {
	filmsData, err := tb.clientKinopoisk.GetRequestCategoryList(mainType, typeId, curPage)
	if err != nil {
		return "", tgbotapi.InlineKeyboardMarkup{}
	}
	var sb strings.Builder
	for _, v := range filmsData.Items {
		name := v.NameRu
		if name == "" && v.NameEn != "" {
			name = v.NameEn
		} else if name == "" {
			name = v.NameOriginal
		}
		sb.WriteString(fmt.Sprintf("%s (%d) %.1f⭐ /id%d\n", name, v.Year, v.RatingKinopoisk, v.KinopoiskId))
	}
	return sb.String(), NewInlinePaginator(mainType, typeId, curPage, filmsData.TotalPages)
}

func (tb *tgBot) buildFilmAnswer(id string) string {
	filmData, err := tb.clientKinopoisk.GetRequestFilm(id)
	if err != nil {
		return ""
	}
	var sb strings.Builder
	if filmData.PosterUrl != "" {
		sb.WriteString(fmt.Sprintf("<a href=\"%s\">&#160;</a>\n", filmData.PosterUrl))
	}
	name := filmData.NameRu
	var origName string
	if name == "" && filmData.NameEn != "" {
		name = filmData.NameEn
	} else if name == "" {
		name = filmData.NameOriginal
	}
	if name != filmData.NameOriginal && filmData.NameOriginal != "" {
		origName = fmt.Sprintf(" (%s)", filmData.NameOriginal)
	}
	sb.WriteString(fmt.Sprintf("<b>%s%s</b>", name, origName))
	if filmData.Year != 0 {
		sb.WriteString(fmt.Sprintf("\nГод производства: %d", filmData.Year))
	}
	if filmData.FilmLength != 0 {
		sb.WriteString(fmt.Sprintf("\nВремя: %s", formatTime(filmData.FilmLength)))
	}
	if len(filmData.Countries) != 0 {
		sb.WriteString("\nСтрана: ")
	}
	for i, country := range filmData.Countries {
		if i == len(filmData.Countries)-1 {
			sb.WriteString(country.Country)
			break
		}
		sb.WriteString(fmt.Sprintf("%s, ", country.Country))
	}
	if len(filmData.Genres) != 0 {
		sb.WriteString("\nЖанр: ")
	}
	for i, genre := range filmData.Genres {
		if i == len(filmData.Genres)-1 {
			sb.WriteString(genre.Genre)
			break
		}
		sb.WriteString(fmt.Sprintf("%s, ", genre.Genre))
	}
	if filmData.RatingKinopoisk != 0 {
		sb.WriteString(fmt.Sprintf("\nРейтинг Кинопоиск: %.1f⭐", filmData.RatingKinopoisk))
	}
	if filmData.RatingImdb != 0 {
		sb.WriteString(fmt.Sprintf("\nРейтинг IMDb: %.1f⭐", filmData.RatingImdb))
	}
	if filmData.Description != "" {
		sb.WriteString(fmt.Sprintf("\n\n%s", filmData.Description))
	}
	return sb.String()
}

func (tb *tgBot) buildTrailerAnswer(id string) *tgbotapi.InlineKeyboardMarkup {
	trailerData, err := tb.clientKinopoisk.GetRequestTrailer(id)
	if err != nil || len(trailerData.Items) == 0 {
		return nil
	}
	trailerInlineKeyboard := newInlineKeyboardTrailer(trailerData.Items[0].Url)
	return &trailerInlineKeyboard
}

func NewInlinePaginator(mainType, typeId, curPage string, maxPage int) tgbotapi.InlineKeyboardMarkup {
	var buttons []tgbotapi.InlineKeyboardButton
	currentPage, err := strconv.Atoi(curPage)
	if err != nil {
		log.Println(err)
	}
	if currentPage != 1 {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Назад", fmt.Sprintf("prev:%s:%s:%d", mainType, typeId, currentPage)))
	}
	if currentPage < maxPage {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Вперёд", fmt.Sprintf("next:%s:%s:%d", mainType, typeId, currentPage)))
	}
	return tgbotapi.NewInlineKeyboardMarkup(buttons)
}

func newInlineKeyboardTrailer(urlTrailer string) tgbotapi.InlineKeyboardMarkup {
	btn := tgbotapi.NewInlineKeyboardButtonURL("Трейлер", urlTrailer)
	return tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{btn})
}

func formatTime(duration int) string {
	dur := time.Duration(duration) * time.Minute
	h := dur / time.Hour
	dur -= h * time.Hour
	m := dur / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}
