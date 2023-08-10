package guestsrepo

import (
	"github.com/sembh1998/wedding_invitation_generation/internal/core/domain"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/ports"
	"gorm.io/gorm"
)

type GuestMysql struct {
	// gorm connection
	Conn *gorm.DB
}

func NewGuestMysql(conn *gorm.DB) ports.GuestRepo {
	conn.AutoMigrate(&domain.Guest{})
	return &GuestMysql{
		Conn: conn,
	}
}

func (u *GuestMysql) FindGuests() ([]domain.Guest, error) {
	var guests []domain.Guest
	err := u.Conn.Joins("User").Order("created_at DESC").Find(&guests).Error
	if err != nil {
		return []domain.Guest{}, err
	}
	return guests, nil
}

func (u *GuestMysql) FetchGuest(id string) (domain.Guest, error) {
	var guest domain.Guest
	err := u.Conn.Joins("User").First(&guest, "guests.id = ?", id).Error
	if err != nil {
		return domain.Guest{}, err
	}
	return guest, nil
}

func (u *GuestMysql) CreateGuest(guest domain.Guest) (domain.Guest, error) {
	err := u.Conn.Create(&guest).Error
	if err != nil {
		return domain.Guest{}, err
	}
	return guest, nil
}

func (u *GuestMysql) UpdateGuest(guest domain.Guest) (domain.Guest, error) {
	err := u.Conn.Save(&guest).Error
	if err != nil {
		return domain.Guest{}, err
	}
	return guest, nil
}

func (u *GuestMysql) DeleteGuest(id string) error {
	err := u.Conn.Delete(&domain.Guest{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
