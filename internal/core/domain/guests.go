package domain

type Guest struct {
	BaseModel
	Name           string `form:"name"`
	LastName       string `form:"last_name"`
	WillAttend     int    `form:"will_attend"`
	SpecialMessage string `gorm:"type:text" form:"special_message"`
	Response       string `gorm:"type:text" form:"response"`
	// foriegn key to user
	CreatedBy string
	User      User `gorm:"foreignKey:CreatedBy"`
}
