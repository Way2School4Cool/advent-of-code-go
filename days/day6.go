// Approx Time:

package days

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Day6 struct{}

var day6File = "input/day6example.txt"

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
	numOfLines := 0
	positions := []int{}
	array := [][]int{}
	//vertArray := [][]int{}
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
		pos1 := ""
		pos2 := ""
		pos3 := ""
		pos4 := ""
		thousandsNum, hundredsNum, tensNum, onesNum := 0, 0, 0, 0

		for j := 0; j < len(array[i]); j++ {
			if array[i][j] >= 1000 {
				pos1 += string(strconv.Itoa(array[i][j])[0])
				pos2 += string(strconv.Itoa(array[i][j])[1])
				pos3 += string(strconv.Itoa(array[i][j])[2])
				pos4 += string(strconv.Itoa(array[i][j])[3])
			} else if array[i][j] >= 100 {
				pos1 += "0"
				pos2 += string(strconv.Itoa(array[i][j])[0])
				pos3 += string(strconv.Itoa(array[i][j])[1])
				pos4 += string(strconv.Itoa(array[i][j])[2])
			} else if array[i][j] >= 10 {
				pos1 += "0"
				pos2 += "0"
				pos3 += string(strconv.Itoa(array[i][j])[0])
				pos4 += string(strconv.Itoa(array[i][j])[1])
			} else if array[i][j] >= 1 {
				pos1 += "0"
				pos2 += "0"
				pos3 += "0"
				pos4 += string(strconv.Itoa(array[i][j])[0])
			}

			thousandsNumTemp, _ := strconv.Atoi(pos1)
			hundredsNumTemp, _ := strconv.Atoi(pos2)
			tensNumTemp, _ := strconv.Atoi(pos3)
			onesNumTemp, _ := strconv.Atoi(pos4)

			thousandsNum = thousandsNumTemp
			hundredsNum = hundredsNumTemp
			tensNum = tensNumTemp
			onesNum = onesNumTemp

		}

		println(thousandsNum, string(symbols[i]), hundredsNum, string(symbols[i]), tensNum, string(symbols[i]), onesNum)

		if string(symbols[i]) == "+" {
			total += thousandsNum + hundredsNum + tensNum + onesNum
		} else {
			total += thousandsNum * hundredsNum * tensNum * onesNum
		}

	}

	/*
		for i := 0; i < len(array); i++ {
			lineTotal := 0

			println(array[i][0], string(symbols[i]), array[i][1], string(symbols[i]), array[i][2], string(symbols[i]), array[i][3])

			if string(symbols[i]) == "+" {
				lineTotal = array[i][0] + array[i][1] + array[i][2] + array[i][3]
			} else {
				lineTotal = array[i][0] * array[i][1] * array[i][2] * array[i][3]
			}

			//println(lineTotal)

			total += lineTotal
			//println(lineTotal)

		}
	*/

	return total
}
