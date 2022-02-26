package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func CalcPointLinesOverlap(input string) int {
	inputSplitted := strings.Split(strings.ReplaceAll(input, " ", ""), "\n")

	var lineSegmentsCoords [][][]int
	for _, v := range inputSplitted {
		lineSegmentsCoordsStr := strings.Split(v, "->")
		var currentLineSegmentCoords [][]int
		for _, s := range lineSegmentsCoordsStr {
			currentLineSegmentsCoordsStr := strings.Split(s, ",")
			var currentLineSegmentsCoords []int
			for _, z := range currentLineSegmentsCoordsStr {
				coord, err := strconv.Atoi(z)
				handleError(err)
				currentLineSegmentsCoords = append(currentLineSegmentsCoords, coord)
			}
			currentLineSegmentCoords = append(currentLineSegmentCoords, currentLineSegmentsCoords)
		}
		lineSegmentsCoords = append(lineSegmentsCoords, currentLineSegmentCoords)
	}

	pointsMap := make(map[int]map[int]int)
	for _, lineSegmentCoords := range lineSegmentsCoords {
		lineSegmentCoordX1 := lineSegmentCoords[0][0]
		lineSegmentCoordY1 := lineSegmentCoords[0][1]
		lineSegmentCoordX2 := lineSegmentCoords[1][0]
		lineSegmentCoordY2 := lineSegmentCoords[1][1]

		// only consider horizontal and vertical lines x1=x2 or y1=y2
		if lineSegmentCoordX1 != lineSegmentCoordX2 && lineSegmentCoordY1 != lineSegmentCoordY2 {
			continue
		}

		var currentPointsToMark [][]int

		// add the mark points at the ends of the line segments
		for _, lineSegmentCoord := range lineSegmentCoords {
			x := lineSegmentCoord[0]
			y := lineSegmentCoord[1]
			currentPointsToMark = append(currentPointsToMark, []int{x, y})
		}

		// add the mark points between line segments
		// TODO: implement this!

		// mark points in the map
		for _, currentPointToMark := range currentPointsToMark {
			_, hasYPointCoord := pointsMap[currentPointToMark[1]]
			if hasYPointCoord {
				prevXpointVal, hasXPointCoord := pointsMap[currentPointToMark[1]][currentPointToMark[0]]
				if hasXPointCoord {
					pointsMap[currentPointToMark[1]][currentPointToMark[0]] = prevXpointVal + 1
				} else {
					pointsMap[currentPointToMark[1]][currentPointToMark[0]] = 1
				}
			} else {
				pointsMap[currentPointToMark[1]] = make(map[int]int)
				pointsMap[currentPointToMark[1]][currentPointToMark[0]] = 1
			}
		}
	}

	// count lines point overlap
	overlapCount := 0
	for _, yPoints := range pointsMap {
		for _, xPoints := range yPoints {
			if xPoints >= 2 {
				overlapCount += 1
			}
		}
	}

	return overlapCount
}

func main() {
	inputBlob, err := ioutil.ReadFile("input.txt")
	handleError(err)
	fmt.Printf("At %v points at least two lines overlap.", CalcPointLinesOverlap(string(inputBlob)))
}
