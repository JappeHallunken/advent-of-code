package day8

import (
	"fmt"

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
  fmt.Println("input: ", input)
  fmt.Println("width: ", width)
  fmt.Println("height: ", height)
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
				fmt.Println("possible pos: ", possiblePositions)

				for _, pos := range possiblePositions {

					if pos.X < 0 || pos.Y < 0 || pos.X >= height || pos.Y >= width { //check if in bounds
						continue
					} else {

						if searchCoordinatesInMap(target, pos) { // look if there is already an antenna
							fmt.Println("here is already an antenna!, add anyway", pos)
              finalPos = append(finalPos, pos)
							continue
						}
						//check if there is already an antinode
						if searchCoordinatesInSlice(finalPos, pos) {
							fmt.Println("here is already an antinode!", pos)
							continue
						}
            fmt.Println("add: ", pos)
						finalPos = append(finalPos, pos)
					}
				}
            fmt.Println()

				//check if there's an antenna there

			}
			// diff :=
		}
	}
	return finalPos

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
	fmt.Println("day8:")

	body, err := fileops.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}

	// coordMap := fileops.ByteToCoordRuneMap(body)
	// fmt.Println(coordMap)

	lines := fileops.MakeStringSlice(body)
	fileops.PrintMap(lines)

	antennas := searchAntennas(lines)
	fmt.Println(antennas)

	list := checkAntinodes(antennas, lines)
  fmt.Println(list)
	fmt.Println(len(list))

	return len(list), 0

}
