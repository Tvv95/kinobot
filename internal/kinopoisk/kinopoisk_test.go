package kinopoisk_test

import (
	"github.com/stretchr/testify/assert"
	"kinobot/internal/dto"
	"kinobot/internal/kinopoisk"
	"kinobot/internal/telegram"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientKinopoisk_GetRequestCategoryList(t *testing.T) {
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
	expectedResponse := &dto.FilmCategoryList{
		TotalPages: 1,
		Items: []dto.FilmCategoryPreview{
			{
				KinopoiskId:     326,
				Year:            1994,
				NameRu:          "Побег из Шоушенка",
				NameEn:          "",
				NameOriginal:    "The Shawshank Redemption",
				RatingKinopoisk: 9.1,
			},
			{
				KinopoiskId:     435,
				Year:            1999,
				NameRu:          "Зеленая миля",
				NameEn:          "",
				NameOriginal:    "The Green Mile",
				RatingKinopoisk: 9.1,
			}},
	}
	clientKinopoisk := kinopoisk.NewClientKinopoisk("")
	clientKinopoisk.ApiURL = server.URL
	response, _ := clientKinopoisk.GetRequestCategoryList(telegram.GenresKey, "2", "1")

	assert.Equal(t, expectedResponse, &response)
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
	expectedResponse := &dto.FilmTopList{
		PagesCount: 1,
		Films: []dto.FilmTopPreview{
			{
				FilmId: 4642708,
				Year:   "2022",
				NameRu: "Нулевой пациент",
				NameEn: "",
				Rating: "8.4",
			},
			{
				FilmId: 915196,
				Year:   "2016",
				NameRu: "Очень странные дела",
				NameEn: "Stranger Things",
				Rating: "8.4",
			}},
	}
	defer server.Close()
	clientKinopoisk := kinopoisk.NewClientKinopoisk("")
	clientKinopoisk.ApiURL = server.URL
	response, _ := clientKinopoisk.GetRequestTopList(telegram.PopularKey, "1")

	assert.Equal(t, expectedResponse, &response)
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
	expectedResponse := &dto.Film{
		KinopoiskId:     325,
		NameRu:          "Крестный отец",
		NameEn:          "",
		NameOriginal:    "The Godfather",
		PosterUrl:       "https://kinopoiskapiunofficial.tech/images/posters/kp/325.jpg",
		RatingKinopoisk: 8.7,
		RatingImdb:      9.2,
		Year:            1972,
		FilmLength:      175,
		Description:     "Криминальная сага, повествующая о нью-йоркской сицилийской мафиозной семье Корлеоне. Фильм охватывает период 1945-1955 годов.\n\nГлава семьи, Дон Вито Корлеоне, выдаёт замуж свою дочь. В это время со Второй мировой войны возвращается его любимый сын Майкл. Майкл, герой войны, гордость семьи, не выражает желания заняться жестоким семейным бизнесом. Дон Корлеоне ведёт дела по старым правилам, но наступают иные времена, и появляются люди, желающие изменить сложившиеся порядки. На Дона Корлеоне совершается покушение.",
		Countries:       []dto.Country{{"США"}},
		Genres:          []dto.Genre{{"драма"}, {"криминал"}},
	}
	clientKinopoisk := kinopoisk.NewClientKinopoisk("")
	clientKinopoisk.ApiURL = server.URL
	response, _ := clientKinopoisk.GetRequestFilm("325")

	assert.Equal(t, expectedResponse, &response)
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
	expectedResponse := &dto.TrailerList{
		Items: []dto.Trailer{{"https://disk.yandex.ru/i/DamZvpQbqDdXvA"}},
	}
	clientKinopoisk := kinopoisk.NewClientKinopoisk("")
	clientKinopoisk.ApiURL = server.URL
	response, _ := clientKinopoisk.GetRequestTrailer("325")

	assert.Equal(t, expectedResponse, &response)
}
