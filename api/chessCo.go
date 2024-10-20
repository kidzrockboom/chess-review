package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kidzrockboom/chess-review/models"
	"golang.org/x/exp/slices"
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

func GetGamePgn(gamesList []models.Game) ([]models.GameData, error) {
	var pgnList []models.GameData

	layoutDate := "2006.01.02"
	layoutTime := "15:04:05"

	for i := 0; i < len(gamesList); i++ {
		var game models.GameData

		temp := strings.Split(gamesList[i].Pgn, "\n")
		re := regexp.MustCompile(`"[^"]+"`)

		tempDate := re.FindAllString(temp[2], -1)[0]
		tempDate = strings.Trim(tempDate, "\"")

		date, err := time.Parse(layoutDate, tempDate)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			os.Exit(1)
		}
		game.Date = date

		tempUtc := re.FindAllString(temp[11], -1)[0]
		tempUtc = strings.Trim(tempUtc, "\"")

		utcTime, err := time.Parse(layoutDate, tempUtc)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			os.Exit(1)
		}
		game.UtcTime = utcTime

		game.TimeControl = strings.Trim(re.FindAllString(temp[15], -1)[0], "\"")
		game.Pgn = temp[22]

		tempStart := re.FindAllString(temp[17], -1)[0]
		tempStart = strings.Trim(tempStart, "\"")
		startTime, err := time.Parse(layoutTime, tempStart)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			os.Exit(1)
		}
		game.StartTime = startTime

		tempEnd := re.FindAllString(temp[19], -1)[0]
		tempEnd = strings.Trim(tempEnd, "\"")
		endTime, err := time.Parse(layoutTime, tempEnd)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			os.Exit(1)
		}
		game.EndTime = endTime
		pgnList = append(pgnList, game)
	}
	return pgnList, nil
}

func GetRecentGame(gameList []models.GameData) (models.GameData, error) {
	// Find all where the time control is 600 or greater
	var rapidGames []models.GameData

	for i := 0; i < len(gameList); i++ {
		timeControl, err := strconv.Atoi(gameList[i].TimeControl)
		if err != nil {
			fmt.Println("Error converting timeControl", err)
			os.Exit(1)
		}
		if timeControl >= 600 {
			rapidGames = append(rapidGames, gameList[i])
		}
	}

	// Sort games into most recently played
	dateCmp := func(a, b models.GameData) int {
		return b.Date.Compare(a.Date)
	}

	slices.SortFunc(rapidGames, dateCmp)

	recentDay := rapidGames[0].Date
	var recentGamesDay []models.GameData

	for j := 0; j < len(rapidGames); j++ {
		if rapidGames[j].Date.Equal(recentDay) {
			recentGamesDay = append(recentGamesDay, rapidGames[j])
		} else {
			break
		}
	}

	// Find most recent game ended
	timeCmp := func(a, b models.GameData) int {
		return b.EndTime.Compare(a.EndTime)
	}

	slices.SortFunc(recentGamesDay, timeCmp)

	return rapidGames[0], nil
}
