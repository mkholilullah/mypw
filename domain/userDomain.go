package domain

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Password string
}

// UserRepository adalah interface untuk layer repository
type UserRepository interface {
	GetById(id int) (*User, error)
	Create(user *User) error
}
