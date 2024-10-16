package services

import (
	"crypto/rand"
	"math/big"
)

func randomPasteIdPrefix(length uint8) string {
	// Define the character set
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLen := big.NewInt(int64(len(charset)))

	// Create a byte slice of the specified length
	randomString := make([]byte, length)

	// Fill the byte slice with random characters from the character set
	for i := range randomString {
		// Generate a secure random index
		num, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			// Handle the error in case of failure in generating a random number
			panic(err)
		}
		randomString[i] = charset[num.Int64()]
	}

	// Convert the byte slice to a string and return it
	return string(randomString)
}
