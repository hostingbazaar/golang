// Three types of variable
// 1.Local variable
// 2.Global variable
// 3.Package Level variable

package main

import "fmt"

//Global function this function you can use any where it is (Pascal case means first word capital)
var Result_num = 800

//pacakge level this variable you can use any where of packages it is (Camel case like someword small other one capital)
var myValue = 700

func main() {

	//local variable come on function (you can declare varibale in any cases Ee anyone it's store local)
	var datapoint = 360
	fmt.Println(datapoint)
}
