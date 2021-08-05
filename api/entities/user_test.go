package entities

import (
	"testing"
)

//go test -v -run ^TestNewUser
func TestNewUser(t *testing.T) {
	u := NewUser()
	if u == nil {
		t.Errorf("TestNewUser(): got -> nil, want: User{}")
	}
}
