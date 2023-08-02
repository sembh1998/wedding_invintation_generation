package userssrv

import (
	"time"

	"github.com/sembh1998/wedding_invitation_generation/internal/core/domain/password"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/domain/token"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/ports"
)

type service struct {
	userRepo ports.UserRepo
}

func New(userRepo ports.UserRepo) ports.UserSrv {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Login(username string, pass string) (string, error) {
	user, err := s.userRepo.FindUser(username)
	if err != nil {
		return "", err
	}
	err = password.CheckPassword(pass, user.Password)
	if err != nil {
		return "", err
	}

	token, err := token.NewTokenizer().Maker.CreateToken(token.CreateTokenParams{
		Username: username,
		UserID:   user.ID,
		Duration: 24 * time.Hour,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}
