package filmdto

type CreateFilmRequest struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"  validate:"required"`
	ThumbnailFilm string `json:"thumbnailfilm" form:"image" gorm:"type: varchar(255)" validate:"required"`
	Year          string `json:"year" form:"year" gorm:"type: varchar(255)" validate:"required"`
	CategoryID    int    `json:"category_id" form:"category_id" gorm:"type: int" validate:"required"`
	Description   string `json:"description" form:"description" gorm:"type: text" validate:"required"`
	LinkFilm      string             `json:"linkfilm"  gorm:"type: text" form:"linkfilm"`

}

type UpdateFilmRequest struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year          string `json:"year" gorm:"type: varchar(255)"`
	CategoryID    int    `json:"category_id" gorm:"type: int"`
	Description   string `json:"description" gorm:"type: text"`
	LinkFilm      string             `json:"linkfilm"  gorm:"type: text" form:"linkfilm"`

}
