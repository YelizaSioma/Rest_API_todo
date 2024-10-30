package main

/*
Endpoints
1. Get All Todos
Endpoint: GET /todos

Description: Retrieves a list of all to-do items.

2. Get a _Todo by ID
Endpoint: GET /todos/:id

Description: Retrieves details of a single to-do item by its ID.

Path Parameter:

id (integer) – ID of the to-do item to retrieve.

3. Toggle _Todo Status
Endpoint: PATCH /todos/:id

Description: Toggles the completed status of the specified to-do item.

Path Parameter:

id (integer) – ID of the to-do item to toggle.

4. Add a New _Todo
Endpoint: POST /todos

Description: Adds a new to-do item.

5. Update a _Todo
Endpoint: PUT /todos/:id

Description: Updates the title and status of an existing to-do item.

Path Parameter:

id (integer) – ID of the to-do item to update.

Common Response Codes
200 OK: Request succeeded.
201 Created: Resource successfully created.
404 Not Found: Resource not found.
400 Bad Request: Invalid request payload or parameters.
*/

import (
	"errors"
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

// client and a server communicate with each other through a JSON
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

	//check if already exists
	for _, todo := range todos {
		if todo.ID == newTodo.ID {
			context.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "Todo already exist"})
			return
		}
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, todo := range todos {
		if todo.ID == id {

			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

// This function will return only specified todofrom all list of todos. If not found - return an error
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func updateTodo(context *gin.Context) {
	id := context.Param("id")
	currTodo, currErr := getTodoById(id)

	if currErr != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	var updatedTodo todo

	if err := context.BindJSON(&updatedTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if updatedTodo.Item != "" {
		currTodo.Item = updatedTodo.Item // Update the title if provided
	}
	currTodo.Completed = updatedTodo.Completed

	context.IndentedJSON(http.StatusOK, currTodo)
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.PUT("/todos/:id", updateTodo)

	err := router.Run("localhost:9090")
	if err != nil {
		return
	}
}
