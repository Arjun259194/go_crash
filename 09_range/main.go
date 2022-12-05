package main

import "fmt"

func main() {
	ids := []int{12, 43, 2, 4, 35, 5}

	// loop through ids
	// for _, id := range ids {
	// 	fmt.Printf("ID: %d\n", id)
	// }

	sum := 0
	for _, id := range ids {
		sum += id
	}

	// fmt.Println(sum)

	//range with map
	emails := map[string]string{
		"arjun": "aa@a.aom",
		"joshi": "nibba@gmai.com",
	}

	emails["arjun"] = "fucked@life.com"
	emails["zeeal"] = "zeeal@bro.com"
	emails["avani"] = "duck@friend.com"

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}
}
