package main

import "fmt"

func main() {
	//Slices are similar to arrays, but are more powerful and flexible.
	//Like arrays, slices are also used to store multiple values of the same type in a single variable.
	//However, unlike arrays, the length of a slice can grow and shrink as you see fit.
	var array []int = []int{1, 2, 3, 4}
	fmt.Println(array)
	array = append(array, 5, 6, 7) //append help to add new value in index
	fmt.Println("Latest Array value", array)
	fmt.Println("Slicing Value is", array[6])
	//added 2 array value in one array index
	var array1 []int = []int{10, 20, 30, 40, 50}
	var array2 []int = []int{60, 70, 80, 90, 100}
	array = append(array1, array2...)
	fmt.Println("2 Array merge", array)
	fmt.Println("7 Array value", array[8])
}
