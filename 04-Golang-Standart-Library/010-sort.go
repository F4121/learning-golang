package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// func (s *UserSlice) Len() int {  define slice without pointer because default value from slice is pointer
func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {

	/**
	temp := s[i]
	s[i] = s[j]
	s[j] = temp

	have same result / shorthand using below
	*/

	s[i], s[j] = s[j], s[i]
}

func main() {
	/**
	Package sort
	- Package sort adalah package yang berisikan utilitas untuk proses pengurutan
	- Agar data kita bisa diurutkan, kita harus mengimplementasi kontrak di interface sort.Interface

	reference
	https://golang.org/pkg/sort
	*/

	/**
	sort.Interface

	type Interface interface {
		- Len is the number of elements in the collection
		Len() int
		- Less reports whether the element with index i should sort before the element with index j.
		Less(i, j int) bool
		- Swap swaps the elements with indices i and j.
		Swap(i, j int)
	}
	*/

	users := []User{
		{"Jhon", 23},
		{"Doe", 18},
		{"Develop", 25},
		{"Agus", 30},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)

}
