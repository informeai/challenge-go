package usecases

import (
	. "github.com/informeai/challenge-go/api/entities"
)

//UserUseCase is provider for methods implementations.
type UserUseCase interface {
	Get(id int) (User, error)
	GetAll() ([]User, error)
	Create(u User) (User, error)
	Update(u User) (int, error)
	Delete(id int) (bool, error)
}
