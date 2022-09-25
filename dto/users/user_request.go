package usersdto

type CreateUserRequest struct {
	Fullname      string `json:"fullname" form:"name" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Gender    string `json:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: varchar(255)"`
	Subscribe bool   `json:"subscribe" gorm:"type: boolean"`
	Role      string `json:"role" gorm:"type: varchar(255)"`
}

type UpdateUserRequest struct {
	Fullname      string `json:"fullname" form:"name" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Gender    string `json:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: varchar(255)"`
	Subscribe bool   `json:"subscribe" gorm:"type: boolean"`
	Role      string `json:"role" gorm:"type: varchar(255)"`
}
