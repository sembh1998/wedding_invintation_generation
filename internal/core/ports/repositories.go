package ports

import "github.com/sembh1998/wedding_invitation_generation/internal/core/domain"

type UserRepo interface {
	FindUser(user string) (domain.User, error)
}

type GuestRepo interface {
	FindGuests() ([]domain.Guest, error)
	FetchGuest(id string) (domain.Guest, error)
	CreateGuest(guest domain.Guest) (domain.Guest, error)
	UpdateGuest(guest domain.Guest) (domain.Guest, error)
	DeleteGuest(id string) error
}
