package main

import "fmt"

func main() {
	var age int32 = 32
	const isCool = true

	email, name := "arjun234@gmail.com", "arjun"
	size := 2.5

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T %T\n", name, age)
}
