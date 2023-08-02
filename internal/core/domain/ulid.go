package domain

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func NewULID() ulid.ULID {
	entropy := rand.New(rand.NewSource(int64(rand.Int())))
	ms := ulid.Timestamp(time.Now())
	new_ulid, _ := ulid.New(ms, entropy)
	return new_ulid
}

func NewULIDString() string {
	return NewULID().String()
}
