package mocks

import "math/rand"

func randomString() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	randomString := make([]byte, 10)

	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(letters))]
	}

	return string(randomString)

}

func RandomEmail() string {

	return randomString() + "@gmail.com"

}
