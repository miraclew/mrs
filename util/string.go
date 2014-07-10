package util

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

func MakeRandomString(strength int) string {
	s := make([]byte, strength)
	if _, err := rand.Read(s); err != nil {
		return ""
	}

	str := base64.URLEncoding.EncodeToString(s)
	str = strings.Replace(str, "+", "", -1)
	str = strings.Replace(str, "/", "", -1)
	str = strings.Replace(str, "=", "", -1)
	str = strings.Replace(str, "-", "", -1)

	return Substr(str, 0, strength)
}

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
