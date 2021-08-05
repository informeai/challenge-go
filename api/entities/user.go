package entities

//User of type entitie the api.
type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

//NewUser return of new user instance.
func NewUser() *User {
	return &User{}
}
