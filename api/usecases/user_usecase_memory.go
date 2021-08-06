package usecases

import (
	. "github.com/informeai/challenge-go/api/entities"
	repo "github.com/informeai/challenge-go/api/repository"
)

//UserUseCaseMemory is connect repository and delivery
type UserUseCaseMemory struct {
	repo repo.UserRepositoryMemory
}

//Get return first user by id.
func (us *UserUseCaseMemory) Get(id int) (User, error) {
	u, err := us.repo.Get(id)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

//GetAll return all users.
func (us *UserUseCaseMemory) GetAll() ([]User, error) {
	users, err := us.repo.GetAll()
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

//Create is responsable by create user.
func (us *UserUseCaseMemory) Create(u User) (User, error) {
	user, err := us.repo.Create(u)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

//Update is responsable by change user.
func (us *UserUseCaseMemory) Update(u User) (int, error) {
	id, err := us.repo.Update(u)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//Delete is delete the user by id.
func (us *UserUseCaseMemory) Delete(id int) (bool, error) {
	b, err := us.repo.Delete(id)
	if err != nil {
		return false, err
	}
	return b, nil
}
