package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go Data Stream")

	user := map[string]interface{}{
		"id":    1,
		"name":  "John Doe",
		"email": "",
	}

	fmt.Println(user)
}
