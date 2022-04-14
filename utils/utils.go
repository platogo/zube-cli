package utils

import (
	"strings"
)

// Returns a string truncated to the given `maxLen`
func TruncateString(s string, maxLen int) string {

	// Supports Japanese
	// Ref: Range loops https://blog.golang.org/strings
	// Source: https://dev.to/takakd/go-safe-truncate-string-9h0
	truncated := ""
	count := 0
	for _, char := range s {
		truncated += string(char)
		count++
		if count >= maxLen {
			break
		}
	}
	return truncated
}

// Converts a snake case string (e.g. `in_progress`) to title case (e.g. `In Progress`)
func SnakeCaseToTitleCase(s string) string {
	spacedWords := strings.ReplaceAll(s, "_", " ")
	return strings.Title(spacedWords)
}
