package main

import "fmt"

func main() {

	//---START FIRST METHOD----///
	//simple variable use//
	var num int
	fmt.Print(num) //default value print 0 becasue that default data is 0

	//how variable print number
	var marks int //declaration
	marks = 98    //Instiliazation
	fmt.Println("your marks is ", marks)
	//---END FIRST METHOD---//

	//---START SECOND METHOD----///

	//a//
	var passmark int = 33 //declaration with instilization
	fmt.Println("Passing mark is", passmark)

	//b//
	var declare = 75 //declaration with instilization//
	fmt.Println("Your declaration value is", declare, "%")

	//---END SECOND METHOD---//

	//---START THIRD METHOD---//

	topper_marks := 99
	fmt.Println("Topper marks is", topper_marks, "%")

	//---END THIRD METHOD---//

}
