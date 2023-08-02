package domain

type Guest struct {
	BaseModel
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	WillAttend     bool   `json:"will_attend"`
	SpecialMessage string `json:"special_message" gorm:"type:text"`
	Response       string `json:"response" gorm:"type:text"`
}
