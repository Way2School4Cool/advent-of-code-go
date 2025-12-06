// Approx Time:

package days

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Day5 struct{}

var day5File = "input/day5.txt"

func (d *Day5) Part1() int {
	fileData, _ := os.Open(day5File)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)
	postSpace := false

	//freshItems := []int{}
	freshCount := 0

	beginRanges := []int{}
	endRanges := []int{}

	for line.Scan() {
		lineData := line.Text()

		//print(lineData)

		if lineData == "" {
			postSpace = true
			continue
		}

		if !postSpace {
			beginItems, _ := strconv.Atoi(strings.Split(lineData, "-")[0])
			endItems, _ := strconv.Atoi(strings.Split(lineData, "-")[1])

			beginRanges = append(beginRanges, beginItems)
			endRanges = append(endRanges, endItems)

		} else {
			currentItem, _ := strconv.Atoi(lineData)
			for i := 0; i < len(beginRanges); i++ {
				if currentItem >= beginRanges[i] && currentItem <= endRanges[i] {
					//println(currentItem)

					freshCount++
					break
				}
			}
		}
	}

	//for i := 0; i < len(freshItems); i++ {
	//	println(freshItems[i])
	//}

	return freshCount
}

func (d *Day5) Part2() int {
	fileData, _ := os.Open(day5File)
	defer fileData.Close()

	line := bufio.NewScanner(fileData)
	postSpace := false

	freshCount := 0

	ranges := [][]int{}
	row := 0

	for line.Scan() {
		lineData := line.Text()

		//print(lineData)

		if lineData == "" {
			postSpace = true
			continue
		}

		if !postSpace {

			beginItems, _ := strconv.Atoi(strings.Split(lineData, "-")[0])
			endItems, _ := strconv.Atoi(strings.Split(lineData, "-")[1])

			innerSlice := []int{}
			innerSlice = append(innerSlice, beginItems, endItems)

			ranges = append(ranges, innerSlice)
			row++

		} else {
			break
		}
	}

	for rotations := 0; rotations < 20; rotations++ {
		newRanges := [][]int{}

		for i := 0; i < len(ranges); i++ {
			matchFound := false

			for j := 0; j < len(ranges); j++ {
				if i != j && ranges[i][0] >= ranges[j][0] && ranges[i][0] <= ranges[j][1] {
					matchFound = true
					if ranges[i][1] <= ranges[j][1] {
						//println(ranges[i][0], ranges[i][1], "within", ranges[j][0], ranges[j][1])

					} else {
						println(ranges[i][0], ranges[i][1], "partially within", ranges[j][0], ranges[j][1])
						innerSlice := []int{}
						innerSlice = append(innerSlice, ranges[i][0], ranges[j][1])
						newRanges = append(newRanges, innerSlice)

					}
				}
			}
			if !matchFound {
				println(ranges[i][1], "-", ranges[i][0], "=", ranges[i][1]-ranges[i][0])
				innerSlice := []int{}
				innerSlice = append(innerSlice, ranges[i][0], ranges[i][1])
				newRanges = append(newRanges, innerSlice)

			}
		}

		ranges = newRanges
	}

	for i := 0; i < len(ranges); i++ {
		freshCount += (ranges[i][1] - ranges[i][0])
	}

	freshCount += len(ranges)

	return freshCount
}
