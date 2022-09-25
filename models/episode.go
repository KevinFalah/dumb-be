package models

type Episode struct {
	ID               int                  `json:"id"`
	Title            string               `json:"title" gorm:"type: varchar(225)"`
	ThumbnailEpisode string               `json:"thumbnailepisode" gorm:"type: varchar(50)"`
	LinkFilm         string               `json:"linkFilm" gorm:"type: varchar(50)"`
	FilmID           int                  `json:"film_id" form:"film_id"`
	Film             FilmCategoryResponse `json:"film"`
	Description      string               `json:"description" gorm:"type :text"`
}

type EpisodeResponse struct {
	ID               int                  `json:"id"`
	Title            string               `json:"title"`
	ThumbnailEpisode string               `json:"thumbnailepisode" `
	LinkFilm         string               `json:"linkFilm"`
	FilmID           int                  `json:"film_id"`
	Film             FilmCategoryResponse `json:"film"`
	Description      string
}

func (EpisodeResponse) TableName() string {
	return "episode"
}
