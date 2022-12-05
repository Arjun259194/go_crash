package main

import "fmt"

func adder() func(int) int { // in this line of code "func(int) int" is return type of the adder function
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
