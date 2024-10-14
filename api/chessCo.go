package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kidzrockboom/chess-review/models"
)

type Archive struct {
	GamesList []string `json:"archives"`
}

func GetChessGames(urlString string) ([]models.Game, error) {
	var games []models.Game

	var jsonData Archive

	res, err := http.Get(urlString)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println("Json unmarshall err")
		log.Fatal(err)
	}

	fmt.Println(jsonData.GamesList)

	// Store archive list of games locally

	// Get games from archive list
	res, err = http.Get(jsonData.GamesList[len(jsonData.GamesList)-1])
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	return games, nil
}
