package delivery

import (
	// "log"
	"testing"
)

//go test -v -run ^TestUserDeliveryMemoryGetUser
func TestUserDeliveryMemoryGetUser(t *testing.T) {
	udm := UserDeliveryMemory{}
	err := udm.GetUser("/user")
	if err != nil {
		t.Errorf("TestUserDeliveryMemoryGetUser(): got -> %v, want: nil", err)
	}
	err = udm.Run(":4000")
	if err != nil {
		t.Errorf("TestUserDeliveryMemoryGetUser(): got -> %v, want: nil", err)
	}

}

//go test -v -run ^TestUserDeliveryMemoryGetAllUsers
func TestUserDeliveryMemoryGetAllUsers(t *testing.T) {
	udm := UserDeliveryMemory{}
	err := udm.GetAllUsers("/users")
	if err != nil {
		t.Errorf("TestUserDeliveryMemoryGetAllUsers(): got -> %v, want: nil", err)
	}
	err = udm.Run(":4000")
	if err != nil {
		t.Errorf("TestUserDeliveryMemoryGetAllUsers(): got -> %v, want: nil", err)
	}
}

//go test -v -run ^TestUserDeliveryMemoryCreateUser
func TestUserDeliveryMemoryCreateUser(t *testing.T) {
	udm := UserDeliveryMemory{}
	err := udm.CreateUser("/user/create")
	if err != nil {
		t.Errorf("TestUserDeliveryMemoryCreateUser(): got -> %v, want: nil", err)
	}
	err = udm.Run(":4000")
	if err != nil {
		t.Errorf("TestUserDeliveryMemoryCreateUser(): got -> %v, want: nil", err)
	}
}

//go test -v -run ^TestUserDeliveryMemoryUpdateUser
func TestUserDeliveryMemoryUpdateUser(t *testing.T) {
	udm := UserDeliveryMemory{}
	err := udm.UpdateUser("/user/update")
	if err != nil {
		t.Errorf("TestUserDeliveryMemoryUpdateUser(): got -> %v, want: nil", err)
	}
	err = udm.Run(":4000")
	if err != nil {
		t.Errorf("TestUserDeliveryMemoryUpdateUser(): got -> %v, want: nil", err)
	}
}

//go test -v -run ^TestUserDeliveryMemoryDeleteUser
func TestUserDeliveryMemoryDeleteUser(t *testing.T) {
	udm := UserDeliveryMemory{}
	err := udm.DeleteUser("/user/delete")
	if err != nil {
		t.Errorf("TestUserDeliveryMemoryDeleteUser(): got -> %v, want: nil", err)
	}
	err = udm.Run(":4000")
	if err != nil {
		t.Errorf("TestUserDeliveryMemoryDeleteUser(): got -> %v, want: nil", err)
	}
}
