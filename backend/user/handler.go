package user

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// These are typically very short, but for sake of simplicity we're just using one token
const tokenTime int64 = 20000 * 60 * 60 * 1000

type UserService struct {
	db   *sqlx.DB
	auth *jwtauth.JWTAuth
}

func (u *UserService) All(w http.ResponseWriter, r *http.Request) {
	// maybe add filters
	users := u.getAllUsers()

	if users == nil {
		http.Error(w, "[]", http.StatusNotFound)
		return
	}

	writeToOutput(w, users)
}

func (u *UserService) Get(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	user := u.getUserById(userID)

	if user == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	writeToOutput(w, user)
}

func (u *UserService) CurrentUser(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userID := claims["id"].(string)
	if len(userID) == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	user := u.getUserById(userID)

	if user == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	writeToOutput(w, user)
}

func (u *UserService) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: create
	writeToOutput(w, User{})
}

type credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"-"`
}

type tokens struct {
	JWT string `json:"jwt"`
}

func (u *UserService) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: updates
	writeToOutput(w, User{})
}

func (u *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: delete
	writeToOutput(w, User{})
}

func (u *UserService) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var creds credentials
	err := decoder.Decode(&creds)
	if err != nil || len(creds.Username) == 0 || len(creds.Password) == 0 {
		log.Print(errors.Wrap(err, "Could not decode user"))
		http.Error(w, "Error logging in user", http.StatusBadRequest)
		return
	}

	user := u.getByUsername(creds.Username)
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

	_, s, err := u.auth.Encode(jwtauth.Claims(jwt.MapClaims{"id": user.ID, "loggedIn": true, "isAdmin": user.IsAdmin, "exp": time.Now().Unix() + (tokenTime)}))

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
