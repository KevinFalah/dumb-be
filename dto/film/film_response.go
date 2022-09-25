package filmdto

type FilmResponse struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailfillm" gorm:"type: varchar(255)"`
	Year          string `json:"year" gorm:"type: varchar(255)"`
	Description   string `json:"description" gorm:"type: text"`
	LinkFilm      string             `json:"linkfilm"  gorm:"type: text" form:"linkfilm"`
}
