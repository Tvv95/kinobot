package dto

type FilmTopList struct {
	PagesCount int              `json:"pagesCount"`
	Films      []FilmTopPreview `json:"films"`
}

type FilmTopPreview struct {
	FilmId int    `json:"filmId"`
	Year   string `json:"year"`
	NameRu string `json:"nameRu"`
	NameEn string `json:"nameEn"`
	Rating string `json:"rating"`
}
