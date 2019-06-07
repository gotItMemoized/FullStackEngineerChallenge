package user

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// These are typically very short, but for sake of simplicity we're just using one token
const tokenTime int64 = 20000 * 60 * 60 * 1000

type UserHandler struct {
	Auth *jwtauth.JWTAuth
	Data Data
}

func (u *UserHandler) All(w http.ResponseWriter, r *http.Request) {
	// maybe add filters
	users := u.Data.getAllUsers()

	if users == nil {
		http.Error(w, "[]", http.StatusNotFound)
		return
	}

	writeToOutput(w, users)
}

func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	user := u.Data.getUserById(userID)

	if user == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	writeToOutput(w, user)
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not decode user"))
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}

	if u.Data.getByUsername(user.Username) != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not hash password"))
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}
	user.PasswordHash = string(hash)

	err = u.Data.create(&user)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not save user"))
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}
	// in some cases it'd be nice to return the userid on create,
	// but in this case we'll skip that
	w.WriteHeader(http.StatusNoContent)
}

func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	idToUpdate := chi.URLParam(r, "id")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var updatedUser User
	err := decoder.Decode(&updatedUser)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not decode user"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedUser.ID = idToUpdate
	possiblyExistingUser := u.Data.getByUsername(updatedUser.Username)
	if possiblyExistingUser != nil && possiblyExistingUser.ID != updatedUser.ID {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	if len(updatedUser.NewPassword) != 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(updatedUser.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Print(errors.Wrap(err, "Could not hash password"))
			http.Error(w, "Error updating user", http.StatusInternalServerError)
			return
		}
		updatedUser.PasswordHash = string(hash)
	}

	err = u.Data.update(&updatedUser)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not update user"))
		http.Error(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	toDeleteId := chi.URLParam(r, "id")

	_, claims, _ := jwtauth.FromContext(r.Context())
	userID := claims["id"].(string)
	if len(userID) == 0 {
		http.Error(w, "Could not confirm user", http.StatusUnauthorized)
		return
	}

	if userID == toDeleteId {
		log.Printf("User tried to delete themselves\n")
		http.Error(w, "Cannot delete yourself", http.StatusUnauthorized)
		return
	}

	err := u.Data.delete(toDeleteId)
	if err != nil {
		log.Printf("Couldn't delete user\n%+v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var creds credentials
	err := decoder.Decode(&creds)
	if err != nil || len(creds.Username) == 0 || len(creds.Password) == 0 {
		log.Print(errors.Wrap(err, "Could not decode user"))
		http.Error(w, "Error logging in user", http.StatusBadRequest)
		return
	}

	user := u.Data.getByUsername(creds.Username)
	if user == nil {
		log.Printf("Attempted to login for unknown username %s\n", creds.Username)
		http.Error(w, "Error logging in user", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password))
	if err != nil {
		log.Printf("Bad password for username %s\n", user.Username)
		http.Error(w, "Error logging in user", http.StatusBadRequest)
		return
	}

	_, s, err := u.Auth.Encode(jwtauth.Claims(jwt.MapClaims{"id": user.ID, "loggedIn": true, "isAdmin": user.IsAdmin, "exp": time.Now().Unix() + (tokenTime)}))

	if err != nil {
		log.Printf("Could not set up tokens %s\n", user.Username)
		http.Error(w, "Error logging in user", http.StatusBadRequest)
		return
	}

	writeToOutput(w, tokens{
		JWT: s,
	})
}

// this is technically inefficient, but allows for fast iterations and we can still get very fast responses locally
func writeToOutput(w http.ResponseWriter, object interface{}) {
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

type credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"-"`
}

type tokens struct {
	JWT string `json:"jwt"`
}
