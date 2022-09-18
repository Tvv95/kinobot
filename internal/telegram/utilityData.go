package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const cmdStart = "start"

const (
	best      = "⭐ Лучшие"
	popular   = "🍿 Популярные"
	genres    = "🎬 Жанры"
	countries = "🌏 Страны"
	menu      = "Меню"
)

const (
	BestKey      = "TOP_250_BEST_FILMS"
	PopularKey   = "TOP_100_POPULAR_FILMS"
	GenresKey    = "genres"
	CountriesKey = "countries"
)

const (
	ussr    = "☭ СССР"
	russia  = "🇷🇺 Россия"
	usa     = "🇺🇸 США"
	france  = "🇫🇷 Франция"
	uk      = "🇬🇧 Великобритания"
	germany = "🇩🇪 Германия"
	italy   = "🇮🇹 Италия"
	denmark = "🇩🇰 Дания"
	india   = "🇮🇳 Индия"
	japan   = "🇯🇵 Япония"
	kr      = "🇰🇷 Южная Корея"
)

var countryToId = map[string]string{
	ussr:    "33",
	russia:  "34",
	usa:     "1",
	france:  "3",
	uk:      "5",
	germany: "9",
	italy:   "10",
	denmark: "17",
	india:   "7",
	japan:   "16",
	kr:      "49",
}

const (
	thriller    = "Триллер"
	drama       = "Драма"
	crime       = "Криминал"
	melodrama   = "Мелодрама"
	detective   = "Детектив"
	sciFi       = "Фантастика"
	adventure   = "Приключение"
	biography   = "Биография"
	noir        = "Нуар"
	western     = "Вестерн"
	action      = "Боевик"
	fantasy     = "Фэнтези"
	comedy      = "Комедия"
	war         = "Военный"
	history     = "История"
	music       = "Музыка"
	horror      = "Ужасы"
	cartoon     = "Мультфильм"
	family      = "Семейный"
	musical     = "Мюзикл"
	sport       = "Спорт"
	documentary = "Документальный"
	shortFilm   = "Корометражка"
	anime       = "Аниме"
)

var genreToId = map[string]string{
	thriller:    "1",
	drama:       "2",
	crime:       "3",
	melodrama:   "4",
	detective:   "5",
	sciFi:       "6",
	adventure:   "7",
	biography:   "8",
	noir:        "9",
	western:     "10",
	action:      "11",
	fantasy:     "12",
	comedy:      "13",
	war:         "14",
	history:     "15",
	music:       "16",
	horror:      "17",
	cartoon:     "18",
	family:      "19",
	musical:     "20",
	sport:       "21",
	documentary: "22",
	shortFilm:   "23",
	anime:       "24",
}

var keyboardButtonsMain = [][]tgbotapi.KeyboardButton{
	{tgbotapi.NewKeyboardButton(best), tgbotapi.NewKeyboardButton(popular)},
	{tgbotapi.NewKeyboardButton(genres), tgbotapi.NewKeyboardButton(countries)}}

var keyboardButtonsCountries = [][]tgbotapi.KeyboardButton{
	{tgbotapi.NewKeyboardButton(ussr), tgbotapi.NewKeyboardButton(russia), tgbotapi.NewKeyboardButton(usa)},
	{tgbotapi.NewKeyboardButton(france), tgbotapi.NewKeyboardButton(denmark), tgbotapi.NewKeyboardButton(germany)},
	{tgbotapi.NewKeyboardButton(italy), tgbotapi.NewKeyboardButton(japan), tgbotapi.NewKeyboardButton(india)},
	{tgbotapi.NewKeyboardButton(uk), tgbotapi.NewKeyboardButton(kr)},
	{tgbotapi.NewKeyboardButton(menu)}}

var keyboardButtonsGenres = [][]tgbotapi.KeyboardButton{
	{tgbotapi.NewKeyboardButton(thriller), tgbotapi.NewKeyboardButton(drama), tgbotapi.NewKeyboardButton(crime)},
	{tgbotapi.NewKeyboardButton(melodrama), tgbotapi.NewKeyboardButton(detective), tgbotapi.NewKeyboardButton(sciFi)},
	{tgbotapi.NewKeyboardButton(adventure), tgbotapi.NewKeyboardButton(biography), tgbotapi.NewKeyboardButton(noir)},
	{tgbotapi.NewKeyboardButton(western), tgbotapi.NewKeyboardButton(action), tgbotapi.NewKeyboardButton(fantasy)},
	{tgbotapi.NewKeyboardButton(comedy), tgbotapi.NewKeyboardButton(war), tgbotapi.NewKeyboardButton(history)},
	{tgbotapi.NewKeyboardButton(music), tgbotapi.NewKeyboardButton(horror), tgbotapi.NewKeyboardButton(cartoon)},
	{tgbotapi.NewKeyboardButton(family), tgbotapi.NewKeyboardButton(musical)},
	{tgbotapi.NewKeyboardButton(sport), tgbotapi.NewKeyboardButton(documentary)},
	{tgbotapi.NewKeyboardButton(shortFilm), tgbotapi.NewKeyboardButton(anime)},
	{tgbotapi.NewKeyboardButton(menu)}}
