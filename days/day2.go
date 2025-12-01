package days

import (
	"bufio"
	"os"
)

type Day2 struct{}

var day2File = "input/day2.txt"

func (d *Day2) Part1() int {
	fileData, _ := os.Open(day2File)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)

	return 0
}

func (d *Day2) Part2() int {
	fileData, _ := os.Open(day2File)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)

	return 0
}
