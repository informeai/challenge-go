package controller

import (
// "github.com/informeai/challenge-go/entities"
// "net/http"
)

//UserInterface is provider for methods the UserController.
type UserInterface interface {
	Get(id int) User
	GetAll() []User
	Create(u User) User
	Update(u User) int
	Delete(id int) bool
}

//UserController responsible for making the connection between the user layers of the api.
type UserController struct {
}

//Get return first user by id.
func (uc *UserController) Get(id int) User {

}

//GetAll return all users.
func (uc *UserController) GetAll() []User {

}

//Create is responsable by create user.
func (uc *UserController) Create(u User) User {

}

//Update is responsable by change user.
func (uc *UserController) Update(u User) int {

}

//Delete is delete the user by id.
func (uc *UserController) Delete(id int) bool {

}
