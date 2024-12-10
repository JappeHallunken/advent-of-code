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
	// fmt.Println("scan blocks: ", block)
	return block
}

func defrag(original []string) (defragedString []string) {

	copySlice := make([]string, len(original))
	copy(copySlice, original)

	dataBlocks := scanBlocks(original)
	freeSpace := scanFreeSpace(original)

	// fmt.Println("copy: ", copySlice)
	// fmt.Println("data blocks: ", dataBlocks)
	// fmt.Println("free space: ", freeSpace)

	// find max key in map for backwards iteration
	var maxKey int
	firstIteration := true

	for key := range dataBlocks {
		if firstIteration || key > maxKey {
			maxKey = key
			firstIteration = false
		}
	}
	// fmt.Println("maxKey: ", maxKey)

	startSearchSpace := 0
	// startSearchBlock := maxKey

	for i := maxKey; i >= 0; i-- {
		nextBlock := false
		//calculate wodth of datablock
		blockWidth := dataBlocks[i].Y - dataBlocks[i].X

		// fmt.Printf("\niteration: %d\n", i)
		// fmt.Println("block width: ", blockWidth)

		// range over free space
		for j := startSearchSpace; j < len(freeSpace) && !nextBlock; j++ { //range freeSpace {

			start := freeSpace[j].X
			oldStart := dataBlocks[i].X
			if start > oldStart {
				break
			}

			//calc space width
			spaceWidth := freeSpace[j].Y - freeSpace[j].X
			// fmt.Printf("calc %v - %v = %v\n", freeSpace[j].Y, freeSpace[j].X, spaceWidth)
			// fmt.Printf("found free space at: %v, width: %v\n", freeSpace[j].X, spaceWidth)

			//if the data fits into the free space, move it
			if (blockWidth <= spaceWidth) && spaceWidth > -1 {
				// fmt.Println("space is big enough! -> move")
				// fmt.Println("select start coord: ", start)
				// fmt.Println("select end coord: ", start+blockWidth)

				for k := 0; k < blockWidth+1; k++ {
					copySlice[start+k] = copySlice[dataBlocks[i].X+k]
					copySlice[dataBlocks[i].X+k] = "."

					// fmt.Printf("for block %d: %v \n", i, copySlice)
					// fmt.Println("free space: ", freeSpace)

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
