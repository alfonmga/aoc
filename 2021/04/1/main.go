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

func CalcWinningBingoBoardScore(bingoInputStr string) (score int64) {
	bingoInputStrLines := strings.Split(bingoInputStr, "\n")
	announcedBingoNumbers := bingoInputStrLines[0]

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
			currentBoardLinesRow := strings.Split(strings.ReplaceAll(currentBoardLines[x], "  ", " "), " ")
			for _, n := range currentBoardLinesRow {
				vInt, err := strconv.Atoi(n)
				handleError(err)
				currentBoardNumbers = append(currentBoardNumbers, vInt)
			}
		}
		bingoBoards = append(bingoBoards, currentBoardNumbers)
	}

	bingoBoardsMarkedNumbers = make(map[int][]int)
	for _, announcedBingoNumber := range announcedBingoNumbers {
		// 1. check each board to see if it contains this number
		// 2. If the board contains it then check winning combinations
		// 3. If contains winning combinations then calculate winning score
	}

	return 0
}

func main() {
	inputBlob, err := ioutil.ReadFile("input.txt")
	handleError(err)
	score := CalcWinningBingoBoardScore(string(inputBlob))
	fmt.Printf("Winning bingo board score is: %v", score)
}
