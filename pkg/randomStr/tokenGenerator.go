package randomStr

import (
	"crypto/rand"
	"math/big"
)

func GenerateRegisterToken() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 32

	var result string

	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, charsetLength)
		result += string(charset[randomIndex.Int64()])
	}

	return result
}

func GenerateRefreshToken() string {
	const length = 64
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string

	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, charsetLength)
		result += string(charset[randomIndex.Int64()])
	}

	return result
}
