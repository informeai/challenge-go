package delivery

//UserDelivery is interface for routes the users.
type UserDelivery interface {
	GetUser(path string) error
	GetAllUsers(path string) error
	CreateUser(path string) error
	UpdateUser(path string) error
	DeleteUser(path string) error
}
