package day4

import (
	"advent-of-code-2024/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	grid := parseFile("sample.txt")
	assert.Equal(t, 10, grid.LenX())
	assert.Equal(t, 10, grid.LenY())
	assert.Equal(t, []uint8{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'}, grid[0])
	assert.Equal(t, []uint8{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'}, grid[9])
}

func TestHuntX(t *testing.T) {
	grid := parseFile("sample.txt")
	assert.Equal(t, uint8('X'), grid[4][6])
	matches := huntX(grid, common.Pos{4, 6})
	assert.Equal(t, 2, matches)
}

func TestSearch(t *testing.T) {
	grid := parseFile("sample.txt")
	assert.Equal(t, 18, search(grid, huntX))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 2575, Part1("input.txt"))
}
