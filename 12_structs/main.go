package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method {value resever}
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthdat method {pointer reciever}
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried {pointer reciver}
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct

	// 1
	person_1 := Person{
		firstName: "Arjun",
		lastName:  "mistry",
		city:      "mehsana",
		gender:    "M",
		age:       19,
	}

	person_2 := Person{
		firstName: "nayna",
		lastName:  "agarvan",
		city:      "ahmadabad",
		gender:    "F",
		age:       19,
	}

	//2
	// person_1 := Person{
	// 	"Joshi",
	// 	"mahir",
	// 	"mehsana",
	// 	"M",
	// 	17,
	// }

	fmt.Println(person_1)
	fmt.Println(person_1.greet())
	person_1.hasBirthday()
	fmt.Println(person_1.greet())

	fmt.Println(person_2.greet())
	person_2.getMarried("singh")
	fmt.Println(person_2.greet())
}
