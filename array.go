package main

import "fmt"

func main() {
	//data type array call
	var arr = [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	//simple array call with data type without length put
	var arr1 = [...]int{4, 5, 6, 7, 8, 9}
	fmt.Println(arr1)
	//check the length or array
	fmt.Println("array length is", len(arr1))
	//call the array value
	fmt.Println("Array value of 3 is", arr1[3])
	//changed the array value
	arr1[3] = 14
	fmt.Println("Changed 3 index array value 7 to ", arr1[3], "latest arra value is", arr1)
	//filter array on range wise it's called slices//
	fmt.Println("Array range value capture", arr1[0:4])

	//Multidimensional Array//
	multi_array := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(multi_array)

	//Call to multidimensional array//
	//full array//
	fmt.Println(multi_array[2])
	//single value//
	fmt.Println(multi_array[2][0])

}
