package day9

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
			// fmt.Println(slice2)
		}
	}
	fmt.Print(slice2, "\n")
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
	// fmt.Println(counter)
	// fmt.Println(len(slice))

	for i := 0; i < endIndex; i++ {
		if slice[i] == "." {
			// fmt.Println("found . at:", i)

			for j := len(slice) - 1; j >= 0; j-- {
				if slice[j] != "." {
					// fmt.Println("found ", slice[j], " at:", j)
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
	slice := input
	var coordinates []fileops.Coordinates
	i := 0

	for i < len(slice) {
		count := 0
		if slice[i] == "." {
			// fmt.Println("start found . at:", i)
			for j := i + 1; j < len(slice)-1; j++ {
				if slice[j] == "." {

					count++
				} else {
					// fmt.Println("no . at: ", j)
					//      fmt.Println("save to coordinates: ", i, i+count)

					coordinates = append(coordinates, fileops.Coordinates{X: i, Y: i + count})
					// fmt.Println("coordinates: ", coordinates)
					i = i + count
					break
				}

			}
		}
		i++
	}
	fmt.Println(coordinates)
	return coordinates
}

func scanBlocks(slice []string) []fileops.Coordinates {

	var coordinates2 []fileops.Coordinates
	i := len(slice) - 1
	reference := slice[i]
	for i > 0 {
		count := 0
		if slice[i] != "." {
			reference = slice[i]
			fmt.Println("ref: ", reference)

			for j := i - 1; j >= 0; j-- {
				fmt.Println("target:  ", slice[j])
				if slice[j] == reference {
					count++
				} else {
					fmt.Println("no . at: ", j)
					fmt.Println("save to coordinates: ", i, i+count)
					break
				}
			}
			coordinates2 = append(coordinates2, fileops.Coordinates{X: i, Y: i - count})
			fmt.Println("coordinates: ", coordinates2)
			fmt.Println("current i: ", i)
			i = i - count
			fmt.Println("new i: ", i)
		}
		i--
		fmt.Println("new i--: ", i)
	}
	fmt.Println(coordinates2)
	return coordinates2
}

func fillEmptySpaces2(slice []string) []string {
	// fmt.Println(counter)
	// fmt.Println(len(slice))
	reference := slice[len(slice)-1]
	counter := 0
	fmt.Printf("Start! reference: %v, counter: %v\n\n", reference, counter)

	for i := len(slice) - 1; i >= 0; i-- {
		time.Sleep(100 * time.Millisecond)

		if slice[i] == reference {
			counter++

			fmt.Printf("%v; reference: %v; counter: %v\n", i, reference, counter)
		} else {

			if slice[i] == "." {
				counter = 0
				continue
			}
			reference = slice[i]

			counter = 1

			fmt.Printf("%v; reference: %v; counter: %v\n", i, reference, counter)
		}
	}
	// for i := range slice {
	// 	if slice[i] == "." {
	// 		// fmt.Println("found . at:", i)
	//
	// 		for j := len(slice) - 1; j >= 0; j-- {
	// 			if slice[j] != "." {
	// 				// fmt.Println("found ", slice[j], " at:", j)
	// 				temp := slice[i]
	// 				slice[i] = slice[j]
	// 				slice[j] = temp
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Printlwan(slice)
	return slice
}

func calculateChecksum(slice []string) (checksum int) {
	for i, v := range slice {
		value, _ := strconv.Atoi(v)
		checksum += i * value
	}
	return checksum
}

func Day9(input string) (checksum int) {
	readFileToRuneSlice(input)
	slice := createBlockString(readFileToRuneSlice(input))

	defraggedSlice := fillEmptySpaces(slice)

	checksum = calculateChecksum(defraggedSlice)
	fmt.Println("checksum:", checksum)

	slice = createBlockString(readFileToRuneSlice(input))
	fmt.Println(len(slice))
	_ = scanFreeSpace(slice)
	_ = scanBlocks(slice)

	// fmt.Println("coordinates:", coordinates)
	// checksum = calculateChecksum(coordinates)

	return checksum
}
