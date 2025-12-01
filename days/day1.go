// Approx Time: 2h 15m

package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Day1 struct{}

var inputFile = "input/day1.txt"

func (d *Day1) Part1() int {
	zeros := 0
	position := 50

	fileData, _ := os.Open(inputFile)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)
	for line.Scan() {
		move := line.Text()
		moveAmount, _ := strconv.Atoi(move[1:])

		if move[0] == 'R' {
			position += moveAmount
		} else if move[0] == 'L' {
			position -= moveAmount
		}

		for {
			if position >= 100 {
				position -= 100
			} else if position < 0 {
				position += 100
			} else {
				break
			}
		}

		if position == 0 {
			zeros++
		}
	}

	return zeros
}

func (d *Day1) Part2() int {
	zeros := 0
	position := 50

	fileData, _ := os.Open(inputFile)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)
	for line.Scan() {
		move := line.Text()
		moveAmount, _ := strconv.Atoi(move[1:])

		if move[0] == 'R' {
			for i := 0; i < moveAmount; i++ {
				position++
				if position >= 100 {
					zeros++
					position -= 100
				}
			}

			println(fmt.Sprintf("Move: %s        Position: %d        Zeros: %d", move, position, zeros))

			continue

		} else if move[0] == 'L' {
			startingPosition := position

			for i := 0; i < moveAmount; i++ {
				position -= 1
				if position < 0 {
					zeros++
					position += 100
				}
			}

			if startingPosition == 0 {
				zeros--
			}

			if position == 0 {
				zeros++
			}

			//println(fmt.Sprintf("Move: %s        Position: %d        Zeros: %d", move, position, zeros))

			continue
		}

	}

	return zeros
}
