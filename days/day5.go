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

	freshItems := []int{}
	freshCount := 0

	for line.Scan() {
		lineData := line.Text()

		if lineData == "" {
			postSpace = true
			continue
		}

		if !postSpace {
			beginItems, _ := strconv.Atoi(strings.Split(lineData, "-")[0])
			endItems, _ := strconv.Atoi(strings.Split(lineData, "-")[1])

			for beginItems <= endItems {
				freshItems = append(freshItems, beginItems)
				beginItems++
			}

		} else {
			currentItem, _ := strconv.Atoi(lineData)
			for i := 0; i < len(freshItems); i++ {
				if freshItems[i] == currentItem {
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
	line.Scan()

	return 0
}
