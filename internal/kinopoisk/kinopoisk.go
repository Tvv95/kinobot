package kinopoisk

import (
	"encoding/json"
	"fmt"
	"io"
	"kinobot/internal/dto"
	"log"
	"net/http"
)

const defaultApiURL = "https://kinopoiskapiunofficial.tech/api/v2.2/films"

type ClientKinopoisk struct {
	consumerKey string
	ApiURL      string
}

func NewClientKinopoisk(consumerKey string) *ClientKinopoisk {
	return &ClientKinopoisk{consumerKey: consumerKey, ApiURL: defaultApiURL}
}

func (ck *ClientKinopoisk) GetRequestTopList(topType string, pageNumber string) (dto.FilmTopList, error) {
	url := fmt.Sprintf("%s/top?type=%s&page=%s", ck.ApiURL, topType, pageNumber)
	jsonData, err := ck.requestData(url)
	if err != nil {
		return dto.FilmTopList{}, err
	}
	return unmarshalJson(jsonData, dto.FilmTopList{})
}

func (ck *ClientKinopoisk) GetRequestFilm(id string) (dto.Film, error) {
	url := fmt.Sprintf("%s/%s", ck.ApiURL, id)
	jsonData, err := ck.requestData(url)
	if err != nil {
		return dto.Film{}, err
	}
	return unmarshalJson(jsonData, dto.Film{})
}

func (ck *ClientKinopoisk) GetRequestTrailer(id string) (dto.TrailerList, error) {
	url := fmt.Sprintf("%s/%s/videos", ck.ApiURL, id)
	jsonData, err := ck.requestData(url)
	if err != nil {
		return dto.TrailerList{}, err
	}
	return unmarshalJson(jsonData, dto.TrailerList{})
}

func (ck *ClientKinopoisk) GetRequestCategoryList(categoryType string, typeId string, pageNumber string) (dto.FilmCategoryList, error) {
	url := fmt.Sprintf("%s?%s=%s&order=RATING&type=ALL&ratingFrom=5&ratingTo=10&yearFrom=1000&yearTo=3000&page=%s", ck.ApiURL, categoryType, typeId, pageNumber)
	jsonData, err := ck.requestData(url)
	if err != nil {
		return dto.FilmCategoryList{}, err
	}
	return unmarshalJson(jsonData, dto.FilmCategoryList{})
}

func (ck *ClientKinopoisk) requestData(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-API-KEY", ck.consumerKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error:", err)
		return nil, err
	}
	return jsonData, nil
}

type dtoConstraint interface {
	dto.Film | dto.FilmCategoryList | dto.FilmTopList | dto.TrailerList
}

func unmarshalJson[T dtoConstraint](jsonData []byte, data T) (T, error) {
	if errJson := json.Unmarshal(jsonData, &data); errJson != nil {
		log.Println("Unmarshal error:", errJson)
		return *new(T), errJson
	}
	return data, nil
}
