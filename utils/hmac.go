package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacCheck(message, key []byte, messageMAC string) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	m, err := hex.DecodeString(messageMAC)
	if err != nil {
		return false
	}
	return hmac.Equal(m, expectedMAC)
}

func HmacSign(message, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return hex.EncodeToString(mac.Sum(nil))
}
