package utils

// Returns a string truncated to the given `maxLen`
func TruncateString(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen]
	}
	return s
}
