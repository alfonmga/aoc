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

func BinaryToDecimal(bN string) int64 {
	n, err := strconv.ParseInt(bN, 2, 64)
	handleError(err)
	return n
}

func calculateGammaRateByBlob(blob []string) string {
	rowsN := len(blob)
	columnsN := len(strings.Split(blob[0], ""))

	gammaRate := ""
	for i := 0; i < columnsN; i++ {

		dict := make(map[int]int)
		for x := 0; x < rowsN; x++ {
			bit, err := strconv.Atoi(strings.Split(blob[x], "")[i])
			handleError(err)
			dict[bit] = dict[bit] + 1
		}

		if dict[0] > dict[1] {
			gammaRate += "0"
		} else {
			gammaRate += "1"
		}
	}

	return gammaRate
}
func calculateEpsilonRateByBlob(blob []string) string {
	rowsN := len(blob)
	columnsN := len(strings.Split(blob[0], ""))

	epsilonRate := ""
	for i := 0; i < columnsN; i++ {

		dict := make(map[int]int)
		for x := 0; x < rowsN; x++ {
			bit, err := strconv.Atoi(strings.Split(blob[x], "")[i])
			handleError(err)
			dict[bit] = dict[bit] + 1
		}

		if dict[0] > dict[1] {
			epsilonRate += "1"
		} else {
			epsilonRate += "0"
		}
	}
	return epsilonRate
}

func CalcOxigenGeneratorRating(blob []string) string {
	var currentBitPos = 0
	var currentBytes = blob
	for 1 < len(currentBytes) {

		currentBitPosBytesGroupedByBitDic := make(map[int][]string)
		for i := 0; i < len(currentBytes); i++ {
			bit, err := strconv.Atoi(strings.Split(currentBytes[i], "")[currentBitPos])
			handleError(err)
			currentBitPosBytesGroupedByBitDic[bit] = append(currentBitPosBytesGroupedByBitDic[bit], currentBytes[i])
		}

		if len(currentBitPosBytesGroupedByBitDic[0]) > len(currentBitPosBytesGroupedByBitDic[1]) {
			currentBytes = currentBitPosBytesGroupedByBitDic[0]
		} else if len(currentBitPosBytesGroupedByBitDic[0]) < len(currentBitPosBytesGroupedByBitDic[1]) {
			currentBytes = currentBitPosBytesGroupedByBitDic[1]
		} else {
			currentBytes = currentBitPosBytesGroupedByBitDic[1]
		}

		currentBitPos++
	}

	return currentBytes[0]
}
func CalcC02ScrubberRating(blob []string) string {
	var currentBitPos = 0
	var currentBytes = blob
	for 1 < len(currentBytes) {

		currentBitPosBytesGroupedByBitDic := make(map[int][]string)
		for i := 0; i < len(currentBytes); i++ {
			bit, err := strconv.Atoi(strings.Split(currentBytes[i], "")[currentBitPos])
			handleError(err)
			currentBitPosBytesGroupedByBitDic[bit] = append(currentBitPosBytesGroupedByBitDic[bit], currentBytes[i])
		}

		if len(currentBitPosBytesGroupedByBitDic[0]) < len(currentBitPosBytesGroupedByBitDic[1]) {
			currentBytes = currentBitPosBytesGroupedByBitDic[0]
		} else if len(currentBitPosBytesGroupedByBitDic[0]) > len(currentBitPosBytesGroupedByBitDic[1]) {
			currentBytes = currentBitPosBytesGroupedByBitDic[1]
		} else {
			currentBytes = currentBitPosBytesGroupedByBitDic[0]
		}

		currentBitPos++
	}

	return currentBytes[0]
}

func SubmarineDiagnosticReport(input string) []int64 {
	diagnosticReportBlob := strings.Split(input, "\n")

	gammaRate := calculateGammaRateByBlob(diagnosticReportBlob)
	epsilonRate := calculateEpsilonRateByBlob(diagnosticReportBlob)
	powerConsumption := BinaryToDecimal(gammaRate) * BinaryToDecimal(epsilonRate)

	oxigenGeneratorRating := CalcOxigenGeneratorRating(diagnosticReportBlob)
	c02ScrubberRating := CalcC02ScrubberRating(diagnosticReportBlob)
	lifeSupportRating := BinaryToDecimal(oxigenGeneratorRating) * BinaryToDecimal(c02ScrubberRating)

	return []int64{powerConsumption, lifeSupportRating}
}

func main() {
	inputBlob, err := ioutil.ReadFile("input.txt")
	handleError(err)
	inputStr := string(inputBlob)
	result := SubmarineDiagnosticReport(inputStr)
	fmt.Printf("Submarine power cosumption [%v]\nSubmarine life support rating [%v]", result[0], result[1])
}
