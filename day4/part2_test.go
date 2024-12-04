package day4

import (
	"advent-of-code-2024/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHuntA(t *testing.T) {
	grid := parseFile("sample.txt")
	matches := huntA(grid, common.Pos{2, 1})
	assert.Equal(t, 1, matches)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 2041, Part2("input.txt"))
}
