// Approx Time: 1hr

package days

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Day2 struct{}

var day2File = "input/day2.txt"

func (d *Day2) Part1() int {
	runningTotal := 0

	fileData, _ := os.Open(day2File)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)
	line.Scan()

	ranges := strings.Split(line.Text(), ",")

	for _, valRange := range ranges {
		values := strings.Split(valRange, "-")
		starting, _ := strconv.Atoi(values[0])
		ending, _ := strconv.Atoi(values[1])

		for i := starting; i <= ending; i++ {
			numStr := strconv.Itoa(i)

			if len(numStr)%2 == 0 {
				firstHalf := numStr[:len(numStr)/2]
				secondHalf := numStr[len(numStr)/2:]

				if firstHalf == secondHalf {
					//println(i)

					runningTotal += i
				}
			}
		}
	}

	return runningTotal
}

func (d *Day2) Part2() int {
	runningTotal := 0

	fileData, _ := os.Open(day2File)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)
	line.Scan()

	ranges := strings.Split(line.Text(), ",")

	for _, valRange := range ranges {
		values := strings.Split(valRange, "-")
		starting, _ := strconv.Atoi(values[0])
		ending, _ := strconv.Atoi(values[1])

		println("Checking range:", starting, "-", ending)

		for i := starting; i <= ending; i++ {
			numStr := strconv.Itoa(i)

			matchFound := false

			if !matchFound {

				for devisor := 1; devisor <= len(numStr)/2; devisor++ {

					if len(numStr)%devisor == 0 && !matchFound {
						numOfSections := len(numStr) / devisor

						sections := make([]string, numOfSections)

						for sectionIndex := 0; sectionIndex < numOfSections; sectionIndex++ {
							sections[sectionIndex] = numStr[sectionIndex*devisor : (sectionIndex+1)*devisor]
						}

						allMatch := true
						matchingNum := sections[0]

						for _, section := range sections {
							if section != matchingNum {
								allMatch = false
								break
							}
						}

						if allMatch {
							runningTotal += i

							matchFound = true
						}
					}
				}
			}
		}
	}

	return runningTotal
}
