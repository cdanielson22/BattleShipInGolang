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
type ship struct {
	ShipType string `json:"shiptype"`
	Bowx     int    `json:"bowx"`
	Bowy     int    `json:"bowy"`
	Sternx   int    `json:"sternx"`
	Sterny   int    `json:"sterny"`
	Sunk     bool   `json:"sunk"`
	Player   string `json:"playerType"`
}

// on the init of the program we will set the game boards to a default state
var battleships = []battleshipBoard{
	{PlayerBoard: defaultBoard, AiBoard: defaultBoard},
}

// init empty structs to track the ships
var ships = []ship{}

func checkShipOverlap(start int, end int, length int) bool {
	return false
}

func validateShip(validShip ship) bool {
	// first check if its with in the bounds of the board
	// check if its horizonal or vertical
	// then check if it touches other ships

	if (validShip.Bowx < 0) || (validShip.Bowx > 9) {
		return false
	}
	if (validShip.Bowy < 0) || (validShip.Bowy > 9) {
		return false
	}
	if (validShip.Sternx < 0) || (validShip.Sternx > 9) {
		return false
	}
	if (validShip.Sterny < 0) || (validShip.Sterny > 9) {
		return false
	}

	if validShip.Bowx != validShip.Sternx {
		// return checkShipOverlap()
	} else {
		// we have a vertical ship
	}

	return true
}

func addShipToBoard(shipToAdd ship) {

}

func addShip(context *gin.Context) {
	var newShip ship

	if err := context.BindJSON(&newShip); err != nil {
		return
	}

	ships = append(ships, newShip)

	//I'll need to validate ships position and then add to board
	// addShipToBoard(newShip)

	context.IndentedJSON(http.StatusCreated, ships)
}

func getShips(context *gin.Context) {
	context.JSON(http.StatusOK, ships)
}

func getCurrentBoard(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, battleships)
}

func main() {
	router := gin.Default()
	router.GET("/getBoard", getCurrentBoard)
	router.GET("/getPlayerShips", getShips)
	router.GET("/todos/:id", getTodo)
	router.POST("/addShip", addShip)
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
