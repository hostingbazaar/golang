package main

import "fmt"

func main() {
	//defer is help to run those things on last ex :- you intilizaton 1,2,3 and diff 2 so result is 1,3,2
	fmt.Println("1")
	defer fmt.Println("2") //2 is printing on last
	fmt.Println("3")

	//real example :- for example you make connection of database then doing code then close database
	//connection but in the case you forget to run those thing better than you put on starting and do defer on it

	//normal do
	fmt.Println("db connection")
	fmt.Println("coding do")
	fmt.Println("db closed") //if you forget to do this

	//do like this

	defer fmt.Println("db closed")
	fmt.Println("db connection")
	fmt.Println("coding do")

}
