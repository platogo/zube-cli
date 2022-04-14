package utils

import (
	"strings"
	"unicode/utf8"
)

// Returns a string truncated to the given `maxLen`
func TruncateString(s string, maxLen int) string {
	if utf8.RuneCountInString(s) > maxLen {
		return s[:maxLen]
	}
	return s
}

// Converts a snake case string (e.g. `in_progress`) to title case (e.g. `In Progress`)
func SnakeCaseToTitleCase(s string) string {
	spacedWords := strings.ReplaceAll(s, "_", " ")
	return strings.Title(spacedWords)
}
