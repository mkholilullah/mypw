package domain

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserRepository adalah interface untuk layer repository
type UserRepository interface {
	GetById(id int) (*User, error)
	Create(user *User) error
	GetByUserName(username string) (*User, error)
}

type UserUseCase interface {
	Register(user *User) (*User, error)
	Login(username, password string) (*User, error)
}