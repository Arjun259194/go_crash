package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)

	//Use * to read val from address
	fmt.Println(*b)

	// change value with pointer
	*b = 10
	fmt.Println(a)
}
