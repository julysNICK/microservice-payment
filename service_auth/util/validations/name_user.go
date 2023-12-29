package validations

func validateName(name string) bool {
	return len(name) > 0
}

func ValidateName(name string) bool {
	return validateName(name)
}
