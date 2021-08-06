package repository

import (
	. "github.com/informeai/challenge-go/api/entities"
	"log"
	"testing"
)

//go test -v -run ^TestUserRepositoryMemoryGet
func TestUserRepositoryMemoryGet(t *testing.T) {
	urm := UserRepositoryMemory{}
	u, err := urm.Get(1)
	if err != nil {
		t.Errorf("TestUserRepositoryMemoryGet(): got -> %v, want: nil", err)
	}
	log.Println(u)
}

//go test -v -run ^TestUserRepositoryMemoryGetAll
func TestUserRepositoryMemoryGetAll(t *testing.T) {
	urm := UserRepositoryMemory{}
	us, err := urm.GetAll()
	if err != nil {
		t.Errorf("TestUserRepositoryMemoryGetAll(): got -> %v, want: nil", err)
	}
	log.Println(us)
}

//go test -v -run ^TestUserRepositoryMemoryCreate
func TestUserRepositoryMemoryCreate(t *testing.T) {
	urm := UserRepositoryMemory{}
	u := User{Id: 2, Name: "informeai", Age: 34, Email: "informeai@email.com"}
	us, err := urm.Create(u)
	if err != nil {
		t.Errorf("TestUserRepositoryMemoryCreate(): got -> %v, want: nil", err)
	}
	log.Println(us)
}

//go test -v -run ^TestUserRepositoryMemoryUpdate
func TestUserRepositoryMemoryUpdate(t *testing.T) {
	urm := UserRepositoryMemory{}
	u := User{Id: 1, Name: "gadelha", Age: 35, Email: "gadelha@email.com"}
	id, err := urm.Update(u)
	if err != nil {
		t.Errorf("TestUserRepositoryMemoryUpdate(): got -> %v, want: nil", err)
	}
	log.Println(id)
}

// go test -v -run ^TestUserRepositoryMemoryDelete
func TestUserRepositoryMemoryDelete(t *testing.T) {
	urm := UserRepositoryMemory{}
	b, err := urm.Delete(1)
	if err != nil {
		t.Errorf("TestUserRepositoryMemoryDelete(): got -> %v, want: nil", err)
	}
	log.Println(b)
}
