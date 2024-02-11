package invite_tokens

import (
	"crypto/rand"
	"math/big"
)

func GenerateRegisterToken(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string

	// Получаем длину набора символов
	charsetLength := big.NewInt(int64(len(charset)))

	// Генерируем случайные символы и добавляем их к результату
	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, charsetLength)
		result += string(charset[randomIndex.Int64()])
	}

	return result
}
