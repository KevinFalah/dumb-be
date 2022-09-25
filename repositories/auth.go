package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

// untuk export function (repository)
type AuthRepository interface {
	Register(user models.User) (models.User, error)
	RegisterUpdateAuth(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	GetAllUsers() ([]models.User, error)
	Getuser(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) RegisterUpdateAuth(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Profile").Find(&users).Error

	return users, err
}

func (r *repository) Getuser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
