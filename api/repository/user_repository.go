package repository

import (
	. "github.com/informeai/challenge-go/api/entities"
)

//UserRepository is interface for implementation.
type UserRepository interface {
	Get(id int) (User, error)
	GetAll() ([]User, error)
	Create(u User) (User, error)
	Update(u User) (int, error)
	Delete(id int) (bool, error)
}
