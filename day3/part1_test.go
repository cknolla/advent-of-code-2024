package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	pairs := parseFile("sample.txt")
	assert.Equal(t, [][]int{
		{2, 4},
		{5, 5},
		{11, 8},
		{8, 5},
	}, pairs)
}

func TestMultiplyPairs(t *testing.T) {
	assert.Equal(t, 161, multiplyPairs([][]int{
		{2, 4},
		{5, 5},
		{11, 8},
		{8, 5},
	}))
}

func TestPart1(t *testing.T) {
	answer := Part1("input.txt")
	assert.Equal(t, 170068701, answer)
}
