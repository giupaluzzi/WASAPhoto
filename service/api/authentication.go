package api

import (
	"strings"
)

// extractToken extracts the Bearer token from the authorization header
func extractToken(auth string) string {
	tokens := strings.Split(auth, " ")
	if len(tokens) < 1 {
		return ""
	}
	return strings.TrimSpace(tokens[1])
}
