package main

import (
	"fmt"
)
type Person struct {
	firstName string
	lastName  string
}
func (p Person) name() string {
	return p.firstName + " " + p.lastName
}
func main() {
	p := Person{"Flavio", "Copes"}
	fmt.Println(p.name())
}
