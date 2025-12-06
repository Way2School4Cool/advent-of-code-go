// Approx Time:

package days

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Day6 struct{}

var day6File = "input/day6.txt"

func (d *Day6) Part1() int {
	fileData, _ := os.Open(day6File)
	defer fileData.Close()
	numOfLines := 0
	positions := []int{}
	array := [][]int{}
	file := []string{}
	symbols := []byte{}
	total := 0

	line := bufio.NewScanner(fileData)
	for line.Scan() {
		numOfLines++
		lineData := line.Text()

		file = append(file, lineData)
	}

	for i := 0; i < len(file[0]); i++ {
		if file[len(file)-1][i] != ' ' {
			positions = append(positions, i)
		}
	}

	for i := 0; i < len(positions)-1; i++ {
		for j := 0; j < numOfLines-1; j++ {
			var innerSlice = []int{}
			stringSpace := strings.Trim(file[j][positions[i]:positions[i+1]], " ")
			numberSpace, _ := strconv.Atoi(stringSpace)

			innerSlice = append(innerSlice, numberSpace)
			array = append(array, innerSlice)
		}
		symbols = append(symbols, file[numOfLines-1][positions[i]])
		//println(string(symbols[i]))
	}

	for i := 0; i < len(array); i++ {
		lineTotal := 0
		for j := 0; j < len(array[i]); j++ {
			if string(symbols[i]) == "+" {
				lineTotal += array[i][j]
			} else if string(symbols[i]) == "*" {
				if lineTotal == 0 {
					lineTotal = array[i][j]
				} else {
					lineTotal = lineTotal * array[i][j]
				}
			}
		}
		total += lineTotal
		println(lineTotal)

	}

	return total
}

func (d *Day6) Part2() int {
	fileData, _ := os.Open(day6File)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)
	line.Scan()

	return 0
}
