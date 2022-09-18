package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
	"kinobot/internal/kinopoisk"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTgBot_buildCategoryAnswer(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
    		"totalPages": 1,
    		"items": [
        	{
            	"kinopoiskId": 326,
            	"nameRu": "Побег из Шоушенка",
            	"nameEn": null,
            	"nameOriginal": "The Shawshank Redemption",
            	"ratingKinopoisk": 9.1,
            	"year": 1994
        	},
        	{
            	"kinopoiskId": 435,
            	"nameRu": "Зеленая миля",
            	"nameEn": null,
            	"nameOriginal": "The Green Mile",
            	"ratingKinopoisk": 9.1,
            	"year": 1999
        	}]
		}`))
	}))
	defer server.Close()
	expectedAnswer := "Побег из Шоушенка (1994) 9.1⭐ /id326\nЗеленая миля (1999) 9.1⭐ /id435\n"
	bot, _ := tgbotapi.NewBotAPI("")
	clientKinopoisk := kinopoisk.NewClientKinopoisk("")
	clientKinopoisk.ApiURL = server.URL
	tgBot := NewTgBot(bot, clientKinopoisk)
	answer, _ := tgBot.buildCategoryAnswer(GenresKey, "2", "1")

	assert.Equal(t, expectedAnswer, answer)
}

func TestClientKinopoisk_GetRequestTopList(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
    		"pagesCount": 1,
    		"films": [
        	{
            	"filmId": 4642708,
            	"nameRu": "Нулевой пациент",
            	"nameEn": null,
            	"year": "2022",
            	"rating": "8.4"
        	},
        	{
            	"filmId": 915196,
            	"nameRu": "Очень странные дела",
            	"nameEn": "Stranger Things",
            	"year": "2016",
            	"rating": "8.4"
        	}]
		}`))
	}))
	defer server.Close()
	expectedAnswer := "Нулевой пациент (2022) 8.4⭐ /id4642708\nОчень странные дела (2016) 8.4⭐ /id915196\n"
	bot, _ := tgbotapi.NewBotAPI("")
	clientKinopoisk := kinopoisk.NewClientKinopoisk("")
	clientKinopoisk.ApiURL = server.URL
	tgBot := NewTgBot(bot, clientKinopoisk)
	answer, _ := tgBot.buildTopAnswer(PopularKey, "1")

	assert.Equal(t, expectedAnswer, answer)
}

func TestClientKinopoisk_GetRequestFilm(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
    		"kinopoiskId": 325,
    		"nameRu": "Крестный отец",
    		"nameEn": null,
    		"nameOriginal": "The Godfather",
    		"posterUrl": "https://kinopoiskapiunofficial.tech/images/posters/kp/325.jpg",
    		"ratingKinopoisk": 8.7,
    		"ratingImdb": 9.2,
    		"year": 1972,
    		"filmLength": 175,
    		"description": "Криминальная сага, повествующая о нью-йоркской сицилийской мафиозной семье Корлеоне. Фильм охватывает период 1945-1955 годов.\n\nГлава семьи, Дон Вито Корлеоне, выдаёт замуж свою дочь. В это время со Второй мировой войны возвращается его любимый сын Майкл. Майкл, герой войны, гордость семьи, не выражает желания заняться жестоким семейным бизнесом. Дон Корлеоне ведёт дела по старым правилам, но наступают иные времена, и появляются люди, желающие изменить сложившиеся порядки. На Дона Корлеоне совершается покушение.",
    		"countries": [
        		{"country": "США"}
    		],
    		"genres": [
        		{"genre": "драма"},
        		{"genre": "криминал"}
    		]
		}`))
	}))
	defer server.Close()

	expectedAnswer := `<a href="https://kinopoiskapiunofficial.tech/images/posters/kp/325.jpg">&#160;</a>
<b>Крестный отец (The Godfather)</b>
Год производства: 1972
Время: 02:55
Страна: США
Жанр: драма, криминал
Рейтинг Кинопоиск: 8.7⭐
Рейтинг IMDb: 9.2⭐

Криминальная сага, повествующая о нью-йоркской сицилийской мафиозной семье Корлеоне. Фильм охватывает период 1945-1955 годов.

Глава семьи, Дон Вито Корлеоне, выдаёт замуж свою дочь. В это время со Второй мировой войны возвращается его любимый сын Майкл. Майкл, герой войны, гордость семьи, не выражает желания заняться жестоким семейным бизнесом. Дон Корлеоне ведёт дела по старым правилам, но наступают иные времена, и появляются люди, желающие изменить сложившиеся порядки. На Дона Корлеоне совершается покушение.`

	bot, _ := tgbotapi.NewBotAPI("")
	clientKinopoisk := kinopoisk.NewClientKinopoisk("")
	clientKinopoisk.ApiURL = server.URL
	tgBot := NewTgBot(bot, clientKinopoisk)
	answer := tgBot.buildFilmAnswer("325")

	assert.Equal(t, expectedAnswer, answer)
}

func TestClientKinopoisk_GetRequestTrailer(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
    		"items": [
        		{"url": "https://disk.yandex.ru/i/DamZvpQbqDdXvA"}
			]
		}`))
	}))
	defer server.Close()
	expectedAnswer := tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonURL("Трейлер", "https://disk.yandex.ru/i/DamZvpQbqDdXvA")})
	bot, _ := tgbotapi.NewBotAPI("")
	clientKinopoisk := kinopoisk.NewClientKinopoisk("")
	clientKinopoisk.ApiURL = server.URL
	tgBot := NewTgBot(bot, clientKinopoisk)
	answer := tgBot.buildTrailerAnswer("325")

	assert.Equal(t, expectedAnswer, *answer)
}

func TestFormatTime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		arg      int
		expected string
	}{
		{
			name:     "10",
			arg:      10,
			expected: "00:10",
		},
		{
			name:     "100",
			arg:      100,
			expected: "01:40",
		},
		{
			name:     "250",
			arg:      250,
			expected: "04:10",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, formatTime(test.arg))
		})
	}
}
