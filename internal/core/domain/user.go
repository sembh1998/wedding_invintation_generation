package domain

type User struct {
	BaseModel
	ID       string `json:"id" gorm:"primary_key;type:varchar(36);column:id" form:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
}
