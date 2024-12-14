package day13

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Game struct {
	AX int
	AY int
	BX int
	BY int
	PX int
	PY int
}

func createGame(text string) Game {
	aPattern := regexp.MustCompile(`Button A: X\+(\d{2}), Y\+(\d{2})`)
	bPattern := regexp.MustCompile(`Button B: X\+(\d{2}), Y\+(\d{2})`)
	pPattern := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	game := Game{}

	matchA := aPattern.FindStringSubmatch(text)
	if len(matchA) == 3 {
		// Extrahiere X und Y für Button A
		fmt.Sscanf(matchA[1], "%d", &game.AX)
		fmt.Sscanf(matchA[2], "%d", &game.AY)
	}

	// Zeile für Button B
	matchB := bPattern.FindStringSubmatch(text)
	if len(matchB) == 3 {
		// Extrahiere X und Y für Button B
		fmt.Sscanf(matchB[1], "%d", &game.BX)
		fmt.Sscanf(matchB[2], "%d", &game.BY)
	}

	// Zeile für Prize
	matchP := pPattern.FindStringSubmatch(text)
	if len(matchP) == 3 {
		// Extrahiere X und Y für Prize
		fmt.Sscanf(matchP[1], "%d", &game.PX)
		fmt.Sscanf(matchP[2], "%d", &game.PY)
	}
	return game
}
func createGames(text string) []Game {
	text = strings.TrimSpace(text)
	sections := strings.Split(text, "\n\n")
	game := Game{}
	games := []Game{}

	for i, section := range sections {
		fmt.Printf("Section %d:\n%s\n\n", i+1, section)

		game = createGame(section)
		games = append(games, game)

	}

	// Ausgabe des extrahierten Spiels
	// fmt.Printf("Game struct: %+v\n", game)

	return games
}
func Day13(input string) (p1 int) {

	body, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("File reading error", err)
		return 0
	}

	games := createGames(string(body))
	for i, game := range games {
		fmt.Printf("Game %v:\nAX: %v, AY: %v, BX: %v, BY: %v, PX: %v, PY: %v\n", i+1, game.AX, game.AY, game.BX, game.BY, game.PX, game.PY)
	}

	return 0
}
