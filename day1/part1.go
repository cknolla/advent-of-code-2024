package day1

import (
	"advent-of-code-2024/common"
	"math"
	"slices"
	"strings"
)

func parseFile(filename string) ([]int, []int) {
	var leftList []int
	var rightList []int
	for l := range common.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		locationIds := strings.Fields(line)
		leftList = append(leftList, common.ToInt(locationIds[0]))
		rightList = append(rightList, common.ToInt(locationIds[1]))
	}
	return leftList, rightList
}

func getDifference(leftList, rightList []int) int {
	difference := 0
	for i := range len(leftList) {
		difference += int(math.Abs(float64(leftList[i] - rightList[i])))
	}
	return difference
}

func Part1(filename string) int {
	leftList, rightList := parseFile(filename)
	slices.Sort(leftList)
	slices.Sort(rightList)
	answer := getDifference(leftList, rightList)
	return answer
}
