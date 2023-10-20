package camelcase

import (
	"strings"
)

// ToCamelCase converts a string into camelCase.
func ToCamelCase(s string) string {
	words := strings.Fields(s) // Split string by whitespace
	if len(words) == 0 {
		return ""
	}

	// Convert the first word to lowercase
	words[0] = strings.ToLower(words[0])

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}

	return strings.Join(words, "")
}
