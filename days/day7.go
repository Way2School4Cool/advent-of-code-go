// Approx Time:

package days

import (
	"bufio"
	"os"
	"slices"
)

type Day7 struct{}

var day7File = "input/day7.txt"

func (d *Day7) Part1() int {
	fileData, _ := os.Open(day7File)
	defer fileData.Close()
	var grid [][]rune

	line := bufio.NewScanner(fileData)
	for line.Scan() {
		lineData := line.Text()
		lineRunes := []rune(lineData)

		grid = append(grid, lineRunes)
	}

	currBeams := []int{}
	splits := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if string(grid[y][x]) == "S" {
				currBeams = append(currBeams, x)
			} else {
				if string(grid[y][x]) == "^" && slices.Contains(currBeams, x) {
					posOfMatch := slices.Index(currBeams, x)
					currBeams = append(currBeams[:posOfMatch], currBeams[posOfMatch+1:]...)

					splits++

					if !slices.Contains(currBeams, x-1) {
						currBeams = append(currBeams, x-1)
					}
					if !slices.Contains(currBeams, x+1) {
						currBeams = append(currBeams, x+1)
					}
				}
			}
			if slices.Contains(currBeams, x) {
				//print("|")
			} else {
				//print(string(grid[y][x]))
			}

		}
		//println()
	}

	return splits
}

func (d *Day7) Part2() int {
	fileData, _ := os.Open(day7File)
	defer fileData.Close()
	var grid [][]rune

	line := bufio.NewScanner(fileData)
	for line.Scan() {
		lineData := line.Text()
		lineRunes := []rune(lineData)

		grid = append(grid, lineRunes)
	}

	currBeams := []int{}
	splits := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if string(grid[y][x]) == "S" {
				currBeams = append(currBeams, x)
			} else {
				if string(grid[y][x]) == "^" && slices.Contains(currBeams, x) {
					posOfMatch := slices.Index(currBeams, x)
					currBeams = append(currBeams[:posOfMatch], currBeams[posOfMatch+1:]...)

					splits++

					if !slices.Contains(currBeams, x-1) {
						currBeams = append(currBeams, x-1)
					}
					if !slices.Contains(currBeams, x+1) {
						currBeams = append(currBeams, x+1)
					}
				}
			}
			if slices.Contains(currBeams, x) {
				print("|")
			} else {
				print(string(grid[y][x]))
			}

		}
		println()
	}

	return splits
}
