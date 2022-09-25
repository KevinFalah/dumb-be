package transactionsdto

import (
	"dumbflix/models"
	"time"
)

type TransactionRequest struct {
	StartDate time.Time            		`json:"startDate"`
	DueDate	  time.Time     	   		`json:"dueDate"`
	UserID	 int				 	    `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price	  int						`json:"price"`
	User     models.UserResponse		`json:"user"`
}


type CreateTransactionRequest struct {
	StartDate string `json:"startdate"`
	DueDate   string `json:"duedate"`
	UserID    int    `json:"user_id" form:"user_id"`
	Attache   string `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    bool   `json:"status" gorm:"type:text" form:"status"`
}

type UpdateTransactionRequest struct {
	StartDate string `json:"startdate"`
	DueDate   string `json:"duedate"`
	UserID    int    `json:"user_id" form:"user_id"`
	Attache   string `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    bool   `json:"status" gorm:"type:text" form:"status"`
}
