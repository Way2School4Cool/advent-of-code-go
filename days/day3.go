// Approx Time:

package days

import (
	"bufio"
	"math/big"
	"os"
	"strconv"
)

type Day3 struct{}

var day3File = "input/day3example.txt"

func (d *Day3) Part1() int {
	fileData, _ := os.Open(day3File)
	defer fileData.Close()
	line := bufio.NewScanner(fileData)

	runningTotal := 0

	for line.Scan() {
		lineData := line.Text()

		//highestFailure := 100
		tens := -1
		ones := -1
		tensFail := -1

		for {
			for i := 0; i < len(lineData); i++ {

				// Look through the line backwards
				currentNumber, _ := strconv.Atoi(string(lineData[i]))

				if currentNumber > tens && (tensFail == -1 || currentNumber < tensFail) {
					tens = currentNumber
					ones = -1

				} else if currentNumber > ones {
					ones = currentNumber
				}

			}

			if ones == -1 {
				tensFail = tens
				tens = -1
			} else {

				//println(lineData)
				//print("largest: ", (tens*10)+ones, "    ")

				runningTotal += tens*10 + ones

				//println(runningTotal)

				break
			}
		}
	}

	return runningTotal
}

func (d *Day3) Part2() string {
	fileData, _ := os.Open(day3File)
	defer fileData.Close()
	line := bufio.NewScanner(fileData)

	runningTotal := big.NewInt(0)

	for line.Scan() {
		lineData := line.Text()

		hundreds := 100
		tens := 100
		ones := 100

		hPos := -1
		tPos := -1
		oPos := -1

		hundredsFail := 100
		tensFail := 100

		for {
			for i := 0; i < len(lineData); i++ {

				// Look through the line backwards
				currentNumber, _ := strconv.Atoi(string(lineData[i]))

				if currentNumber < hundreds && (hundredsFail == 100 || currentNumber > hundredsFail) {

					hPos = i
					tPos = -1
					oPos = -1

					hundreds = currentNumber
					tens = 100
					ones = 100

				} else if currentNumber < tens && (tensFail == 100 || currentNumber > tensFail) {
					tPos = i
					oPos = -1

					tens = currentNumber
					ones = 100

				} else if currentNumber < ones {
					oPos = i
					ones = currentNumber
				}

			}

			if tens == 100 {
				hundredsFail = hundreds
				hundreds = 100
				hPos = -1
			} else if ones == 100 {
				tensFail = tens
				tens = 100
				hundreds = 100
				hPos = -1
				tPos = -1
			} else {

				println("Smallest number positions:", hPos, tPos, oPos)

				println(lineData)

				// remove all numbers at the positions
				outputLine := lineData[:hPos] + lineData[hPos+1:tPos] + lineData[tPos+1:oPos] + lineData[oPos+1:]

				println(outputLine)

				unintTotal, _ := new(big.Int).SetString(outputLine, 10)
				runningTotal.Add(runningTotal, unintTotal)

				println(runningTotal.String(), "\n")

				break
			}
		}
	}

	return runningTotal.String()
}
