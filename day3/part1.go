package day3

import (
	"advent-of-code-2024/common"
	"regexp"
)

func parseFile(filename string) [][]int {
	var pairs [][]int
	for l := range common.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := pattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			pairs = append(pairs, []int{common.ToInt(match[1]), common.ToInt(match[2])})
		}
	}
	return pairs
}

func multiplyPairs(pairs [][]int) int {
	sum := 0
	for _, pair := range pairs {
		sum += pair[0] * pair[1]
	}
	return sum
}

func Part1(filename string) int {
	pairs := parseFile(filename)
	return multiplyPairs(pairs)
}
