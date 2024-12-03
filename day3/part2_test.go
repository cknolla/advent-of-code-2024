package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFileWithEnable(t *testing.T) {
	pairs := parseFileWithEnable("sample2.txt")
	assert.Equal(t, [][]int{
		{2, 4},
		{8, 5},
	}, pairs)
}

func TestPart2(t *testing.T) {
	answer := Part2("input.txt")
	assert.Equal(t, 78683433, answer)
}
