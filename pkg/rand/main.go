package rand

import "math/rand"

func GenerateBase58String(length int) string {
	base64Chars := "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		randPosition := rand.Intn(len(base64Chars))
		result[i] = base64Chars[randPosition]
	}
	return string(result)
}
