package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

		deltaX := lineSegmentCoordX2 - lineSegmentCoordX1
		deltaY := lineSegmentCoordY2 - lineSegmentCoordY1

		rad := math.Atan2(float64(deltaY), float64(deltaX))
		deg := rad * (180 / math.Pi)

		isVerticalOrHorizontal := math.Abs(deg) == 0 || math.Abs(deg) == 90 || math.Abs(deg) == 135 || math.Abs(deg) == 180
		isDiagonalCoords := deg == 45

		if !isDiagonalCoords && !isVerticalOrHorizontal {
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
		if isDiagonalCoords {
			isAscX := lineSegmentCoordX1 < lineSegmentCoordX2
			isAscY := lineSegmentCoordY1 < lineSegmentCoordY2
			var targetXCoord int
			var startXCoord int
			var startYCoord int
			var stopAt int
			if isAscX {
				targetXCoord = lineSegmentCoordX2
				startXCoord = lineSegmentCoordX1
				startYCoord = lineSegmentCoordY1
				stopAt = targetXCoord
			} else {
				targetXCoord = lineSegmentCoordX2
				startXCoord = lineSegmentCoordX1
				startYCoord = lineSegmentCoordY1
				stopAt = targetXCoord + 1
			}

			pointsAway := targetXCoord - startXCoord
			if pointsAway != 1 {
				for {
					if isAscX {
						startXCoord++
						if isAscY {
							startYCoord++
						} else {
							startYCoord--
						}
					} else {
						startXCoord--
						startYCoord++
					}

					currentPointsToMark = append(currentPointsToMark, []int{startXCoord, startYCoord})

					if startXCoord == stopAt {
						break
					}
				}
			}
		} else {
			if deltaX != 0 {
				var targetCoord int
				var startCoord int
				if lineSegmentCoordX1 > lineSegmentCoordX2 {
					targetCoord = lineSegmentCoordX1
					startCoord = lineSegmentCoordX2
				} else {
					targetCoord = lineSegmentCoordX2
					startCoord = lineSegmentCoordX1
				}

				if targetCoord-startCoord > 1 {
					for {
						startCoord++
						currentPointsToMark = append(currentPointsToMark, []int{startCoord, lineSegmentCoordY1})
						if startCoord == targetCoord-1 {
							break
						}
					}
				}
			} else if deltaY != 0 {
				var targetCoord int
				var startCoord int
				if lineSegmentCoordY1 > lineSegmentCoordY2 {
					targetCoord = lineSegmentCoordY1
					startCoord = lineSegmentCoordY2
				} else {
					targetCoord = lineSegmentCoordY2
					startCoord = lineSegmentCoordY1
				}

				if targetCoord-startCoord > 1 {
					for {
						startCoord++
						currentPointsToMark = append(currentPointsToMark, []int{lineSegmentCoordX1, startCoord})
						if startCoord == targetCoord-1 {
							break
						}
					}
				}
			}
		}

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
			if xPoints > 1 {
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
