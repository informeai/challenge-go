package repository

import (
	"errors"
	. "github.com/informeai/challenge-go/api/entities"
)

//slice the mockusers for test api.
var users []User = []User{
	{Id: 1, Name: "wellington", Age: 37, Email: "wellington@email.com"},
}

//UserRepository for mocking user.
type UserRepositoryMemory struct {
}

//Get of retrieve first user by id.
func (ur *UserRepositoryMemory) Get(id int) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}
	return User{}, nil
}

//GetAll return all user of repository.
func (ur *UserRepositoryMemory) GetAll() ([]User, error) {
	if len(users) > 0 {
		return users, nil
	}
	return []User{}, errors.New("slice of users is empty")
}

//Create is responsable by save user in mock users.
func (ur *UserRepositoryMemory) Create(u User) (User, error) {
	for _, user := range users {
		if user.Id == u.Id {
			return User{}, errors.New("user already exists.")
		}
	}
	users = append(users, u)
	return u, nil
}

//Update is change user of mock users.
func (ur *UserRepositoryMemory) Update(u User) (int, error) {
	for i := 0; i < len(users); i++ {
		if users[i].Id == u.Id {
			users[i] = u
			return users[i].Id, nil
		}
	}
	return 0, errors.New("user not found.")
}

//Delete is delete user of mock users.
func (ur *UserRepositoryMemory) Delete(id int) (bool, error) {
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			users = append(users[:i], users[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("error of delete user.")
}
