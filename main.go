package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// todo struct.
type todo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Checked *bool  `json:"is_checked"`
}

// default value for assigning boolean addresses.
var DefaultTrue, DefaultFalse = true, false

// todos data collection.
var todos = []todo{
	{ID: "1", Name: "Learn Golang", Checked: &DefaultTrue},
	{ID: "2", Name: "Learn MongoDB", Checked: &DefaultTrue},
	{ID: "3", Name: "Build Simple REST API with Golang & MongoDB", Checked: &DefaultFalse},
}

// postTodos adds a todo from JSON received in the request body.
func postTodos(c *gin.Context) {
	var newTodo todo

	// Call BindJSON to bind the received JSON to newTodo.
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	// check for ID duplication before adding newTodo.
	for _, t := range todos {
		if t.ID == newTodo.ID {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "ID already exist"})
			return
		}
	}

	// Add the new TODO to the slice.
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// getTodoByID locates the todo whose ID value matches the id parameter sent by the client, then returns that todo as a response.
func getTodoByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of todos, looking for a todo whose ID value matches the parameter.
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

// updateTodos updating todo whose ID matches the id parameter esent by the client.
func updateTodos(c *gin.Context) {
	id := c.Param("id")

	var newTodo todo

	// Call BindJSON to bind the received JSON to newTodo.
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	// Loop over the list of todos, looking for a todo whose ID value matches the parameter.
	for idx, a := range todos {
		if a.ID == id {
			if newTodo.Name != "" {
				todos[idx].Name = newTodo.Name
			}
			if newTodo.Checked != nil {
				todos[idx].Checked = newTodo.Checked
			}
			c.IndentedJSON(http.StatusOK, gin.H{"message": "todo's updated"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

// deleteTodos delete todo whose ID matches the id parameter by the client.
func deleteTodos(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of todos, looking for a todo whose ID value matches the parameter.
	for idx, a := range todos {
		if a.ID == id {
			todos = append(todos[:idx], todos[idx+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "todo's deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {
	router := gin.Default()

	// GET routes
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)

	// POST routes
	router.POST("/todos", postTodos)

	// PUT routes
	router.PUT("/todos/:id", updateTodos)

	// DELETE routes
	router.DELETE("/todos/:id", deleteTodos)

	router.Run("localhost:8080")
}
