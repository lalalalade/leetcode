package codetop4

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	if s == "" {
		return s
	}

	if unicode.IsLetter(rune(s[0])) {
		return s[:1] + decodeString(s[1:])
	}

	i := strings.IndexByte(s, '[')
	balance := 1
	for j := i + 1; ; j++ {
		if s[j] == '[' {
			balance++
		} else if s[j] == ']' {
			balance--
			if balance == 0 {
				k, _ := strconv.Atoi(s[:i])
				return strings.Repeat(decodeString(s[i+1:j]), k) + decodeString(s[j+1:])
			}
		}
	}
}
