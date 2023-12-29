package validations

import "regexp"

func removeNonNumeric(s string) string {
	reg, _ := regexp.Compile("[^0-9]+")

	return reg.ReplaceAllString(s, "")
}

func areAllDigitsEqual(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}

	return true
}

func calculateCPFCheckDigit(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i]-'0') * (len(s) + 1 - i)
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}

	return 11 - remainder
}

func IsValidCPF(cpf string) bool {
	cpf = removeNonNumeric(cpf)

	if len(cpf) != 11 || areAllDigitsEqual(cpf) {
		return false
	}

	if areAllDigitsEqual(cpf) {
		return false
	}

	digit1 := calculateCPFCheckDigit(cpf[:9])

	digit2 := calculateCPFCheckDigit(cpf[:9] + string(digit1+'0'))

	return cpf == cpf[:9]+string(digit1+'0')+string(digit2+'0')
}
