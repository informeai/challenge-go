package repository

//slice the mockusers for test api.
var users []User

//RepositoryUser for mocking user.
type RepositoryUser struct {
}

//Get of retrieve first user by id.
func (r *RepositoryUser) Get(id int) MockUser {

}

//GetAll return all user of repository.
func (r *RepositoryUser) GetAll() []User {

}

//Create is responsable by save user in mock users.
func (r *RepositoryUser) Create(u User) User {

}

//Update is change user of mock users.
func (r *RepositoryUser) Update(u User) int {

}

//Delete is delete user of mock users.
func (r *RepositoryUser) Delete(id int) bool {

}
