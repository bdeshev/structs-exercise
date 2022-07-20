package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	age  int
	DOB  time.Time
}

func main() {
	var p Person
	p.name = "Pesho"
	p.age = 21
	p.DOB = time.Now()

	p1 := Person{
		name: "Gosho",
		age:  21,
		DOB:  time.Now(),
	}

	p1Copy := p1
	p1Copy.age = 2222

	fmt.Println(p.name, p1.name, p1.age)
}
