package day13

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Game struct {
	AX         int
	AY         int
	BX         int
	BY         int
	PX         int
	PY         int
	legitGames []Pushs
	minPrize   int
}
type Pushs struct {
	A int
	B int
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

	for _, section := range sections {
		// fmt.Printf("Section %d:\n%s\n\n", i+1, section)

		game = createGame(section)
		games = append(games, game)

	}

	// Ausgabe des extrahierten Spiels
	// fmt.Printf("Game struct: %+v\n", game)

	return games
}
func addToXAndY(game *Game) {

	game.PX = game.PX + 10000000000000
	game.PY = game.PY + 10000000000000
}
func calculateGame(game *Game) {
	var legitGame Pushs
	expectedX := game.PX
	expectedY := game.PY
	// limit := game.PX / game.AX

	for a := 0; a*game.AX <= game.PX; a++ {
		if a%100 == 0 {
			fmt.Println("A: ", a)
		}
		if a*game.AX > expectedX || a*game.AY > expectedY {
			break
		}
		remaining := expectedX - (a * game.AX)
		if remaining%game.BX == 0 {
			b := remaining / game.BX

			if a*game.AY+b*game.BY == expectedY {
				fmt.Println("treffer: ", a, b)
				legitGame.A = a
				legitGame.B = b

				game.legitGames = append(game.legitGames, legitGame)
			}
		}

	}
}

func calcMinPrize(game *Game) {
	minPrize := 0
	for _, push := range game.legitGames {
		prize := calculatePrize(push)
		if minPrize == 0 || prize < minPrize {
			minPrize = prize
			game.minPrize = minPrize
		}
	}
}

func calculatePrize(push Pushs) int {
	return (push.A * 3) + push.B
}

func findButtonPresses(game Game) {
	// Werte aus der Aufgabe
	targetX := game.PX
	targetY := game.PY
	aX, aY := game.AX, game.AY
	bX, bY := game.BX, game.BY

	// Iteriere über mögliche Werte von i
	for i := 0; i*aX <= targetX; i++ {
		if i%100000000 == 0 {
			fmt.Println("i: ", i)
		}
		// Berechne j aus der ersten Gleichung
		remainingX := targetX - i*aX
		if remainingX%bX != 0 {
			continue // Überspringe, wenn keine gültige Lösung für j
		}
		j := remainingX / bX

		// Überprüfe, ob die zweite Gleichung ebenfalls erfüllt ist
		if i*aY+j*bY == targetY {
			fmt.Printf("Lösung gefunden: i = %d, j = %d\n", i, j)
		}
	}
}

// Erweiterter Euklidischer Algorithmus
// Gibt den größten gemeinsamen Teiler und die Koeffizienten für ax + by = gcd(a, b) zurück
func extendedGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

// Lösung der linearen Diophantischen Gleichung: ax + by = c
func solveLinearDiophantine(a, b, c int) (int, int, int, bool) {
	// Berechne den ggT der beiden Zahlen a und b
	gcd, x, y := extendedGCD(a, b)

	// Überprüfen, ob die rechte Seite durch den ggT teilbar ist
	if c%gcd != 0 {
		return 0, 0, 0, false // Keine Lösung
	}

	// Lösung finden, indem wir den Wert der rechten Seite durch den ggT teilen
	x *= c / gcd
	y *= c / gcd

	return x, y, gcd, true
}

// Hauptlogik zur Berechnung von i und j
func solveSystem(game *Game) {

	var legitGame Pushs
	a1, b1, c1, a2, b2, c2 := game.AX, game.BX, game.PX, game.AY, game.BY, game.PY
	// Solve the first equation: a1 * a + b1 * b = c1
	x1, y1, gcd1, ok1 := solveLinearDiophantine(a1, b1, c1)
	if !ok1 {
		return
	}

	k := a2*(b1/gcd1) - b2*(a1/gcd1)
	rhs := c2 - a2*x1 - b2*y1

	// Solve t * k = rhs
	if k == 0 {
		if rhs == 0 {
			// Infinite solutions for t
			return
		} else {
			// No solution
			return
		}
	}

	// Solve for t
	if rhs%k != 0 {
		return // No integer solution for t
	}

	t := rhs / k

	// Substitute t back to find a and b
	legitGame.A = x1 + t*(b1/gcd1)
	legitGame.B = y1 - t*(a1/gcd1)
	game.legitGames = append(game.legitGames, legitGame)

	return
}
func Day13(input1, input2 string) (p1, p2 int) {
	// puzzle 1
	body, err := os.ReadFile(input1)
	if err != nil {
		fmt.Println("File reading error", err)
		return 0, 0
	}

	games := createGames(string(body))
	for i, game := range games {
		calculateGame(&game)

		calcMinPrize(&game)
		p1 += game.minPrize

		fmt.Printf("\nGame %v:\nAX: %v, AY: %v,\nBX: %v, BY: %v,\nPX: %v, PY: %v,\nlegitGames: %v,\nminPrize: %v\n", i+1, game.AX, game.AY, game.BX, game.BY, game.PX, game.PY, game.legitGames, game.minPrize)
	}

	//   puzzle 2
	body2, err := os.ReadFile(input1)
	if err != nil {
		fmt.Println("File reading error", err)
		return 0, 0
	}

	games2 := createGames(string(body2))

	for l, game := range games2 {
		addToXAndY(&game)

		// a1, b1, c1 := game.AX, game.BX, game.PX
		// a2, b2, c2 := game.AY, game.BY, game.PY //
		// Lösungen finden
		solveSystem(&game)

		fmt.Printf("\nGame %v:\nAX: %v, AY: %v,\nBX: %v, BY: %v,\nPX: %v, PY: %v,\nlegitGames: %v,\nminPrize: %v\n", l+1, game.AX, game.AY, game.BX, game.BY, game.PX, game.PY, game.legitGames, game.minPrize)
		// if valid {
		// 	fmt.Printf("Lösungen gefunden: i = %d, j = %d\n", i, j)
		// } else {
		// 	fmt.Println("Keine Lösung gefunden.")
		// }
		calcMinPrize(&game)
		p2 += game.minPrize

	}

	return p1, p2
}
