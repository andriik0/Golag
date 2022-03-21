package logs
import "strings"


var applications = map[string]string{
	"❗": "recommendation",
	"🔍": "search",
	"☀": "weather",
}


// Application identifies the application emitting the given log.
func Application(log string) string {
	var index = strings.IndexAny(log, "❗🔍☀")
	if index == -1 {
		return "default"
	}
	return applications[string([]rune(log)[index])]
	
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	var index = strings.IndexAny(log, "❗🔍☀")
	if index == -1 {
		return true
	}
	return index < limit
}
