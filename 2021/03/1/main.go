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

func binaryToDecimal(bN string) int64 {
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

func SubmarineDiagnosticReport(input string) int64 {
	diagnosticReportBlob := strings.Split(input, "\n")

	gammaRate := calculateGammaRateByBlob(diagnosticReportBlob)
	epsilonRate := calculateEpsilonRateByBlob(diagnosticReportBlob)

	powerConsumption := binaryToDecimal(gammaRate) * binaryToDecimal(epsilonRate)

	return powerConsumption
}

func main() {
	inputBlob, err := ioutil.ReadFile("input.txt")
	handleError(err)
	inputStr := string(inputBlob)
	result := SubmarineDiagnosticReport(inputStr)
	fmt.Printf("Current submarine power cosumption [%v].", result)
}
