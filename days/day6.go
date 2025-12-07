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

			//println(i)
			positions = append(positions, i)
		}
	}

	for i := 0; i < len(positions)-1; i++ {
		var innerSlice = []int{}
		for j := 0; j < numOfLines; j++ {
			stringSpace := strings.Trim(file[j][positions[i]:positions[i+1]], " ")
			numberSpace, _ := strconv.Atoi(stringSpace)
			innerSlice = append(innerSlice, numberSpace)
		}
		array = append(array, innerSlice)
		symbols = append(symbols, file[numOfLines-1][positions[i]])
		//println(string(symbols[i]))
	}

	//last one is getting lost somewhere
	val1, _ := strconv.Atoi(strings.Trim(file[0][positions[len(positions)-1]:], " "))
	val2, _ := strconv.Atoi(strings.Trim(file[1][positions[len(positions)-1]:], " "))
	val3, _ := strconv.Atoi(strings.Trim(file[2][positions[len(positions)-1]:], " "))
	val4, _ := strconv.Atoi(strings.Trim(file[3][positions[len(positions)-1]:], " "))

	symbols = append(symbols, file[numOfLines-1][positions[len(positions)-1]])

	array = append(array, []int{val1, val2, val3, val4})

	//println(len(array), len(symbols))

	for i := 0; i < len(array); i++ {
		lineTotal := 0

		//println(array[i][0], string(symbols[i]), array[i][1], string(symbols[i]), array[i][2], string(symbols[i]), array[i][3])

		if string(symbols[i]) == "+" {
			lineTotal = array[i][0] + array[i][1] + array[i][2] + array[i][3]
		} else {
			lineTotal = array[i][0] * array[i][1] * array[i][2] * array[i][3]
		}

		/*
			for j := 0; j < len(symbols); j++ {
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
		*/

		//println(lineTotal)

		total += lineTotal
		//println(lineTotal)

	}

	return total
}

func (d *Day6) Part2() int {
	fileData, _ := os.Open(day6File)
	defer fileData.Close()
	//vertArray := [][]int{}
	var grid [][]rune
	total := 0

	line := bufio.NewScanner(fileData)
	for line.Scan() {
		lineData := line.Text()
		lineRunes := []rune(lineData)

		grid = append(grid, lineRunes)

	}

	numbers := []int{}
	symbol := ""

	for x := 0; x < len(grid[0]); x++ {
		wipNumbers := ""
		isAllSpaces := true
		for y := 0; y < len(grid); y++ {
			char := grid[y][x]

			if char != ' ' {
				isAllSpaces = false
				if char == '+' {
					symbol = "+"
				} else if char == '*' {
					symbol = "*"
				} else {
					wipNumbers += string(char)
				}
			}
		}

		// After iterating a full column, if we built a number, add it.
		if wipNumbers != "" {
			number, _ := strconv.Atoi(wipNumbers)
			numbers = append(numbers, number)
		}

		// If the column was all spaces or it's the last column, perform calculation.
		if isAllSpaces || x == len(grid[0])-1 {
			if len(numbers) > 0 {
				wipTotal := 0
				if symbol == "*" {
					wipTotal = 1
				}

				for _, num := range numbers {
					if symbol == "+" {
						wipTotal += num
					} else if symbol == "*" {
						wipTotal *= num
					}
				}
				total += wipTotal
			}
			// Reset for the next group.
			numbers = []int{}
			symbol = ""
		}
	}

	return total
}
