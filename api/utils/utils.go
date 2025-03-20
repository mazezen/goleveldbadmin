package utils

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum)
}
