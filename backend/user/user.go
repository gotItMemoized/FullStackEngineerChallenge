package user

type User struct {
	ID           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Username     string `json:"username" db:"username"`
	Password     string `json:"password,omitempty" db:"-"`
	NewPassword  string `json:"newPassword,omitempty" db:"-"`
	PasswordHash string `json:"-" db:"password"`
	// maybe not a good idea to send this to the frontend,
	// but frontend should only see it if you're an admin
	IsAdmin bool `json:"isAdmin,omitempty" db:"isadmin"`
}
