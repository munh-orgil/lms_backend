package helpers

import (
	"crypto/sha1"
	"encoding/hex"
	"math"
	"math/rand"
	"strings"
	"time"
)

func GeneratePassword(password string) string {
	sha1pwd := sha1.Sum([]byte(strings.ToUpper(password)))
	return strings.ToUpper(hex.EncodeToString(sha1pwd[:]))
}

func GenerateRandom(len int) uint {
	min := int(math.Pow10(len - 1))
	max := int(math.Pow10(len) - 1)
	rand.NewSource(time.Now().UnixNano())
	return uint(rand.Intn(max-min) + min)
}
