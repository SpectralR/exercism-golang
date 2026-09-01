package advanced

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var found []string
	for _, char := range log {
		switch {
		case string(char) == "❗":
			found = append(found, "recommendation")
		case string(char) == "🔍":
			found = append(found, "search")
		case string(char) == "☀":
			found = append(found, "weather")
		}
	}
	if len(found) == 0 {
		return "default"
	}

	return found[0]
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.Replace(log, string(oldRune), string(newRune), -1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
