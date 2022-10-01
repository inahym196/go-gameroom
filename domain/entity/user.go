package entity

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) GetName() string {
	return u.name
}
