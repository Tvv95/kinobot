package dto

type FilmCategoryList struct {
	TotalPages int                   `json:"totalPages"`
	Items      []FilmCategoryPreview `json:"items"`
}

type FilmCategoryPreview struct {
	KinopoiskId     int     `json:"kinopoiskId"`
	Year            int     `json:"year"`
	NameRu          string  `json:"nameRu"`
	NameEn          string  `json:"nameEn"`
	NameOriginal    string  `json:"nameOriginal"`
	RatingKinopoisk float32 `json:"ratingKinopoisk"`
}
