// Approx Time:

package days

import (
	"bufio"
	"math/big"
	"os"
	"strconv"
)

type Day3 struct{}

var day3File = "input/day3.txt"

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
		pos1 := -1
		pos2 := -1
		pos3 := -1

		for i := 0; i < len(lineData)-2; i++ {
			currentNumber, _ := strconv.Atoi(string(lineData[i]))
			nextNumber, _ := strconv.Atoi(string(lineData[i+1]))

			if currentNumber < nextNumber || currentNumber == 1 {
				print("current number:", currentNumber, " next number:", nextNumber, "\n")

				if pos1 == -1 {
					pos1 = i
					//println("Setting Pos 1 to ", pos1)
				} else if pos2 == -1 {
					pos2 = i
					//println("Setting Pos 2 to ", pos2)
				} else if pos3 == -1 {
					pos3 = i
					//println("Setting Pos 3 to ", pos3)
					break
				}
			}
		}

		if pos3 == -1 {
			pos3 = len(lineData) - 1
			print("AAAAAAAAAAAAAAAAAA")
		}
		if pos2 == -1 {
			pos2 = len(lineData) - 2
		}
		if pos1 == -1 {
			pos1 = len(lineData) - 3
		}

		strippedLine := lineData[:pos1] + lineData[pos1+1:pos2] + lineData[pos2+1:pos3] + lineData[pos3+1:]

		strippedLineBigInt := new(big.Int)
		strippedLineBigInt.SetString(strippedLine, 10)

		runningTotal.Add(runningTotal, strippedLineBigInt)

		println(lineData)
		println(strippedLine)
		//println(runningTotal.String(), "\n")
	}

	return runningTotal.String()
}

//2222222322411112221272122222252212213222213415221122222152422222222222142321222222222122561232221
//2222222322411112221272122222252212213222213415221122222152422222222222142321222222222122561232221
