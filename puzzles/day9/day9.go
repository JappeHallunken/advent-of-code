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

// make a string with free space as .
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
		} else { //when uneven, append v-times just a "." to mark fress space

			for k := 0; k < length; k++ {
				slice2 = append(slice2, ".")
			}
			// fmt.Println(slice2)
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
	// fmt.Println("checksum:", checksum)

	return checksum
}
