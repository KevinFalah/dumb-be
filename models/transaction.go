package models

import "time"

type Transaction struct {
	ID        int                  		`json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time            		`json:"startDate"`
	DueDate	  time.Time     	   		`json:"dueDate"`
	UserID    int	                    `json:"-" form:"user_id"`
	User	  UserTransactionResponse 	`json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price	  int						`json:"price"`
	Status    string               		`json:"status" gorm:"type:varchar(50)"`
	CreatedAt time.Time            		`json:"-"`
	UpdatedAt time.Time            		`json:"-"`
}

type TransactionResponse struct {
	ID        int                  		`json:"id"`
	StartDate time.Time            		`json:"startDate"`
	DueDate	  time.Time     	   		`json:"dueDate"`
	UserID    int	                    `json:"-"`
	User	  UserTransactionResponse 	`json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price	  string					`json:"price"`
	Status    string               		`json:"status"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name" form:"name" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	Phone     int    `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	Subscribe string `json:"subscribe" form:"subscribe"`
	Status    string `json:"status" form:"status"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}

func (UserResponse) TableName() string {
	return "users"
}
