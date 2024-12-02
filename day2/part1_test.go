package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	reports := parseFile("sample.txt")
	assert.Equal(t, reports, [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	})
}

func TestIsSafe(t *testing.T) {
	minDifference := 1
	maxDifference := 3
	assert.True(t, isSafe([]int{7, 6, 4, 2, 1}, minDifference, maxDifference))
	assert.False(t, isSafe([]int{1, 2, 7, 8, 9}, minDifference, maxDifference))
	assert.False(t, isSafe([]int{9, 7, 6, 2, 1}, minDifference, maxDifference))
	assert.False(t, isSafe([]int{1, 3, 2, 4, 5}, minDifference, maxDifference))
	assert.False(t, isSafe([]int{8, 6, 4, 4, 1}, minDifference, maxDifference))
	assert.True(t, isSafe([]int{7, 6, 4, 2, 1}, minDifference, maxDifference))
}

func TestPart1(t *testing.T) {
	answer := Part1("input.txt")
	assert.Equal(t, 230, answer)
}
