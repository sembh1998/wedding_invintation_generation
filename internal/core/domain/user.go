package domain

type User struct {
	BaseModel
	User     string `json:"user"`
	Password string `json:"password"`
}
