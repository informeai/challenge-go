package delivery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	// "path/filepath"
	. "github.com/informeai/challenge-go/api/entities"
	usecase "github.com/informeai/challenge-go/api/usecases"
	"net/http"
)

type Mux http.ServeMux

func (m *Mux) ServeHTTP(http.ResponseWriter, *http.Request) {}

//UserDeliveryMemory is responsable by route the users.
type UserDeliveryMemory struct {
	usecase usecase.UserUseCaseMemory
	server  http.Server
}

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
			fmt.Println(newUser)
			user, err := ud.usecase.Create(newUser)
			if err != nil {
				return
			}
			fmt.Println(user)
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

func (ud *UserDeliveryMemory) UpdateUser(path string) error {
	return nil
}

func (ud *UserDeliveryMemory) DeleteUser(path string) error {
	return nil
}

func (ud *UserDeliveryMemory) Run(address string) error {

	ud.server = http.Server{
		Addr:    address,
		Handler: nil,
	}
	err := ud.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
