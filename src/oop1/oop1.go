package main

import (
	"fmt"
)

type Amount int

func (a *Amount) Add(add Amount) {
	*a += add
}

func main() {
	var a Amount
	a = 1
	a.Add(2)
	fmt.Println(a)
}
