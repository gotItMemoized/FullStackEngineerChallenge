package handlers

import (
	"encoding/json"
	"errors"
	"log"
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

// this is technically inefficient, but allows for fast iterations and we can still get very fast responses locally
func WriteToOutput(w http.ResponseWriter, object interface{}) {
	output, err := json.Marshal(object)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(output)
	if err != nil {
		log.Printf("error while writing output: %+v\n", err)
	}
}
