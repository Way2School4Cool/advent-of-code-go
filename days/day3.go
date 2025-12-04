// Approx Time:

package days

import (
	"bufio"
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

func (d *Day3) Part2() int {
	fileData, _ := os.Open(day3File)
	defer fileData.Close()
	line := bufio.NewScanner(fileData)

	runningTotal := 0

	for line.Scan() {
		lineData := line.Text()

		currentLargest := -1
		newList := []int{}

		for {
			largest := -1
			positionOfLargest := -1

			// Find largest number
			for i := currentLargest + 1; i < len(lineData)-12; i++ {
				currentNumber, _ := strconv.Atoi(string(lineData[i]))
				if currentNumber > largest {
					largest, _ = strconv.Atoi(string(lineData[i]))
					positionOfLargest = i
					currentLargest = i
				}
			}

			print(largest)
			newList = append(newList, largest)

			// remake the line removing evething before the largest number found inclusive
			lineData = lineData[positionOfLargest+1:]

			if len(lineData) == 12 {
				//combine the new list and the end of the old list to make the final number of 12 digits
				lineData = lineData[len(newList):]
				//newlist

				var numberString string
				for _, digit := range newList {
					numberString += strconv.Itoa(digit)
				}

				print(newList, lineData)
				// Join the string from newList with the remaining lineData
				finalNumberString := numberString + lineData

				// Convert the final 12-digit string to a number and add to the total
				finalNumber, _ := strconv.ParseInt(finalNumberString, 10, 64)
				runningTotal += int(finalNumber)

				break
			}
		}
	}

	return runningTotal

}

//2222222322411112221272122222252212213222213415221122222152422222222222142321222222222122561232221
//2222222322411112221272122222252212213222213415221122222152422222222222142321222222222122561232221
