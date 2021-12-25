package predicate

import "strings"

func findSubstr(f func(s, substr string) bool, s string, substr ...string) (string, bool) {
	for _, ss := range substr {
		if f(s, ss) {
			return ss, true
		}
	}

	return "", false
}

// Helper for strings.contains
func Contains(s string, substr ...string) (string, bool) {
	return findSubstr(strings.Contains, s, substr...)
}
