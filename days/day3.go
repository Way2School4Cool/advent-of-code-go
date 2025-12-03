// Approx Time:

package days

import (
	"bufio"
	"os"
)

type Day3 struct{}

var day3File = "input/day3.txt"

func (d *Day3) Part1() int {
	fileData, _ := os.Open(day3File)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)
	line.Scan()

	return 0
}

func (d *Day3) Part2() int {
	fileData, _ := os.Open(day3File)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)
	line.Scan()

	return 0
}
