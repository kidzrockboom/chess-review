package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/kidzrockboom/chess-review/apis"
)

func main() {
	var username string
	var numOfGames int

	help := flag.String("help", "", "A simple program to get games from chess.com and use lichess analysis on them, to use the program, enter your username and number of games wanted. Ex: chessview 'username' '5'")
	flag.Parse()

	fmt.Print("", *help)

	userInput := os.Args[1:]

	if len(userInput) < 2 || len(userInput) > 2 {
		fmt.Println("Invalid number of inputs")
		os.Exit(2)
	}

	username = userInput[0]
	numOfGames, err := strconv.Atoi(userInput[1])
	if err != nil {
		fmt.Printf("Improper value (%s) for number of games \n", userInput[1])
		os.Exit(2)
	}

	urlString := fmt.Sprintf("https://api.chess.com/pub/player/%s/games/archives", username)

	api.GetChessGames(urlString)

	fmt.Printf("Username is: %s and number of games requested: %d \n", username, numOfGames)
}
