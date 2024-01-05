package mocks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormtatCpf(t *testing.T) {
	t.Run("should format cpf", func(t *testing.T) {
		cpf := formatCPF([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1})

		require.Equal(t, "12345678911", cpf)
	})
}

func TestCalculateCPFCheckDigit(t *testing.T) {
	t.Run("should calculate cpf check digit", func(t *testing.T) {
		digit := calculateCPFCheckDigit([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

		require.Equal(t, 0, digit)
	})
}

func TestGenerateRandomCPF(t *testing.T) {
	t.Run("should generate random cpf", func(t *testing.T) {
		cpf := GenerateRandomCPF()

		require.NotEmpty(t, cpf)
	})
}
