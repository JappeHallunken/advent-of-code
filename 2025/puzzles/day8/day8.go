package day8

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
)

var TopN = 1000 // package variable to make it modifiable in the test

type Coord [3]int

type Box struct {
	ID    int
	Coord Coord
}

type DistancePair struct {
	I, J     int // Box.IDs of the Boxes in the pair
	Distance float64
}

type Circuits map[int]int

func createSlice(input string) []Box {

	lines := strings.Split(input, "\n")
	boxes := make([]Box, len(lines))
	for i, line := range lines {
		var a, b, c int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &a, &b, &c)
		if err != nil {
			fmt.Println("error parsing: ", err)
			continue
		}
		co := Coord{a, b, c}
		boxes[i] = Box{
			ID:    i,
			Coord: co,
		}
	}
	return boxes
}

func getDistance(a, b Coord) float64 {
	dx := float64(b[0] - a[0])
	dy := float64(b[1] - a[1])
	dz := float64(b[2] - a[2])
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func getDistances(b []Box) []DistancePair {
	var distances []DistancePair

	for i := range b {
		for j := i + 1; j < len(b); j++ {
			d := getDistance(b[i].Coord, b[j].Coord)
			distances = append(distances, DistancePair{
				I:        b[i].ID,
				J:        b[j].ID,
				Distance: d,
			})
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})
	return distances
}

func createCircuits(d []DistancePair) map[int]int {
	clusters := make(Circuits)
	nextClusterID := 1

	for _, pair := range d {
		id1, id2 := pair.I, pair.J
		c1, ok1 := clusters[id1]
		c2, ok2 := clusters[id2]

		switch {
		case !ok1 && !ok2:
			clusters[id1] = nextClusterID
			clusters[id2] = nextClusterID
			nextClusterID++
		case ok1 && !ok2:
			clusters[id2] = c1
		case !ok1 && ok2:
			clusters[id1] = c2
		case ok1 && ok2 && c1 != c2:
			oldCluster := c2
			for k, v := range clusters {
				if v == oldCluster {
					clusters[k] = c1
				}
			}
		}
	}
	return clusters
}

func countBoxesPerCircuit(circs Circuits) []int {
	counts := make(map[int]int)

	for _, clusterID := range circs {
		counts[clusterID]++
	}
	clusterSizes := make([]int, 0, len(counts))
	for _, size := range counts {
		clusterSizes = append(clusterSizes, size)
	}
	return clusterSizes
}

func P1(input string) int {
	boxes := createSlice(input)

	td := getDistances(boxes)

	circuits := createCircuits(td[:TopN]) // only the top 10

	counts := countBoxesPerCircuit(circuits)

	//sort them, we need just the top 3 for the result
	slices.Sort(counts)
	slices.Reverse(counts)

	result := 1
	for _, c := range counts[:3] {
		result *= c
	}

	return result
}

func findFinalConnection(pairs []DistancePair, boxCount int) DistancePair {
	uf := NewUF(boxCount)

	for _, pair := range pairs {
		merged := uf.Union(pair.I, pair.J)

		if merged && uf.groups == 1 {
			return pair // ← genau dieses Paar ist die gesuchte finale Verbindung
		}
	}

	return DistancePair{} // sollte nie passieren
}

func P2(input string) int {

	boxes := createSlice(input)
	boxCount := len(boxes)

	dists := getDistances(boxes)

	finalPair := findFinalConnection(dists, boxCount)

	var one, two int
	for _, v := range boxes {
		if v.ID == finalPair.I {
			one = v.Coord[0]
		}
		if v.ID == finalPair.J {
			two = v.Coord[0]
		}
	}
	return one * two
}
