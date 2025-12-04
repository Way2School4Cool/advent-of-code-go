// Approx Time: 45 min

package days

import (
	"bufio"
	"os"
)

type Day4 struct{}

var day4File = "input/day4.txt"

func (d *Day4) Part1() int {
	fileData, _ := os.Open(day4File)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)

	vert := 0
	girdSize := 140

	// Make a 2d array
	var positions [][]string = make([][]string, girdSize)
	for i := range positions {
		positions[i] = make([]string, girdSize)
	}

	for line.Scan() {
		for i, char := range line.Text() {
			positions[vert][i] = string(char)
			//print(positions[vert][i])
		}
		//println()
		vert++
	}

	total := 0

	for r, row := range positions {
		for c, _ := range row {
			adjacent := 0

			var up = r - 1
			var down = r + 1
			var left = c - 1
			var right = c + 1

			if positions[r][c] == "@" {

				//print(positions[r][c])

				// Check the top row first
				if r == 0 {

					// If it is in the first or last column, we are safe
					if c == 0 || c == len(positions[r])-1 {

						// less than 4 adjacent
						//updatedPositions[r][c] = "x"

						// check the
					} else {
						if positions[r][left] == "@" {
							adjacent++
						}
						if positions[r][right] == "@" {
							adjacent++
						}
						if positions[down][left] == "@" {
							adjacent++
						}
						if positions[down][c] == "@" {
							adjacent++
						}
						if positions[down][right] == "@" {
							adjacent++
						}
					}
					// do the same for the last row
				} else if r == len(positions)-1 {
					if c == 0 || c == len(positions[r])-1 {
						// less than 4 adjacent
						//updatedPositions[r][c] = "x"

					} else {
						if positions[r][left] == "@" {
							adjacent++
						}
						if positions[r][right] == "@" {
							adjacent++
						}
						if positions[up][left] == "@" {
							adjacent++
						}
						if positions[up][c] == "@" {
							adjacent++
						}
						if positions[up][right] == "@" {
							adjacent++
						}
					}
					// if we arent at the top or bottom
				} else {
					if c == 0 {
						if positions[up][c] == "@" {
							adjacent++
						}
						if positions[down][c] == "@" {
							adjacent++
						}
						if positions[up][right] == "@" {
							adjacent++
						}
						if positions[r][right] == "@" {
							adjacent++
						}
						if positions[down][right] == "@" {
							adjacent++
						}
					} else if c == len(positions[r])-1 {
						if positions[up][c] == "@" {
							adjacent++
						}
						if positions[down][c] == "@" {
							adjacent++
						}
						if positions[up][left] == "@" {
							adjacent++
						}
						if positions[r][left] == "@" {
							adjacent++
						}
						if positions[down][left] == "@" {
							adjacent++
						}
					} else {
						if positions[up][left] == "@" {
							adjacent++
						}
						if positions[up][c] == "@" {
							adjacent++
						}
						if positions[up][right] == "@" {
							adjacent++
						}
						if positions[r][left] == "@" {
							adjacent++
						}
						if positions[r][right] == "@" {
							adjacent++
						}
						if positions[down][left] == "@" {
							adjacent++
						}
						if positions[down][c] == "@" {
							adjacent++
						}
						if positions[down][right] == "@" {
							adjacent++
						}
					}
				}

				if adjacent < 4 {
					//positions[r][c] = "x"
					total++
				}

			}

			print(adjacent)
		}
		println()
	}

	return total
}

func (d *Day4) Part2() int {
	fileData, _ := os.Open(day4File)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)

	vert := 0
	girdSize := 140

	// Make a 2d array
	var positions [][]string = make([][]string, girdSize)
	for i := range positions {
		positions[i] = make([]string, girdSize)
	}

	for line.Scan() {
		for i, char := range line.Text() {
			positions[vert][i] = string(char)
			//print(positions[vert][i])
		}
		//println()
		vert++
	}

	var newPositions = positions

	total := 0

	max := -1

	for {
		positions = newPositions

		for r, row := range positions {
			for c, _ := range row {
				adjacent := 0

				var up = r - 1
				var down = r + 1
				var left = c - 1
				var right = c + 1

				if positions[r][c] == "@" {

					//print(positions[r][c])

					// Check the top row first
					if r == 0 {

						// If it is in the first or last column, we are safe
						if c == 0 || c == len(positions[r])-1 {

							// less than 4 adjacent
							//updatedPositions[r][c] = "x"

							// check the
						} else {
							if positions[r][left] == "@" {
								adjacent++
							}
							if positions[r][right] == "@" {
								adjacent++
							}
							if positions[down][left] == "@" {
								adjacent++
							}
							if positions[down][c] == "@" {
								adjacent++
							}
							if positions[down][right] == "@" {
								adjacent++
							}
						}
						// do the same for the last row
					} else if r == len(positions)-1 {
						if c == 0 || c == len(positions[r])-1 {
							// less than 4 adjacent
							//updatedPositions[r][c] = "x"

						} else {
							if positions[r][left] == "@" {
								adjacent++
							}
							if positions[r][right] == "@" {
								adjacent++
							}
							if positions[up][left] == "@" {
								adjacent++
							}
							if positions[up][c] == "@" {
								adjacent++
							}
							if positions[up][right] == "@" {
								adjacent++
							}
						}
						// if we arent at the top or bottom
					} else {
						if c == 0 {
							if positions[up][c] == "@" {
								adjacent++
							}
							if positions[down][c] == "@" {
								adjacent++
							}
							if positions[up][right] == "@" {
								adjacent++
							}
							if positions[r][right] == "@" {
								adjacent++
							}
							if positions[down][right] == "@" {
								adjacent++
							}
						} else if c == len(positions[r])-1 {
							if positions[up][c] == "@" {
								adjacent++
							}
							if positions[down][c] == "@" {
								adjacent++
							}
							if positions[up][left] == "@" {
								adjacent++
							}
							if positions[r][left] == "@" {
								adjacent++
							}
							if positions[down][left] == "@" {
								adjacent++
							}
						} else {
							if positions[up][left] == "@" {
								adjacent++
							}
							if positions[up][c] == "@" {
								adjacent++
							}
							if positions[up][right] == "@" {
								adjacent++
							}
							if positions[r][left] == "@" {
								adjacent++
							}
							if positions[r][right] == "@" {
								adjacent++
							}
							if positions[down][left] == "@" {
								adjacent++
							}
							if positions[down][c] == "@" {
								adjacent++
							}
							if positions[down][right] == "@" {
								adjacent++
							}
						}
					}

					if adjacent < 4 {
						newPositions[r][c] = "x"
						total++
					}
				}
			}
		}

		if max == total {
			break
		}

		max = total

		println(total)
	}

	return total
}
