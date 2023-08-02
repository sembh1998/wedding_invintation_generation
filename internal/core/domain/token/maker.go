package token

import "time"

type CreateTokenParams struct {
	Username string
	Duration time.Duration
	UserID   string
}
type Maker interface {
	CreateToken(params CreateTokenParams) (string, error)

	VerifyToken(token string) (*Payload, error)
}
