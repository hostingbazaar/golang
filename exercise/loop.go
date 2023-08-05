package main

import "fmt"

func main() {
	//printing table of 2
	table := 5
	multiple_by := 10

	for i := 1; i <= multiple_by; i++ {
		multiply := table * i
		fmt.Println(table, "x", i, "=", multiply)
	}
}
