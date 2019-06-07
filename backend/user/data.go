package user

type Data interface {
	getAllUsers() []User
	getUserById(userId string) *User
	getByUsername(username string) *User
	create(user *User) error
	update(user *User) error
	delete(userId string) error
	Start()
	Stop()
}
