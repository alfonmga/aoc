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

func DepthMeasurementCounter(input string) int {
	sMeasurements := strings.Split(input, "\n")

	numTimesDepthMeasurementIncreased := 0
	for i := 0; i < len(sMeasurements); i++ {
		if i == 0 {
			continue
		}

		currentMeasurement, err := strconv.Atoi(sMeasurements[i])
		handleError(err)
		prevMeasurement, err := strconv.Atoi(sMeasurements[i-1])
		handleError(err)

		if currentMeasurement > prevMeasurement {
			numTimesDepthMeasurementIncreased++
		}
	}

	return numTimesDepthMeasurementIncreased
}

func main() {
	inputBlob, err := ioutil.ReadFile("input.txt")
	handleError(err)
	inputStr := string(inputBlob)
	result := DepthMeasurementCounter(inputStr)
	fmt.Printf("Depth measurement has increased [%v] times.", result)
}
