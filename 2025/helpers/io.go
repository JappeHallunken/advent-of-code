// Package helpers contains functions to read in files and return the types needed for solving the puzzles
package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileToString(path string) ([]string, error) {
	readFile, err := os.Open(path)
	var s []string

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s = append(s, fileScanner.Text())
	}
	return s, nil
}
