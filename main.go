package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kidzrockboom/chess-review/api"
)

func main() {
	var username string
	var externalEngine int

	help := flag.String("help", "", "A simple program to get games from chess.com and use lichess analysis on them, to use the program, enter your username and external engine to use as analysis if you have one. Ex: chessview 'username' '5'")
	flag.Parse()

	fmt.Print("", *help)

	userInput := os.Args[1:]

	if len(userInput) < 1 || len(userInput) > 2 {
		fmt.Println("Invalid number of inputs")
		os.Exit(2)
	}

	username = userInput[0]

	if len(userInput) > 1 {
		externalEngine, err := strconv.Atoi(userInput[1])
		if err != nil {
			fmt.Printf("Improper value (%d) for chess engine \n", externalEngine)
			os.Exit(2)
		}
	}

	urlString := fmt.Sprintf("https://api.chess.com/pub/player/%s/games/archives", username)

	gamesArchive, err := api.GetGameArchive(urlString)
	if err != nil {
		log.Fatal(err)
	}

	games, err := api.GetChessGames(gamesArchive)
	if err != nil {
		log.Fatal(err)
	}

	gamesData, err := api.GetGamePgn(games)
	if err != nil {
		log.Fatal(err)
	}

	game, err := api.GetRecentGame(gamesData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(game)
	fmt.Printf("Username is: %s and number of games requested: %d \n", username, externalEngine)
}
