package entity

type User struct {
	Name string
}

type UserRepository interface {
	Authenticate(*User) (bool, error)
	Regist(*User) error
}
