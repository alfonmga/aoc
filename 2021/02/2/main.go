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

func SubmarineMovementCmdHandler(input string) int {
	sCommands := strings.Split(input, "\n")

	submarineXPos := 0
	submarineDepthVal := 0
	submarineAimVal := 0
	for _, v := range sCommands {
		cmd := strings.TrimSpace(strings.SplitAfter(v, " ")[0])
		cmdVal, err := strconv.Atoi(strings.SplitAfter(v, " ")[1])
		handleError(err)

		switch cmd {
		case "forward":
			submarineXPos += cmdVal
			if submarineAimVal != 0 {
				submarineDepthVal += submarineAimVal * cmdVal
			}
		case "up":
			submarineAimVal -= cmdVal
		case "down":
			submarineAimVal += cmdVal
		}
	}

	return submarineXPos * submarineDepthVal
}

func main() {
	inputBlob, err := ioutil.ReadFile("input.txt")
	handleError(err)
	inputStr := string(inputBlob)
	result := SubmarineMovementCmdHandler(inputStr)
	fmt.Printf("Final horizontal position multiplied by final depth [%v]", result)
}
