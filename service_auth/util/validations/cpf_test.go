package validations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidCPF(t *testing.T) {
	validCpf := IsValidCPF("12345678911")

	require.True(t, validCpf)

}
