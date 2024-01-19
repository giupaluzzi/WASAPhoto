package api

import (
	"strings"
)

// removeBearer extracts the Bearer token from the authorization header
func removeBearer(auth string) string {
	tokens := strings.Split(auth, " ")
	if len(tokens) < 1 {
		return ""
	}
	return strings.TrimSpace(tokens[1])
}
