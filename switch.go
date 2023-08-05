package main

import "fmt"

func main() {
	var value int = 9
	switch value {
	case 2, 9:
		fmt.Println("value is not right 1 ")
		fallthrough //is used for print next condition also
	case 4:
		fmt.Println("value is not right 2")
	case 8:
		fmt.Println("value is right 1")
		break //break is help to break the input with two print you break those result
		fmt.Println("value is right 2")
	default:
		fmt.Println("something went wrong")
	}

}
