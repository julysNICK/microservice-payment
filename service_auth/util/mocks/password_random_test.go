package mocks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomPassword(t *testing.T) {
	t.Run("should generate random password", func(t *testing.T) {
		password := GenerateRandomPassword()

		require.NotEmpty(t, password)
	})
}
