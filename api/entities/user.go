package entities

type User struct {
	Name  string
	Age   int
	Email string
}

func NewUser() *User {
	return &User{}
}
