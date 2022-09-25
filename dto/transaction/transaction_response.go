package transactionsdto

type TransactionResponse struct {
	ID        int    `json:"id"`
	StartDate string `json:"startdate"`
	DueDate   string `json:"duedate"`
	UserID    int    `json:"user_id" form:"user_id"`
	User      UserResponse `json:"user"`
	Attache   string `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    bool   `json:"status" gorm:"type:text" form:"status"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Fullname      string `json:"fullname" form:"name" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	Phone     int    `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	Subscribe string `json:"subscribe" form:"subscribe"`
	Status    string `json:"status" form:"status"`
}

func (UserResponse) TableName() string {
	return "users"
}
