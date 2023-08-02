package token

import (
	"testing"
	"time"

	"github.com/sembh1998/wedding_invitation_generation/internal/core/domain"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {

	maker, err := NewPasetoMaker("b4cbfb43df4ce210727d953e4a713307fa19bb7d9f85041438d9e11b942a37741eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2", "1eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	require.NoError(t, err)

	username := "sem"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(CreateTokenParams{
		Username: username,
		UserID:   domain.NewULIDString(),
		Duration: duration,
	})
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

}
