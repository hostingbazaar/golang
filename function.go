package main

import "fmt"

//simple funciton call//
func simple() {
	var name string = "string value is that"
	fmt.Println(name)
}

//return method
func percentage(num int, percent int) int {
	var getpercentage int = num * percent / 100
	return getpercentage
}

//end number of values

func endno(nums ...int) {
	sum := 0
	for _, num := range nums {
		fmt.Println(num)
		sum += num
	}
	// Print the sum
	fmt.Println("Sum:", sum)

}

func main() {
	//simple print
	simple()

	//getting value
	fmt.Println("Your Percentage Value is :- ", percentage(1000, 8))

	//print end number value
	endno(1, 2, 3, 4, 5)
}
