package day2

import (
	"advent-of-code-2024/common"
	"strings"
)

func parseFile(filename string) [][]int {
	var reports [][]int
	for l := range common.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		levels := strings.Fields(line)
		levelInts := make([]int, len(levels))
		for i, level := range levels {
			levelInts[i] = common.ToInt(level)
		}
		reports = append(reports, levelInts)
	}
	return reports
}

func isSafe(report []int, minDifference int, maxDifference int) bool {
	if len(report) < 2 {
		return false
	}
	if report[1] == report[0] {
		return false
	}
	if report[1] < report[0] {
		tmpMinDifference := minDifference
		minDifference = maxDifference * -1
		maxDifference = tmpMinDifference * -1
	}
	for i := 1; i < len(report); i++ {
		difference := report[i] - report[i-1]
		if difference < minDifference || difference > maxDifference {
			return false
		}
	}
	return true
}

func Part1(filename string) int {
	safeReports := 0
	reports := parseFile(filename)
	for _, report := range reports {
		if isSafe(report, 1, 3) {
			safeReports++
		}
	}
	return safeReports
}
