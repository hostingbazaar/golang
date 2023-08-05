package main

import "fmt"

func main() {

	var valueis int = 15

	if valueis > 10 {
		fmt.Println("value is big")
	} else if valueis == 10 {
		fmt.Println("value is equal")
	} else {
		fmt.Println("value is small")
	}
}
