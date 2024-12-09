package day8

import (
	"fmt"
	// "time"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

func getMapDimensions(lines []string) (width int, height int) {
	width = len(lines[0])
	height = len(lines)
	return width, height
}

func searchAntennas(lines []string) map[rune][]fileops.Coordinates {
	// create a map where the key is the frequencie and the value is a slice of fileops.Coordinates

	antennaMap := make(map[rune][]fileops.Coordinates)
	for i, line := range lines {
		for j, char := range line {

			if lines[i][j] != '.' {

				if _, exists := antennaMap[char]; exists {

					antennaMap[char] = append(antennaMap[char], fileops.Coordinates{X: i, Y: j})
				} else {
					antennaMap[char] = []fileops.Coordinates{{X: i, Y: j}}
				}
			}
		}
	}

	return antennaMap
}

func checkAntinodes(target map[rune][]fileops.Coordinates, input []string) []fileops.Coordinates {
	width, height := getMapDimensions(input)
	// fmt.Println("input: ", input)
	// fmt.Println("width: ", width)
	// fmt.Println("height: ", height)
	var finalPos []fileops.Coordinates
	for _, freqs := range target {
		for i := range freqs {
			for j := i + 1; j < len(freqs); j++ {
				first := freqs[i]
				second := freqs[j]

				diff := fileops.Coordinates{X: first.X - second.X, Y: first.Y - second.Y}

				invertDiff := fileops.Coordinates{X: -diff.X, Y: -diff.Y}

				// fmt.Println(diff)
				// fmt.Println(diff2)
				// get potential antinode positions
				var possiblePositions []fileops.Coordinates
				pos1 := fileops.Coordinates{X: first.X + diff.X, Y: first.Y + diff.Y}
				pos2 := fileops.Coordinates{X: second.X + invertDiff.X, Y: second.Y + invertDiff.Y}
				possiblePositions = append(possiblePositions, pos1, pos2)
				// fmt.Println("possible pos: ", possiblePositions)

				for _, pos := range possiblePositions {

					if pos.X < 0 || pos.Y < 0 || pos.X >= height || pos.Y >= width { //check if in bounds
						continue
					} else {

						if searchCoordinatesInMap(target, pos) { // look if there is already an antenna
							// fmt.Println("here is already an antenna!, add anyway", pos)
							finalPos = append(finalPos, pos)
							continue
						}
						//check if there is already an antinode
						if searchCoordinatesInSlice(finalPos, pos) {
							// fmt.Println("here is already an antinode!", pos)
							continue
						}
						// fmt.Println("add: ", pos)
						finalPos = append(finalPos, pos)
					}
				}

				//check if there's an antenna there

			}
			// diff :=
		}
	}
	return finalPos

}

func checkAntinodes2(target map[rune][]fileops.Coordinates, input []string) []fileops.Coordinates {
	width, height := getMapDimensions(input)
	// fmt.Println("input: ", input)
	// fmt.Println("width: ", width)
	// fmt.Println("height: ", height)
	var finalPos2 []fileops.Coordinates
	for _, freqs := range target {
		// fmt.Println("freq: ", string(k), freqs)
		for i := range freqs {
			// fmt.Println("i: ", i)
			for j := range freqs {
				// fmt.Println("j: ", j)
				if len(freqs) == 1 {
					continue
				}
				if i == j {
					continue
				}

        // fmt.Println("positions: ", finalPos2)
				first := freqs[i]
				second := freqs[j]
				// fmt.Println("first pos: ", first, "second pos: ", second)

				if !searchCoordinatesInSlice(finalPos2, second) {

					finalPos2 = append(finalPos2, second) //

				}

				diff := fileops.Coordinates{X: second.X - first.X, Y: second.Y - first.Y}
				// fmt.Println("diff: ", diff)

				// fmt.Println("possible pos: ", possiblePositions)
				currentPos, currentPos2 := second, second

				for {
					currentPos.X += diff.X
					currentPos.Y += diff.Y
					// fmt.Println("cur pos: ", currentPos)
					// time.Sleep(5 * time.Millisecond)

					if currentPos.X < 0 || currentPos.X >= height || currentPos.Y < 0 || currentPos.Y >= width { //check if in bounds
						break
					}

					//check if there is already an antinode
					if searchCoordinatesInSlice(finalPos2, currentPos) {
						// fmt.Println("here is already an antinode!", currentPos)
						continue
					}
					// fmt.Println("add: ", currentPos)
					finalPos2 = append(finalPos2, currentPos)
				}
				for {
					currentPos2.X -= diff.X
					currentPos2.Y -= diff.Y
					// fmt.Println("cur pos: ", currentPos)
					// time.Sleep(500 * time.Millisecond)

					if currentPos2.X < 0 || currentPos2.X >= height || currentPos2.Y < 0 || currentPos2.Y >= width { //check if in bounds
						break
					}

					//check if there is already an antinode
					if searchCoordinatesInSlice(finalPos2, currentPos2) {
						// fmt.Println("here is already an antinode!", currentPos2)
						continue
					}
					// fmt.Println("add: ", currentPos2)
					finalPos2 = append(finalPos2, currentPos2)
				}

				//check if there's an antenna there

			}
		}
		// diff :=
	}

	return finalPos2

}

func searchCoordinatesInMap(input map[rune][]fileops.Coordinates, targetCoord fileops.Coordinates) bool {
	for _, freqs := range input {
		for _, coord := range freqs {
			if coord == targetCoord {
				return true
			}
		}
	}
	return false
}

func searchCoordinatesInSlice(slice []fileops.Coordinates, targetCoord fileops.Coordinates) bool {
	for _, coord := range slice {
		if coord == targetCoord {
			return true
		}
	}
	return false
}
func Day8(input string) (puzzle1, puzzle2 int) {

	body, err := fileops.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}

	// coordMap := fileops.ByteToCoordRuneMap(body)
	// fmt.Println(coordMap)

	lines := fileops.MakeStringSlice(body)
	// fileops.PrintMap(lines)

	antennas := searchAntennas(lines)
	// fmt.Println(antennas)

	list := checkAntinodes(antennas, lines)
	// fmt.Println(len(list))
	// fmt.Println("puzzle 1 done")
	// fmt.Println(antennas)
	// fmt.Println(len(list))

	list2 := checkAntinodes2(antennas, lines)

	return len(list), len(list2)

}
