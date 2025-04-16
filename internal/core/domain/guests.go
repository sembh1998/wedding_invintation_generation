package domain

type Guest struct {
	BaseModel
	ID              string `json:"id" form:"id"`
	Order           int    `form:"order"  gorm:"primary_key;column:order"`
	Name            string `form:"name"`
	LastName        string `form:"last_name"`
	WillAttend      int    `form:"will_attend"`
	SpecialMessage  string `gorm:"type:text" form:"special_message"`
	Response        string `gorm:"type:text" form:"response"`
	AttendanceLimit int    `form:"attendance_limit"`
	// foriegn key to user
	CreatedBy string
	User      User `gorm:"foreignKey:CreatedBy"`
}
