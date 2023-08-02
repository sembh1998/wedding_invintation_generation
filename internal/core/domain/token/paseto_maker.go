package token

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"

	"github.com/o1egl/paseto"
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto     *paseto.V2
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

func NewPasetoMaker(privateKey string, publicKey string) (Maker, error) {
	b, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	newPrivateKey := ed25519.PrivateKey(b)

	b, err = hex.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}
	newPublicKey := ed25519.PublicKey(b)

	maker := &PasetoMaker{
		paseto:     paseto.NewV2(),
		privateKey: newPrivateKey,
		publicKey:  newPublicKey,
	}
	return maker, nil
}

// CreateToken creates a new token for a specific username, privilage and duration
func (maker *PasetoMaker) CreateToken(params CreateTokenParams) (string, error) {
	payload, err := NewPayload(NewPayloadParams(params))
	if err != nil {
		return "", err
	}
	return maker.paseto.Sign(maker.privateKey, payload, nil)
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Verify(token, maker.publicKey, &payload, nil)

	if err != nil {
		return nil, fmt.Errorf("error verifying token: %w", err)
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
