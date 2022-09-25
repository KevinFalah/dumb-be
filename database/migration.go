package database

import (
	"dumbflix/models"
	"dumbflix/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Film{},
		&models.Profile{},
		&models.Category{},
		&models.Transaction{},
		&models.Episode{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
