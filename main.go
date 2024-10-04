package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// client and a server communicate w each other through a JSON
var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Wash the Dishes", Completed: false},
	{ID: "3", Item: "Learn English", Completed: false},
	{ID: "4", Item: "Go for a walk", Completed: false},
	{ID: "5", Item: "Cook a meal", Completed: false},
}

func getTodos(context *gin.Context) {
	if todos == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "todos list is empty"})
		return
	}

	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

/*
change that so we'll have a function that returns atodo object and error by surching via ID
then change the function, so it'll use that helper func
*/
func updateTodo(context *gin.Context) {
	//In the body, you only send the fields that you want to update, like title or completed
	//It'll specify the ID through the endpoint
	id := context.Param("id")

	var updatedTodo todo

	if err := context.BindJSON(&updatedTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	//add a getTodoById(id string)(*todo, error){}
	for i, t := range todos {
		if t.ID == id {
			if updatedTodo.Item != "" {
				todos[i].Item = updatedTodo.Item // Update the title if provided
			}
			todos[i].Completed = updatedTodo.Completed
			context.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.PUT("/todos/:id", updateTodo)

	router.Run("localhost:9090")
}
