package day8

import (
	"fmt"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

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

func checkAntinodes(antennaMap map[rune][]fileops.Coordinates) {

	for _, freqs := range antennaMap {
		for i := range freqs {
			for j := i + 1; j < len(freqs); j++ {
				first := freqs[i]
				second := freqs[j]

        diff := fileops.Coordinates{X: first.X - second.X, Y: first.Y - second.Y}
        diff2 := fileops.Coordinates{X: -diff.X, Y: -diff.Y}

				fmt.Println(diff)
				fmt.Println(diff2)

			}
			// diff :=
		}
	}

}

func searchCoordinatesInMap(antennaMap map[rune][]fileops.Coordinates, targetCoord fileops.Coordinates) bool {
	for _, freqs := range antennaMap {
		for _, coord := range freqs {
			if coord == targetCoord {
				return true
			}
		}
	}
	return false
}
func Day8(input string) int {
	fmt.Println("day8:")

	body, err := fileops.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	// coordMap := fileops.ByteToCoordRuneMap(body)
	// fmt.Println(coordMap)

	lines := fileops.MakeStringSlice(body)
	fileops.PrintMap(lines)

	antennas := searchAntennas(lines)
	fmt.Println(antennas)

	return 0

}
