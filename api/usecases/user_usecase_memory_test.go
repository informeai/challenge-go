package usecases

import (
	. "github.com/informeai/challenge-go/api/entities"
	"log"
	"testing"
)

//go test -v -run ^TestUserUseCaseMemoryGet
func TestUserUseCaseMemoryGet(t *testing.T) {
	ucm := UserUseCaseMemory{}
	u, err := ucm.Get(1)
	if err != nil {
		t.Errorf("TestUserUseCaseMemoryGet(): got -> %v, want: nil", err)
	}
	log.Println(u)

}

//go test -v -run ^TestUserUseCaseMemoryGetAll
func TestUserUseCaseMemoryGetAll(t *testing.T) {
	ucm := UserUseCaseMemory{}
	us, err := ucm.GetAll()
	if err != nil {
		t.Errorf("TestUserUseCaseMemoryGetAll(): got -> %v, want: nil", err)
	}
	log.Println(us)
}

//go test -v -run ^TestUserUseCaseMemoryCreate
func TestUserUseCaseMemoryCreate(t *testing.T) {
	ucm := UserUseCaseMemory{}
	user := User{Id: 2, Name: "informeai", Age: 45, Email: "informeai@email.com"}
	u, err := ucm.Create(user)
	if err != nil {
		t.Errorf("TestUserUseCaseMemoryCreate(): got -> %v, want: nil", err)
	}
	log.Println(u)
}

//go test -v -run ^TestUserUseCaseMemoryUpdate
func TestUserUseCaseMemoryUpdate(t *testing.T) {
	ucm := UserUseCaseMemory{}
	newUser := User{Id: 1, Name: "informeai", Age: 34, Email: "informeai@email.com"}
	id, err := ucm.Update(newUser)
	if err != nil {
		t.Errorf("TestUserUseCaseMemoryUpdate(): got -> %v, want: nil", err)
	}
	log.Println(id)
}

//go test -v -run ^TestUserUseCaseMemoryDelete
func TestUserUseCaseMemoryDelete(t *testing.T) {
	ucm := UserUseCaseMemory{}
	b, err := ucm.Delete(1)
	if err != nil {
		t.Errorf("TestUserUseCaseMemoryDelete(): got -> %v, want: nil", err)
	}
	log.Println(b)
}
