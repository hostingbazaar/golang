package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//import the pacakges

type Todo struct {
	Id     string `json : "id"`
	Item   string `json : "item"` //mapping
	Status bool   `json : "status`
}

//creating to do array

var Todolist = []Todo{
	{Id: "1", Item: "wake up", Status: false},
	{Id: "2", Item: "meditation", Status: false},
	{Id: "3", Item: "book reading", Status: false},
}

// context is bunc of information like header and incoming http request
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Todolist) //convert this thing on json
}

// adding new struct data on array
func addTodo(context *gin.Context) {
	var newAdd Todo
	error := context.BindJSON(&newAdd)
	if error != nil {
		err := "Null value not taken"
		context.JSON(400, gin.H{"error": err})
	} else {
		Todolist = append(Todolist, newAdd)
		context.IndentedJSON(http.StatusCreated, newAdd)
	}
}

// gettodo id//
func getId(context *gin.Context) {
	id := context.Param("id")
	Todo, err := getbyID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, Todo)
	return

}

//get individual data of task

func getbyID(id string) (*Todo, error) {
	for i, t := range Todolist {
		if t.Id == id {
			return &Todolist[i], nil //nill is error show
		}
	}
	return nil, errors.New("to do not found")

}

func main() {
	router := gin.Default() //router is our server

	//creating endpoint
	router.GET("/todos", getTodos)    //function you called do run on rotuter link
	router.POST("/addtodo", addTodo)  //function you called do run on rotuter link
	router.GET("/getbyid/:id", getId) //function you called do run on rotuter link

	// where our project run that our router
	router.Run("localhost:9090")

}
