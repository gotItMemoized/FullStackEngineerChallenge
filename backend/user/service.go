package user

import (
	"log"

	"github.com/go-chi/jwtauth"
	"github.com/jmoiron/sqlx"
)

func (u *UserService) getAllUsers() []User {
	return []User{}
}

func (u *UserService) getUserById(userId string) *User {
	return nil
}

func (u *UserService) getByUsername(username string) *User {
	var result User
	err := u.db.Get(&result, `
		select *
		from users
		where username = $1
	`, username)

	if err != nil {
		log.Printf("Could not get the user\n%+v\n", err)
		return nil
	}

	return &result
}

func (u *UserService) create(user *User) error {
	return nil
}

func (u *UserService) Start(db *sqlx.DB, auth *jwtauth.JWTAuth) {
	u.db = db
	u.auth = auth
}

func (u *UserService) Stop() {
	if u.db != nil {
		u.db.Close()
	}
}
