package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func hasWinningBingoCombination(drawnNumbers map[int]bool, candidateNumbers []int) bool {
	var hasWinCombination = false
	for idx, v := range candidateNumbers {
		_, found := drawnNumbers[v]
		if !found {
			break
		}
		if idx == 4 {
			hasWinCombination = true
		}
	}

	return hasWinCombination
}

func calcWinningBoardScore(drawnNumbers map[int]bool, lastDrawnNumber int, boardNumbers []int) int {
	var unmarkedBoardNumbers []int
	for _, v := range boardNumbers {
		_, isNumberMarked := drawnNumbers[v]
		if !isNumberMarked {
			unmarkedBoardNumbers = append(unmarkedBoardNumbers, v)
		}
	}
	var sumUnmarkedNumbers = 0
	for _, v := range unmarkedBoardNumbers {
		sumUnmarkedNumbers += v
	}

	return sumUnmarkedNumbers * lastDrawnNumber
}

func PlayBingo(bingoInputStr string) (score int) {
	bingoInputStrLines := strings.Split(bingoInputStr, "\n")

	toBeAnnouncedBingoNumbersInput := bingoInputStrLines[0]
	toBeAnnouncedBingoNumbersStr := strings.Split(toBeAnnouncedBingoNumbersInput, ",")
	var toBeAnnouncedBingoNumbers []int
	for _, nStr := range toBeAnnouncedBingoNumbersStr {
		nInt, err := strconv.Atoi(nStr)
		handleError(err)
		toBeAnnouncedBingoNumbers = append(toBeAnnouncedBingoNumbers, nInt)
	}

	bingoBoardsInput := bingoInputStrLines[2:]
	var bingoBoards [][]int
	for i := 0; i < len(bingoBoardsInput); i++ {
		currentBingoBoardInputLine := bingoBoardsInput[i]
		if currentBingoBoardInputLine == "" {
			return
		}
		currentBoardLines := bingoBoardsInput[i : i+5]
		i = i + 5
		var currentBoardNumbers []int
		for x := 0; x < len(currentBoardLines); x++ {
			re := regexp.MustCompile(`\d+`)
			currentBoardLinesRow := re.FindAllString(currentBoardLines[x], -1)
			for _, n := range currentBoardLinesRow {
				vInt, err := strconv.Atoi(n)
				handleError(err)
				currentBoardNumbers = append(currentBoardNumbers, vInt)
			}
		}
		bingoBoards = append(bingoBoards, currentBoardNumbers)
	}

	var drawnBingoNumbers = make(map[int]bool)
	wonBoards := make(map[int]bool, len(bingoBoards))
	var lastBoardWinningScore int
	for _, announcedBingoNumber := range toBeAnnouncedBingoNumbers {
		drawnBingoNumbers[announcedBingoNumber] = true

		if len(drawnBingoNumbers) < 5 {
			continue // There cannot be winners if less than 5 numbers have been drawn!
		}

		hasWinningBoard := false
		for bingoBoardIdx, bingoBoardNumbers := range bingoBoards {
			hasWinningCombination := false

			// check rows wins
			for i := 0; i < 5; i++ {
				rowHeadPos := i * 5
				candidateRowNumbers := bingoBoardNumbers[rowHeadPos : rowHeadPos+5]

				if hasWinningBingoCombination(drawnBingoNumbers, candidateRowNumbers) {
					hasWinningCombination = true
					break
				}
			}
			if !hasWinningCombination {
				// check columns wins
				for i := 0; i < 5; i++ {
					columnHeadPos := i
					var candidateColumnNumbers []int
					for x := 0; x < 5; x++ {
						targetIdx := columnHeadPos
						if x > 0 {
							targetIdx += x * 5
						}
						candidateColumnNumbers = append(candidateColumnNumbers, bingoBoardNumbers[targetIdx])
					}

					if hasWinningBingoCombination(drawnBingoNumbers, candidateColumnNumbers) {
						hasWinningCombination = true
						break
					}
				}
			}

			if hasWinningCombination {
				_, thisBoardAlreadyWon := wonBoards[bingoBoardIdx]
				if !thisBoardAlreadyWon {
					wonBoards[bingoBoardIdx] = true
					numBoardsWon := 0
					for _, v := range wonBoards {
						if v {
							numBoardsWon += 1
						}
					}
					if numBoardsWon == len(bingoBoards) {
						hasWinningBoard = true
						lastBoardWinningScore = calcWinningBoardScore(drawnBingoNumbers, announcedBingoNumber, bingoBoardNumbers)
						break
					}
				}

			}
		}
		if hasWinningBoard {
			break
		}
	}

	return lastBoardWinningScore
}

func main() {
	inputBlob, err := ioutil.ReadFile("input.txt")
	handleError(err)
	score := PlayBingo(string(inputBlob))
	fmt.Printf("The last bingo board to win score was: %v", score)
}
