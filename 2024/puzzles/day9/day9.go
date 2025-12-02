package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JappeHallunken/advent-of-code/2024/fileops"
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
	// fmt.Println(string(runeSlice))
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
	// for i, v := range slice2 {
	//   fmt.Println(i, v)
	// }
	// fmt.Println(slice2)
	return slice2
}

func fillEmptySpaces(slice []string) []string {
	counter := 0
	for i := range slice {
		if slice[i] == "." { //count the "."
			counter++
		}
	}
	endIndex := len(slice) - counter

	for i := 0; i < endIndex; i++ {
		if slice[i] == "." {

			for j := len(slice) - 1; j >= 0; j-- {
				if slice[j] != "." {
					temp := slice[i] // swap the positions
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
	slice := input
	var coordinates []fileops.Coordinates
	i := 0

	for i < len(slice) {
		// fmt.Println("i: ",i)
		if slice[i] == "." {
			for j := i + 1; j < len(slice)-1; j++ {
				if slice[j] == "." {
					continue
				} else {
					coordinates = append(coordinates, fileops.Coordinates{X: i, Y: j - 1})
					i = j - 1
					break
				}
			}
		}
		i++
	}

	// fmt.Println("free space: ", coordinates)
	return coordinates
}

//creates a map of occupied space in the style of rune:{startIndex, endIndex}
func scanBlocks(slice []string) (blocks map[int]fileops.Coordinates) {

	block := make(map[int]fileops.Coordinates)
	i := len(slice) - 1

  // set a reference for comparison
	reference := slice[i]
	for i > 0 {
		count := 0
		if slice[i] != "." {
			reference = slice[i]

			for j := i - 1; j >= 0; j-- {
				if slice[j] == reference {
					count++
          // fmt.Println(i, j, count)
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
	// fmt.Println("scan blocks: ", block)
	return block
}

func defrag(original []string) (defragedString []string) {

	copySlice := make([]string, len(original))
	copy(copySlice, original)
	// make the slice/map for free and occupied space
	dataBlocks := scanBlocks(original)
	freeSpace := scanFreeSpace(original)

	// find max key in map for backwards iteration
	var maxKey int
	firstIteration := true
	// find datablock with highest key
	for key := range dataBlocks {
		if firstIteration || key > maxKey {
			maxKey = key
			firstIteration = false
		}
	}
	startSearchSpace := 0
  // iterate von block with highest key to lowest
	for i := maxKey; i >= 0; i-- {
		nextBlock := false
		//calculate width of datablock
		blockWidth := dataBlocks[i].Y - dataBlocks[i].X

		// range over free space
		for j := startSearchSpace; j < len(freeSpace) && !nextBlock; j++ {

			start := freeSpace[j].X
			oldStart := dataBlocks[i].X
      // when new position would be behind the old one -> break
			if start > oldStart {
				break
			}
			//calc space width
			spaceWidth := freeSpace[j].Y - freeSpace[j].X
			//if the data fits into the free space, move it
			if (blockWidth <= spaceWidth) && spaceWidth > -1 {
				for k := 0; k < blockWidth+1; k++ {
					copySlice[start+k] = copySlice[dataBlocks[i].X+k]
					copySlice[dataBlocks[i].X+k] = "."
				}
				// update the free space coordinates
				freeSpace[j].X += blockWidth + 1
				nextBlock = true
			}
		}
	}

	// fmt.Println(copySlice)
	return copySlice
}

func calculateChecksum(slice []string) (checksum int) {
	for i, v := range slice {
		// fmt.Printf("%v / %v\n", i, len(slice))
		value, _ := strconv.Atoi(v)
		checksum += i * value
	}
	return checksum
}

func Day9(input string) (checksum, checksum2 int) {
	readFileToRuneSlice(input)
	slice := createBlockString(readFileToRuneSlice(input))
	// fmt.Println("part 1 block: ", slice)

	defraggedSlice := fillEmptySpaces(slice)
	checksum = calculateChecksum(defraggedSlice)

	slice2 := createBlockString(readFileToRuneSlice(input))
	blockSlice := defrag(slice2)
	// fmt.Printf("\n %v \n", blockSlice)
	checksum2 = calculateChecksum(blockSlice)

	return checksum, checksum2
}
