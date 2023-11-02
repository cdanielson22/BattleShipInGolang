package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var defaultBoard = [10][10]string{
	{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
	{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
	{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
	{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
	{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
	{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
	{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
	{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
	{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
	{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
}

type battleship struct {
	Board [10][10]string `json:"board"`
}

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

func getCurrentBoard(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, battleships)
}

var battleships = []battleship{
	{Board: defaultBoard},
}

// Every thing below here is no longer needed but is still being used for notes

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Dishes", Completed: false},
	{ID: "3", Item: "Walk Dog", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/getBoard", getCurrentBoard)
	router.GET("/todos/:id", getTodo)
	router.POST("/todo", addTodo)
	router.Run("localhost:9090")
}