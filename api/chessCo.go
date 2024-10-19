package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/kidzrockboom/chess-review/models"
)

type Archive struct {
	GamesList []string `json:"archives"`
}

type Games struct {
	List []models.Game `json:"games"`
}

func GetGameArchive(urlString string) ([]string, error) {
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

	// Store archive list of games locally

	return jsonData.GamesList, nil
}

func GetChessGames(archiveList []string) ([]models.Game, error) {
	var games Games

	// Get games from archive list
	res, err := http.Get(archiveList[len(archiveList)-1])
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Get chess Games")
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &games)
	if err != nil {
		fmt.Println("Json unmarshall err")
		log.Fatal(err)
	}

	return games.List, nil
}

func GetGamePgn(gamesList []models.Game) ([]string, error) {
	var pgnList []string

	for i := 0; i < len(gamesList); i++ {
		temp := strings.Split(gamesList[i].Pgn, "\n")
		re := regexp.MustCompile(`"[^"]+"`)
		date := re.FindAllString(temp[1], -1)
		utcTime := re.FindAllString(temp[11], -1)
		timeControl := re.FindAllString(temp[14], -1)
		gamePgn := re.FindAllString(temp[19], -1)
	}
	return pgnList, nil
}
