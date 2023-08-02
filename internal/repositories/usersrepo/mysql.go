package usersrepo

import (
	"github.com/sembh1998/wedding_invitation_generation/internal/core/domain"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/ports"
	"gorm.io/gorm"
)

type UserMysql struct {
	// gorm connection
	Conn *gorm.DB
}

func NewUserMysql(conn *gorm.DB) ports.UserRepo {
	conn.AutoMigrate(&domain.User{})
	return &UserMysql{
		Conn: conn,
	}
}

func (u *UserMysql) FindUser(user string) (domain.User, error) {
	var usr domain.User
	err := u.Conn.First(&usr, "user = ?", user).Error
	if err != nil {
		return domain.User{}, err
	}
	return usr, nil
}
