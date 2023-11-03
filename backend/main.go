package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// makling a default board that can be used on init and when starting a new game
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

// struct that will keep track of game boards
type battleshipBoard struct {
	PlayerBoard [10][10]string `json:"playerboard"`
	AiBoard     [10][10]string `json:"aiboard"`
}

// Structure to hold the player ship information
type playership struct {
	ShipType string `json:"shiptype"`
	Bowx     int    `json:"bowx"`
	Bowy     int    `json:"bowy"`
	Sternx   int    `json:"sternx"`
	Sterny   int    `json:"sterny"`
	Sunk     bool   `json:"sunk"`
}
type aiship struct {
	ShipType string `json:"shiptype"`
	Bowx     int    `json:"bowx"`
	Bowy     int    `json:"bowy"`
	Sternx   int    `json:"sternx"`
	Sterny   int    `json:"sterny"`
	Sunk     bool   `json:"sunk"`
}

// on the init of the program we will set the game boards to a default state
var battleships = []battleshipBoard{
	{PlayerBoard: defaultBoard, AiBoard: defaultBoard},
}

// init empty structs to track the ships
var playerships = []playership{}
var aiships = []aiship{}

func addPlayerShip(context *gin.Context) {
	var newShip playership

	if err := context.BindJSON(&newShip); err != nil {
		return
	}

	playerships = append(playerships, newShip)

	context.IndentedJSON(http.StatusCreated, playerships)
}

func getCurrentBoard(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, battleships)
}

func main() {
	router := gin.Default()
	router.GET("/getBoard", getCurrentBoard)
	router.GET("/todos/:id", getTodo)
	router.POST("/addShip/", addTodo)
	router.Run("localhost:9090")
}

// Every thing below here is no longer needed but is still being used for notes
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Dishes", Completed: false},
	{ID: "3", Item: "Walk Dog", Completed: false},
}

/*
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}*/

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
