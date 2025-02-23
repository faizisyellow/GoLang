package user

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func New(name, email, password string) *User {
	return &User{
		ID:       1,
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (u *User) ResetPassword() User {
	u.Password = ""
	return *u
}

func (u *User) SetPassword(password string) {

	u.Password = password
}
