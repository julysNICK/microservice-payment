package validations

func isValidCardNumber(cardNumber string) bool {
	return len(cardNumber) == 16
}

func isValidCardSecurityCode(cardSecurityCode string) bool {
	return len(cardSecurityCode) == 3
}

func isValidCardExpirationDate(cardExpirationDate string) bool {
	return len(cardExpirationDate) == 4
}

func isLuhnValid(cardNumber string) bool {
	sum := 0
	alternate := false
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')
		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}

func IsValidCreditCard(cardNumber string) bool {
	return isValidCardNumber(cardNumber) && isLuhnValid(cardNumber)
}
