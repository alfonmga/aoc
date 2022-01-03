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

		if len(sMeasurements[i:]) < 3 || len(sMeasurements[i+1:]) < 3 {
			continue
		}

		aWindow := sMeasurements[i : i+3]
		bWindow := sMeasurements[i+1 : i+1+3]

		aWindowSum := 0
		for _, v := range aWindow {
			vSum, err := strconv.Atoi(v)
			handleError(err)
			aWindowSum += vSum
		}
		bWindowSum := 0
		for _, v := range bWindow {
			vSum, err := strconv.Atoi(v)
			handleError(err)
			bWindowSum += vSum
		}

		if bWindowSum > aWindowSum {
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
	fmt.Printf("Depth measurement (three-measurement windows) has increased [%v] times.", result)
}
