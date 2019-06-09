package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth"
)

// looks at the HTTP request and returns the user id from the JWT claims
func GetCurrentUserId(r *http.Request) (string, error) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	if len(claims) == 0 {
		return "", errors.New("Could not determine current user")
	}
	userID := claims["id"].(string)
	if len(userID) == 0 {
		return "", errors.New("Could not determine current user")
	}

	return userID, nil
}
