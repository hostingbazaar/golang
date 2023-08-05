package main

import "fmt"

func main() {
	//1.decleration //condidtion //insitliaztion

	//differencease all element of lopp
	for i := 1; i <= 5; {
		fmt.Println(i)
		i++
	}

	//simple way to write down
	for x := 0; x <= 2; x++ {
		fmt.Println("Value is", x)
	}

}

//Sr. No.	Loop Type	Description
//1.	While Loop	In while loop, a condition is evaluated before processing a body of the loop. If a condition is true then and only then the body of a loop is executed.
//2.	Do-While Loop	In a do…while loop, the condition is always executed after the body of a loop. It is also called an exit-controlled loop.
//3.	For Loop	In a for loop, the initial value is performed only once, then the condition tests and compares the counter to a fixed value after each iteration, stopping the for loop when false is returned.
