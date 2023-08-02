package ports

import "github.com/sembh1998/wedding_invitation_generation/internal/core/domain"

type UserSrv interface {
	Login(user, password string) (string, error)
}

type GuestSrv interface {
	FindGuests() ([]domain.Guest, error)
	FetchGuest(id string) (domain.Guest, error)
	CreateGuest(guest domain.Guest) (domain.Guest, error)
	UpdateGuest(guest domain.Guest) (domain.Guest, error)
	AttendConfirmation(id string, attend bool, response string) (domain.Guest, error)
	DeleteGuest(id string) error
}
