package user

import (
	"errors"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type SqlData struct {
	DB *sqlx.DB
}

func (u *SqlData) getAllUsers() []User {
	var results []User
	err := u.DB.Select(&results, `
		select * 
		from users
	`)

	if err != nil {
		log.Printf("Could not get all the users\n%+v\n", err)
		return nil
	}

	return results
}

func (u *SqlData) getUserById(userId string) *User {
	var result User
	pid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		log.Printf("Couldn't parse int for getting\n%+v\n", err)
		return nil
	}
	err = u.DB.Get(&result, `
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

func (u *SqlData) getByUsername(username string) *User {
	var result User
	err := u.DB.Get(&result, `
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

func (u *SqlData) create(user *User) error {
	if len(user.Username) == 0 {
		return errors.New("Invalid username cannot save user")
	}

	_, err := u.DB.Exec(`
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

func (u *SqlData) update(user *User) error {
	if len(user.Username) == 0 {
		return errors.New("Invalid username cannot save user")
	}

	m := map[string]interface{}{"id": user.ID, "name": user.Name, "isadmin": user.IsAdmin, "username": user.Username}
	if len(user.PasswordHash) != 0 {
		m["password"] = user.PasswordHash
	}

	sql := `update users 
		set name=:name, username=:username, isadmin=:isadmin `
	if len(user.PasswordHash) != 0 {
		sql += `, password=:password `
	}
	sql += ` where id=:id`
	_, err := u.DB.NamedExec(sql, m)

	if err != nil {
		return err
	}
	return nil
}

func (u *SqlData) delete(id string) error {
	_, err := u.DB.Exec(`
		DELETE FROM users WHERE ID = $1
	`, id)
	if err != nil {
		log.Printf("Error deleting user \n%+v\n", err)
		return err
	}
	return nil
}

func (u *SqlData) Start() {
}

func (u *SqlData) Stop() {
	if u.DB != nil {
		u.DB.Close()
	}
}
