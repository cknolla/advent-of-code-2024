package day3

import (
	"advent-of-code-2024/common"
	"regexp"
	"strings"
)

func parseFileWithEnable(filename string) [][]int {
	var builder strings.Builder
	for l := range common.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		builder.WriteString(line)
	}
	var pairs [][]int
	allButLast := regexp.MustCompile(`don't\(\).*?do\(\)`)
	finalDont := regexp.MustCompile(`don't\(\).*`)
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	partiallyEnabledLine := allButLast.ReplaceAllString(builder.String(), "")
	enabledLine := finalDont.ReplaceAllString(partiallyEnabledLine, "")
	matches := pattern.FindAllStringSubmatch(enabledLine, -1)
	for _, match := range matches {
		pairs = append(pairs, []int{common.ToInt(match[1]), common.ToInt(match[2])})
	}
	return pairs
}

func Part2(filename string) int {
	pairs := parseFileWithEnable(filename)
	return multiplyPairs(pairs)
}
