package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

// read input to []rune
func readFileToRuneSlice(input string) []rune {
	body, err := fileops.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	bodyStr := string(body)
	bodyStr = strings.TrimSpace(bodyStr)
	runeSlice := []rune(bodyStr)
	return runeSlice
}

// make a string with free space marked as "."
func createBlockString(runeSlice []rune) []string {
	var slice2 []string
	for i, v := range runeSlice {

		if v < '0' || v > '9' {
			// check for non digits
			fmt.Println("non digit:", string(v))
			continue
		}

		//convert rune to int for next loop
		length, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println(i, v, err)
			return nil
		}
		if i%2 == 0 { //when index is even, append v-times the value of index to string
			for j := 0; j < length; j++ {
				slice2 = append(slice2, strconv.Itoa(i/2))
			}
		} else { //when uneven, append v-times just a "." to mark free space

			for k := 0; k < length; k++ {
				slice2 = append(slice2, ".")
			}
		}
	}
	// fmt.Print(slice2, "\n")
	return slice2
}

func fillEmptySpaces(slice []string) []string {
	counter := 0
	for i := range slice {
		if slice[i] == "." {

			counter++
		}
	}
	endIndex := len(slice) - counter

	for i := 0; i < endIndex; i++ {
		if slice[i] == "." {

			for j := len(slice) - 1; j >= 0; j-- {
				if slice[j] != "." {
					temp := slice[i]
					slice[i] = slice[j]
					slice[j] = temp
					break
				}
			}
		}
	}
	// fmt.Println(slice)
	return slice
}

func scanFreeSpace(input []string) []fileops.Coordinates {
  fmt.Println("scan free space")
	slice := input
	var coordinates []fileops.Coordinates
	i := 0

	for i < len(slice) {
    fmt.Println("i: ",i)
		if slice[i] == "." {
      fmt.Println("start gefunden: ", i)

			for j := i + 1; j < len(slice)-1; j++ {
				if slice[j] == "." {
          continue
				} else {

					coordinates = append(coordinates, fileops.Coordinates{X: i, Y: j - 1})
          
					fmt.Println("free space: ", coordinates)
					i = j - 1
					break
				}

			}
		}
      i++
	}
	return coordinates
}

func scanBlocks(slice []string) (blocks map[int]fileops.Coordinates) {
	//creates a map in the style of rune:{startIndex, endIndex}

	block := make(map[int]fileops.Coordinates)
	i := len(slice) - 1
	reference := slice[i]
	for i > 0 {
		count := 0
		if slice[i] != "." {
			reference = slice[i]

			for j := i - 1; j >= 0; j-- {
				if slice[j] == reference {
					count++
					// i--
				} else {
					break
				}
			}
			integer, err := strconv.Atoi(reference)
			if err != nil {
				fmt.Println(err)
				return
			}
			block[integer] = fileops.Coordinates{X: i - count, Y: i}
			i = i - count
		}
		i--
	}
	// fmt.Println("blocks: start, end", block)
	return block
}

func defrag(original []string) (defragedString []string) {

	copySlice := make([]string, len(original))
	copy(copySlice, original)

	dataBlocks := scanBlocks(original)
	freeSpace := scanFreeSpace(original)

	// fmt.Println("lenght original: ", length)
	for i := range dataBlocks {

		// fmt.Printf("defrag: %v / %v\n\n", i, len(dataBlocks))
		blockWidth := dataBlocks[i].X - dataBlocks[i].Y + 1

		for j := range freeSpace {
			freeWidth := freeSpace[j].Y - freeSpace[j].X + 1
			filled := false
			if blockWidth <= freeWidth {
				start := freeSpace[j].X

				for k := range blockWidth {

					if start >= dataBlocks[i].Y {
						break
					}
					copySlice[start+k] = original[dataBlocks[i].Y]
					copySlice[dataBlocks[i].Y+k] = "."
					// fmt.Println(copySlice)

				}

				filled = true
				// freeSpace = scanFreeSpace(copySlice)

			}
			if filled {
				break
			}
		}
	}
	return copySlice
}

func calculateChecksum(slice []string) (checksum int) {
	for i, v := range slice {
		value, _ := strconv.Atoi(v)
		checksum += i * value
	}
	return checksum
}

func Day9(input string) (checksum, checksum2 int) {
	readFileToRuneSlice(input)
	slice := createBlockString(readFileToRuneSlice(input))
	defraggedSlice := fillEmptySpaces(slice)

	checksum = calculateChecksum(defraggedSlice)

	slice = createBlockString(readFileToRuneSlice(input))
	blockSlice := defrag(slice)
	fmt.Printf("\n %v \n", blockSlice)
	checksum2 = calculateChecksum(blockSlice)

	return checksum, checksum2
}
