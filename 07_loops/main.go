package main

import "fmt"

// * fizzbuzz challange
func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

//* loops
// func main() {
// 	// long method for loop
// 	// i := 1
// 	// for i <= 10 {
// 	// 	fmt.Printf("%d ", i)
// 	// 	i++
// 	// }
// 	// fmt.Println()

// 	//short method
// 	// for j := 0; j < 10; j++ {
// 	// 	fmt.Println(j)
// 	// }
// }
