package user

import (
	"errors"
	"log"
	"strconv"
)

type MapData struct {
	Seed      bool
	userCount int
	users     map[string]User
}

func (u *MapData) GetAllUsers() []User {
	return u.getAllUsers()
}

func (u *MapData) getAllUsers() []User {
	results := make([]User, len(u.users))
	ind := 0
	for _, user := range u.users {
		results[ind] = user
		ind += 1
	}
	return results
}

func (u *MapData) GetUserById(userId string) *User {
	return u.getUserById(userId)
}

func (u *MapData) getUserById(userId string) *User {
	user, ok := u.users[userId]

	if !ok {
		log.Printf("Could not get the user. Not found in map\n")
		return nil
	}

	return &user
}

func (u *MapData) GetByUsername(username string) *User {
	return u.getByUsername(username)
}

func (u *MapData) getByUsername(username string) *User {
	var result User
	for _, user := range u.users {
		if user.Username == username {
			result = user
			break
		}
	}

	if len(result.ID) == 0 {
		log.Printf("Could not get the user. Not found in map\n")
		return nil
	}

	return &result
}

func (u *MapData) Create(user *User) error {
	return u.create(user)
}

func (u *MapData) create(user *User) error {
	if len(user.Username) == 0 {
		return errors.New("Invalid username cannot save user")
	}

	u.userCount += 1
	user.ID = strconv.Itoa(u.userCount)
	u.users[user.ID] = *user

	return nil
}

func (u *MapData) Update(user *User) error {
	return u.update(user)
}

func (u *MapData) update(user *User) error {
	if len(user.Username) == 0 {
		return errors.New("Invalid username cannot save user")
	}

	u.users[user.ID] = *user
	return nil
}

func (u *MapData) Delete(id string) error {
	return u.delete(id)
}

func (u *MapData) delete(id string) error {
	delete(u.users, id)
	return nil
}

func (u *MapData) Start() {
	log.Println("Setting up default users")
	u.users = make(map[string]User)
	if u.Seed {
		u.userCount = 2
		u.users["1"] = User{
			ID:          "1",
			IsAdmin:     true,
			Name:        "James",
			Username:    "admin",
			Password:    "ssh",
			NewPassword: "ssh",
			// asdfasdfasdf
			PasswordHash: "$2a$10$p3cYUFyWcYvd1FNs/2nSD.Z9ZEbU8TCxlsrAbAR3jAXC2dspqjPKy",
		}
		u.users["2"] = User{
			ID:          "2",
			IsAdmin:     false,
			Name:        "Jamie",
			Username:    "user",
			Password:    "ssh",
			NewPassword: "ssh",
			// asdfasdfasdf
			PasswordHash: "$2a$10$p3cYUFyWcYvd1FNs/2nSD.Z9ZEbU8TCxlsrAbAR3jAXC2dspqjPKy",
		}
	}
	log.Println(" - Success")
}

func (u *MapData) Stop() {
}
