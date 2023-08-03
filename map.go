package main

import "fmt"

func main() {
	//mapping data//
	// userDetails := make(map[string]string) //deceleration without this given below example you can declare and initialization also
	userDetails := map[string]string{ //initialization
		"fullname":   "Sachin Tedulkar",
		"email":      "sachin@gmail.com",
		"account_id": "33883839",
	}
	//simply print map
	fmt.Println(userDetails)
	//get data from map
	fmt.Println("User Fullname is :- ", userDetails["fullname"])
	//add new map key and value
	userDetails["phone_no"] = "1234567890"
	fmt.Println("New Map Data", userDetails)
	//how delete key data
	delete(userDetails, "email")
	fmt.Println("After Deleted Current Data", userDetails)
	//check those key is right or not
	password, validate := userDetails["password"] // you can choose any name validate behalf like ok, abc, etc
	fmt.Println("Checked Key Type Availble = ", password, validate)
	//without showing variable input you can get validate data
	_, pvalidate := userDetails["fullname"] // you can use _ to hide variable name where no need to show varibale value
	fmt.Println("Without getting variable data only show validate bollen data here = ", pvalidate)

}
