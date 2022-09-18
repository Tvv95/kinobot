package dto

type Film struct {
	KinopoiskId     int       `json:"kinopoiskId"`
	NameRu          string    `json:"nameRu"`
	NameEn          string    `json:"nameEn"`
	NameOriginal    string    `json:"nameOriginal"`
	PosterUrl       string    `json:"posterUrl"`
	RatingKinopoisk float32   `json:"ratingKinopoisk"`
	RatingImdb      float32   `json:"RatingImdb"`
	Year            int       `json:"year"`
	FilmLength      int       `json:"filmLength"`
	Description     string    `json:"description"`
	Countries       []Country `json:"countries"`
	Genres          []Genre   `json:"genres"`
}

type Country struct {
	Country string `json:"country"`
}

type Genre struct {
	Genre string `json:"genre"`
}
