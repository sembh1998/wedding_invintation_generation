package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}

type View struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

// Payload contanis the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	UserID    string    `json:"user_id"`
	Issuer    string    `json:"issuer"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type NewPayloadParams struct {
	Username string
	Duration time.Duration
	UserID   string
}

// NewPayload chreates a new token payload with a specific username, privilage and duration
func NewPayload(params NewPayloadParams) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  params.Username,
		Issuer:    "wedding_invitation_generation",
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(params.Duration),
		UserID:    params.UserID,
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return fmt.Errorf("token is expired")
	}
	return nil
}
