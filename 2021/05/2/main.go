package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}
type Line struct {
	slope int
	yint  int
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateLine(a, b Point) Line {
	slope := (b.y - a.y) / (b.x - a.x)
	yint := a.y - slope*a.x
	return Line{slope, yint}
}
func EvalX(l Line, x int) int {
	return l.slope*x + l.yint
}
func Intersection(l1, l2 Line) (Point, bool) {
	if l1.slope == l2.slope {
		return Point{}, false
	}
	x := (l2.yint - l1.yint) / (l1.slope - l2.slope)
	y := EvalX(l1, x)
	return Point{x, y}, true
}

func CalcPointLinesOverlap(input string) int {
	inputSplitted := strings.Split(strings.ReplaceAll(input, " ", ""), "\n")

	var linesSegments [][]Point
	for _, v := range inputSplitted {
		lineSegmentsCoordsStr := strings.Split(v, "->")
		var currentLineSegmentCoords []Point
		for _, s := range lineSegmentsCoordsStr {
			currentLineSegmentsCoordsStr := strings.Split(s, ",")
			var currentLineSegmentsCoords []int
			for _, z := range currentLineSegmentsCoordsStr {
				coord, err := strconv.Atoi(z)
				handleError(err)
				currentLineSegmentsCoords = append(currentLineSegmentsCoords, coord)
			}
			currentLineSegmentCoords = append(currentLineSegmentCoords, Point{int(currentLineSegmentsCoords[0]), int(currentLineSegmentsCoords[1])})
		}
		linesSegments = append(linesSegments, currentLineSegmentCoords)
	}

	intersectionPointsMap := make(map[int]map[int]int)

	for lineSegmentsIdx, lineSegments := range linesSegments {
		lineSegment1 := Point{lineSegments[0].x, lineSegments[0].y}
		lineSegment2 := Point{lineSegments[1].x, lineSegments[1].y}

		deltaX := lineSegment2.x - lineSegment1.x
		deltaY := lineSegment2.y - lineSegment1.y

		lineRad := math.Atan2(float64(deltaY), float64(deltaX))
		lineDeg := lineRad * (180 / math.Pi)

		if lineSegmentCoordX1 != lineSegmentCoordX2 && lineSegmentCoordY1 != lineSegmentCoordY2 {
			continue
		}

		line := CreateLine(Point{lineSegments[0].x, lineSegments[0].y}, Point{lineSegments[0].x, lineSegments[1].y})

		for targetLineSegmentsIdx, targetLineSegments := range linesSegments {
			if lineSegmentsIdx == targetLineSegmentsIdx {
				continue
			}
			targetLine := CreateLine(Point{targetLineSegments[0].x, targetLineSegments[0].y}, Point{targetLineSegments[0].x, targetLineSegments[1].y})

			intersectedPoint, intersected := Intersection(line, targetLine)
			if intersected {
				_, hasYPointCoord := intersectionPointsMap[intersectedPoint.y]
				if hasYPointCoord {
					prevXpointVal, hasXPointCoord := intersectionPointsMap[intersectedPoint.y][intersectedPoint.x]
					if hasXPointCoord {
						intersectionPointsMap[intersectedPoint.y][intersectedPoint.x] = prevXpointVal + 1
					} else {
						intersectionPointsMap[intersectedPoint.y][intersectedPoint.x] = 1
					}
				} else {
					intersectionPointsMap[intersectedPoint.y] = make(map[int]int)
					intersectionPointsMap[intersectedPoint.y][intersectedPoint.x] = 1
				}
			}

		}

	}

	intersectionsCounter := 0
	for _, yPoints := range intersectionPointsMap {
		for _, xPoints := range yPoints {
			intersectionsCounter += xPoints
		}
	}

	return intersectionsCounter
}

func main() {
	inputBlob, err := ioutil.ReadFile("input.txt")
	handleError(err)
	fmt.Printf("At %v points at least two lines overlap.", CalcPointLinesOverlap(string(inputBlob)))
}
