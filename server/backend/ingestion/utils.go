package ingestion

import (
	"strings"
	"unicode"
)

// ToSnake convert the given string to snake case following the Golang format:
// acronyms are converted to lower-case and preceded by an underscore.
func ToSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		if !unicode.IsSpace(runes[i]) {
			out = append(out, unicode.ToLower(runes[i]))
		}
	}

	return string(out)
}

func StripNewLines(text string) string {
	return strings.Replace(strings.Replace(text, "\r", "", -1), "\n", "", -1)
}