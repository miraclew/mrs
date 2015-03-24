package util

import (
	"crypto/rand"
	"encoding/base64"
	"sort"
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

func RemoveDuplicates(xs *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

func SplitUniqSort(s string) []string {
	members := strings.Split(s, ",")
	var newMembers []string
	for i := 0; i < len(members); i++ {
		m := members[i]
		m = strings.TrimSpace(m)
		if len(m) <= 0 {
			continue
		}
		newMembers = append(newMembers, m)
	}

	RemoveDuplicates(&newMembers)
	sorted := sort.StringSlice(newMembers)
	sorted.Sort()
	return sorted
}
