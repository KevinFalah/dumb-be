package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	AddProfile(profile models.Profile) (models.Profile, error)
	GetProfile(ID int) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddProfile(profile models.Profile) (models.Profile, error) {
	err := r.db.Create(&profile).Error

	return profile, err
}

func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").First(&profile, "user_id=?", ID).Error // memanggil user di models yang sudah melakukan relasi dengan profile

	return profile, err
}
