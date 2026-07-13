package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateCode(url string) string {
	hash := sha256.Sum256([]byte(url))
	return hex.EncodeToString(hash[:])[:10]
}
