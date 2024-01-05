package mocks

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomCPF() string {
	rand.Seed(time.Now().UnixNano())

	digits := make([]int, 9)
	for i := range digits {
		digits[i] = rand.Intn(10)
	}

	d1 := calculateCPFCheckDigit(digits)
	digits = append(digits, d1)
	d2 := calculateCPFCheckDigit(digits)
	digits = append(digits, d2)

	return formatCPF(digits)
}

func calculateCPFCheckDigit(digits []int) int {
	sum := 0
	for i, digit := range digits {
		sum += digit * (10 - i)
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

func formatCPF(digits []int) string {
	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d", digits[0], digits[1], digits[2], digits[3], digits[4], digits[5], digits[6], digits[7], digits[8], digits[9], digits[10])
}
