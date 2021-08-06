package delivery

import (
	"encoding/json"
	"fmt"
	. "github.com/informeai/challenge-go/api/entities"
	usecase "github.com/informeai/challenge-go/api/usecases"
	"io/ioutil"
	"net/http"
	"strconv"
)

//UserDeliveryMemory is responsable by route the users.
type UserDeliveryMemory struct {
	usecase usecase.UserUseCaseMemory
	server  http.Server
}

//GetUser is connect of route the get single user by id.
func (ud *UserDeliveryMemory) GetUser(path string) error {
	var err error
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			//receiver uniq number id. eg: 1 or 2,3...
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()
			id, err := strconv.Atoi(string(b))
			if err != nil {
				return
			}
			user, err := ud.usecase.Get(id)
			if err != nil {
				return
			}
			userMarshal, err := json.Marshal(user)
			if err != nil {
				return
			}
			userJson := string(userMarshal)
			fmt.Fprintf(w, userJson)
		} else {
			fmt.Fprintf(w, "method not allowed.")
		}

	})

	return err
}

//GetAllUsers is connect of route for list all users.
func (ud *UserDeliveryMemory) GetAllUsers(path string) error {
	var err error
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			//get all users.
			users, err := ud.usecase.GetAll()
			if err != nil {
				return
			}
			usersMarshal, err := json.Marshal(users)
			if err != nil {
				return
			}
			usersJson := string(usersMarshal)
			fmt.Fprintf(w, usersJson)
		} else {
			fmt.Fprintf(w, "method not allowed.")
		}

	})

	return err
}

//CreateUser is connect of route the create user.
func (ud *UserDeliveryMemory) CreateUser(path string) error {
	var err error
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			newUser := User{}
			//receiver user
			u, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()
			err = json.Unmarshal(u, &newUser)
			if err != nil {
				return
			}
			user, err := ud.usecase.Create(newUser)
			if err != nil {
				return
			}
			userMarshal, err := json.Marshal(user)
			if err != nil {
				return
			}
			userJson := string(userMarshal)
			fmt.Fprintf(w, userJson)
		} else {
			fmt.Fprintf(w, "method not allowed.")
		}

	})

	return err
}

//UpdateUser is connect of route the update user.
func (ud *UserDeliveryMemory) UpdateUser(path string) error {
	var err error
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			updateUser := User{}
			//receiver user
			u, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()
			err = json.Unmarshal(u, &updateUser)
			if err != nil {
				return
			}
			id, err := ud.usecase.Update(updateUser)
			if err != nil {
				return
			}
			idMarshal, err := json.Marshal(id)
			if err != nil {
				return
			}
			idJson := string(idMarshal)
			fmt.Fprintf(w, idJson)
		} else {
			fmt.Fprintf(w, "method not allowed.")
		}

	})

	return err
}

//DeleteUser is connect of route the delete user by id.
func (ud *UserDeliveryMemory) DeleteUser(path string) error {
	var err error
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			//receiver id for delete
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()
			id, err := strconv.Atoi(string(b))
			if err != nil {
				return
			}
			sucess, err := ud.usecase.Delete(id)
			if err != nil {
				return
			}
			idMarshal, err := json.Marshal(sucess)
			if err != nil {
				return
			}
			idJson := string(idMarshal)
			fmt.Fprintf(w, idJson)
		} else {
			fmt.Fprintf(w, "method not allowed.")
		}

	})

	return err
}

//Run is responsable by running server.
func (ud *UserDeliveryMemory) Run(address string) error {

	ud.server = http.Server{
		Addr:    address,
		Handler: nil,
	}
	err := ud.GetUser("/user")
	if err != nil {
		return err
	}
	err = ud.GetAllUsers("/users")
	if err != nil {
		return err
	}
	err = ud.CreateUser("/user/create")
	if err != nil {
		return err
	}
	err = ud.UpdateUser("/user/update")
	if err != nil {
		return err
	}
	err = ud.DeleteUser("/user/delete")
	if err != nil {
		return err
	}
	err = ud.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
