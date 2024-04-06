package utils

import (
	"strings"
	"unicode"
)

func ToLowerCamelCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
func ToUpperCamelCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func SnakeToCamel(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] == '_' {
			if i+1 < len(runes) {
				runes[i+1] = unicode.ToUpper(runes[i+1])
			}
		}
	}
	return strings.ReplaceAll(string(runes), "_", "")
}

func UpperCamelToSnake(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) {
			runes[i] = unicode.ToLower(runes[i])
			runes = append(runes[:i], append([]rune{'_'}, runes[i:]...)...)
			i++
		}
	}
	if after, found := strings.CutPrefix((string(runes)), "_"); found {
		return after
	}
	return string(runes)
}
