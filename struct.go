package main

import "fmt"

type Userdata struct {
	userId     int
	fullName   string
	emailId    string
	userPhotos []string //array data define like this

}

func main() {

	// user := Userdata{ //how define value on struct type but it's not right way why because if top key value
	// 				  // change their data value going wrong key value so define their key value also give below example
	// 	01,
	// 	"Ajeet Choubhey",
	// 	"ajeet@gmail.com",
	// 	[]string{"profile.png", "cover.jpg"},
	// }

	user := Userdata{
		userId:     01,
		fullName:   "Ajeet Choubhey",
		emailId:    "ajeet@gmail.com",
		userPhotos: []string{"profile.png", "cover.jpg"},
	}

	//print struct value
	fmt.Println(user)
	//get specify key value of struct
	fmt.Println("Welcome New User", user.fullName)
	//changed value of specify struct key data value
	user.emailId = "ajeet@developerbazaar.com"
	fmt.Println(user.fullName, "Your new email id updated", user.emailId)
	fmt.Println(user.fullName, "Your profile picture", user.userPhotos[0])
	//current struct value
	fmt.Println(user) // & Variable help to address same variable come on same data learn it more//
	//get length of value
}
