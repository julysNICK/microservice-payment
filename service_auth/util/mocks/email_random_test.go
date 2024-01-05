package mocks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	t.Run("should generate random string", func(t *testing.T) {
		randomString := randomString()

		require.NotEmpty(t, randomString)
	})
}

func TestRandomEmail(t *testing.T) {
	t.Run("should generate random email", func(t *testing.T) {
		email := RandomEmail()

		require.Contains(t, email, "@gmail.com")
	})
}
