package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// todo represents data about a record todo.
type todo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Checked bool   `json:"is_checked"`
}

// todos slice to seed record todo data.
var todos = []todo{
	{ID: "1", Name: "Learn Golang", Checked: true},
	{ID: "2", Name: "Learn MongoDB", Checked: true},
	{ID: "3", Name: "Build Simple REST API with Golang & MongoDB", Checked: false},
}

// postTodos adds a todo from JSON received in the request body.
func postTodos(c *gin.Context) {
	var newTodo todo

	// Call BindJSON to bind the received JSON to
	// newTodo.
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	// Add the new TODO to the slice.
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// getTodoByID locates the todo whose ID value matches the id
// parameter sent by the client, then returns that todo as a response.
func getTodoByID(c *gin.Context) {
	// get parameter id, different from query
	id := c.Param("id")

	// Loop over the list of todos, looking for
	// an todo whose ID value matches the parameter.
	for _, a := range todos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

// getTodos responds with the list of all todos as JSON.
func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()

	// GET routes
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)

	// POST routes
	router.POST("/todos", postTodos)

	router.Run("localhost:8080")
}
