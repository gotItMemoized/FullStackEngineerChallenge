package user

import (
	"errors"
	"log"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/jmoiron/sqlx"
)

func (u *UserService) getAllUsers() []User {
	var results []User
	err := u.db.Select(&results, `
		select * 
		from users
	`)

	if err != nil {
		log.Printf("Could not get all the users\n%+v\n", err)
		return nil
	}

	return results
}

func (u *UserService) getUserById(userId string) *User {
	var result User
	pid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		log.Printf("Couldn't parse int for getting\n%+v\n", err)
		return nil
	}
	err = u.db.Get(&result, `
		select *
		from users
		where id = $1
	`, pid)

	if err != nil {
		log.Printf("Could not get the user\n%+v\n", err)
		return nil
	}

	return &result
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
	if len(user.Username) == 0 {
		return errors.New("Invalid username cannot save user")
	}

	_, err := u.db.Exec(`
		insert into users 
            (name, password, username, isadmin) 
		values 
			($1, $2, $3, $4)
	`, user.Name, user.PasswordHash, user.Username, user.IsAdmin)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) update(user *User) error {
	if len(user.Username) == 0 {
		return errors.New("Invalid username cannot save user")
	}

	m := map[string]interface{}{"id": user.ID, "name": user.Name, "isadmin": user.IsAdmin, "username": user.Username}
	if len(user.PasswordHash) != 0 {
		m["password"] = user.PasswordHash
	}

	sql := `update users 
	set name=:name, username=:username, isadmin=:isadmin`
	if len(user.PasswordHash) != 0 {
		sql += `, password=:password `
	}
	sql += `where id=:id`
	_, err := u.db.NamedExec(sql, m)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) delete(id int64) error {
	_, err := u.db.Exec(`
		DELETE FROM users WHERE ID = $1
	`, id)
	if err != nil {
		log.Printf("Error deleting user \n%+v\n", err)
		return err
	}
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
