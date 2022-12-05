package main

import "fmt"

func main() {
	// define map
	// emails := make(map[string]string)

	//assing value
	// emails["arjun"] = "arjun259194@gmail.com"
	// emails["joshi"] = "joshi@gmail.com"
	// emails["zeeal"] = "zeeal@gmail.com"
	// emails["avani"] = "duck@gmail.com"

	// declare map and add key
	emails := map[string]string{"arjun": "aa@a.aom", "joshi": "nibba@gmai.com"}

	emails["avani"] = "duck@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["arjun"])

	//delete
	delete(emails, "arjun")

	fmt.Println(emails)
}
