package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Fullname      string    `json:"fullname" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	Phone     string    `json:"phone" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Role      string    `json:"role" gorm:"type: varchar(255)"`
	Subscribe bool                  `json:"subscribe"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserTransactionResponse struct {
	ID        int                   `json:"id"`
	Fullname  string                `json:"fullname" gorm:"type: varchar(255)"`
	Email     string                `json:"email" gorm:"type: varchar(255)"`
	Gender    string                `json:"gender" gorm:"type: varchar(100)"`
	Phone     string                `json:"phone" gorm:"type: varchar(255)"`
	Address   string                `json:"address" gorm:"type: text"`
	Subscribe bool                  `json:"subscribe"`
  }

  func (UserTransactionResponse) TableName() string {
	return "users"
  }