package guestssrv

import (
	"fmt"

	"github.com/sembh1998/wedding_invitation_generation/internal/core/domain"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/ports"
)

type service struct {
	guestRepo ports.GuestRepo
}

func New(guestRepo ports.GuestRepo) ports.GuestSrv {
	return &service{
		guestRepo: guestRepo,
	}
}

func (s *service) CreateGuest(entry domain.Guest) (domain.Guest, error) {
	entry.ID = domain.NewULIDString()
	guest, err := s.guestRepo.CreateGuest(entry)
	if err != nil {
		return domain.Guest{}, fmt.Errorf("error creating guest: %v", err)
	}
	return guest, nil
}

func (s *service) DeleteGuest(id string) error {
	return s.guestRepo.DeleteGuest(id)
}

func (s *service) FindGuests() ([]domain.Guest, error) {
	guests, err := s.guestRepo.FindGuests()
	if err != nil {
		return []domain.Guest{}, fmt.Errorf("error finding guests: %v", err)
	}
	return guests, nil
}

func (s *service) UpdateGuest(entry domain.Guest) (domain.Guest, error) {
	guest, err := s.guestRepo.UpdateGuest(entry)
	if err != nil {
		return domain.Guest{}, fmt.Errorf("error updating guest: %v", err)
	}
	return guest, nil
}
