package day7

import "strings"

type Coord struct {
	Row int
	Col int
}

func P1(input string) int {
	var r, c int // number of rows and cols of the input

	s := Coord{}
	lines := strings.Split(input, "\n")
	r = len(lines)
	c = len(lines[0])
	s.Row = 0 // row coordinate of 'S'

	for i, ch := range lines[0] {
		if ch == 'S' {
			s.Col = i // col coordinate of 'S'
		}
	}

	active := make([]Coord,2)
	active = append(active, Coord{s.Row+1, s.Col}) // the first beam is on row under the 'S'

	
}
