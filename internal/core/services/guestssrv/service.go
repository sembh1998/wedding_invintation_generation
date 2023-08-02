package guestssrv

import (
	"fmt"
	"time"

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
	entry.CreatedAt = time.Now()
	entry.UpdatedAt = time.Now()
	entry.WillAttend = 99
	guest, err := s.guestRepo.CreateGuest(entry)
	if err != nil {
		return domain.Guest{}, fmt.Errorf("error creating guest: %v", err)
	}
	return guest, nil
}

func (s *service) FetchGuest(id string) (domain.Guest, error) {
	guest, err := s.guestRepo.FetchGuest(id)
	if err != nil {
		return domain.Guest{}, fmt.Errorf("error fetching guest: %v", err)
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
	guest_before, err := s.guestRepo.FetchGuest(entry.ID)
	if err != nil {
		return domain.Guest{}, fmt.Errorf("error fetching guest: %v", err)
	}
	guest_before.UpdatedAt = time.Now()
	guest_before.Name = entry.Name
	guest_before.LastName = entry.LastName
	guest_before.SpecialMessage = entry.SpecialMessage

	guest, err := s.guestRepo.UpdateGuest(guest_before)
	if err != nil {
		return domain.Guest{}, fmt.Errorf("error updating guest: %v", err)
	}
	return guest, nil
}

func (s *service) AttendConfirmation(id string, attend bool, response string) (domain.Guest, error) {
	guest, err := s.guestRepo.FetchGuest(id)
	if err != nil {
		return domain.Guest{}, fmt.Errorf("error fetching guest: %v", err)
	}
	if attend {
		guest.WillAttend = 1
	} else {
		guest.WillAttend = 0
	}
	guest.Response = response
	guest, err = s.guestRepo.UpdateGuest(guest)
	if err != nil {
		return domain.Guest{}, fmt.Errorf("error updating guest: %v", err)
	}
	return guest, nil
}
