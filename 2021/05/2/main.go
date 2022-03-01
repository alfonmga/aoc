package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Point struct{ X, Y int }
type Segment struct{ P1, P2 Point }

func minMax(x1, x2 int) (int, int) {
	if x1 <= x2 {
		return x1, x2
	}
	return x2, x1
}

func CalcPointLinesOverlap(input string) int {
	visited := map[Point]int{}

	data := strings.Split(input, "\n")
	var segments = make([]Segment, len(data))
	for i, l := range data {
		fmt.Sscanf(l, "%d,%d -> %d,%d", &segments[i].P1.X, &segments[i].P1.Y, &segments[i].P2.X, &segments[i].P2.Y)
	}

	for _, s := range segments {

		switch {

		// vertical
		case s.P1.X == s.P2.X:
			minY, maxY := minMax(s.P1.Y, s.P2.Y)
			for y := minY; y <= maxY; y++ {
				visited[Point{s.P1.X, y}]++
			}

		// horizontal
		case s.P1.Y == s.P2.Y:
			minX, maxY := minMax(s.P1.X, s.P2.X)
			for x := minX; x <= maxY; x++ {
				visited[Point{x, s.P1.Y}]++
			}

		// diagonal
		default:
			delta := abs(s.P1.X - s.P2.X)

			x, y := s.P1.X, s.P1.Y
			mx, my := sign(s.P2.X-s.P1.X), sign(s.P2.Y-s.P1.Y)

			for i := 0; i <= delta; i++ {
				visited[Point{x + i*mx, y + i*my}]++
			}
		}
	}

	count := 0
	for _, v := range visited {
		if v >= 2 {
			count++
		}
	}
	return count
}

func main() {
	inputBlob, err := ioutil.ReadFile("input.txt")
	handleError(err)
	fmt.Printf("At %v points at least two lines overlap.", CalcPointLinesOverlap(string(inputBlob)))
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}
func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
func sign(i int) int {
	if i >= 0 {
		return +1
	}
	return -1
}
