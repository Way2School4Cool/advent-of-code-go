// Approx Time:

package days

import (
	"bufio"
	"os"
)

type Day8 struct{}

var day8File = "input/day8.txt"

func (d *Day8) Part1() int {
	fileData, _ := os.Open(day7File)
	defer fileData.Close()
	var grid [][]rune

	line := bufio.NewScanner(fileData)
	for line.Scan() {
		lineData := line.Text()
		lineRunes := []rune(lineData)

		grid = append(grid, lineRunes)

	}

	return 0
}

func (d *Day8) Part2() int {
	fileData, _ := os.Open(day7File)
	defer fileData.Close()
	var grid [][]rune

	line := bufio.NewScanner(fileData)
	for line.Scan() {
		lineData := line.Text()
		lineRunes := []rune(lineData)

		grid = append(grid, lineRunes)

	}

	return 0
}
