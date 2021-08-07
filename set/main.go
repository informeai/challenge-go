package main

import (
	"errors"
	"fmt"
)

//Set is struct of implements sets in go.
type Set struct {
	Elements map[string]struct{}
}

//NewSet return isntance of Set.
func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set
}

//Add is responsable by add elements.
func (s *Set) Add(elem string) {
	s.Elements[elem] = struct{}{}
}

//Delete is delete elements of Set.
func (s *Set) Delete(elem string) error {
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("not element existed.")
	}
	delete(s.Elements, elem)
	return nil
}

//Contains is verify if exist element of Set.
func (s *Set) Contains(elem string) bool {
	_, exists := s.Elements[elem]
	return exists
}

//List print elements of Set.
func (s *Set) List() {
	for k, _ := range s.Elements {
		fmt.Println(k)
	}
}

func main() {
	set := NewSet()

	set.Add("green")
	set.Add("blue")
	set.Add("red")
	set.Add("blue")

	set.List()

	set.Delete("red")
	set.List()
}
