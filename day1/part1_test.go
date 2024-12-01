package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	leftList, rightList := parseFile("sample.txt")
	assert.Equal(t, leftList, []int{3, 4, 2, 1, 3, 3})
	assert.Equal(t, rightList, []int{4, 3, 5, 3, 9, 3})
}

func TestGetDifferences(t *testing.T) {
	difference := getDifference(
		[]int{1, 2, 3, 3, 3, 4},
		[]int{3, 3, 3, 4, 5, 9},
	)
	assert.Equal(t, 11, difference)
}

func TestPart1(t *testing.T) {
	answer := Part1("input.txt")
	assert.Equal(t, 2264607, answer)
}
