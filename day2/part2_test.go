package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSafeWithDampener(t *testing.T) {
	minDifference := 1
	maxDifference := 3
	assert.True(t, isSafeWithDampener([]int{7, 6, 4, 2, 1}, minDifference, maxDifference))
	assert.False(t, isSafeWithDampener([]int{1, 2, 7, 8, 9}, minDifference, maxDifference))
	assert.False(t, isSafeWithDampener([]int{9, 7, 6, 2, 1}, minDifference, maxDifference))
	assert.True(t, isSafeWithDampener([]int{1, 3, 2, 4, 5}, minDifference, maxDifference))
	assert.True(t, isSafeWithDampener([]int{8, 6, 4, 4, 1}, minDifference, maxDifference))
	assert.True(t, isSafeWithDampener([]int{7, 6, 4, 2, 1}, minDifference, maxDifference))
	assert.True(t, isSafeWithDampener([]int{16, 13, 10, 9, 8, 6, 6}, minDifference, maxDifference))
	assert.True(t, isSafeWithDampener([]int{75, 72, 69, 67, 68}, minDifference, maxDifference))
	assert.False(t, isSafeWithDampener([]int{15, 12, 12, 9, 9}, minDifference, maxDifference))
	assert.True(t, isSafeWithDampener([]int{21, 19, 17, 17, 16}, minDifference, maxDifference))
	assert.True(t, isSafeWithDampener([]int{44, 41, 39, 37, 35, 32, 27}, minDifference, maxDifference))
	assert.True(t, isSafeWithDampener([]int{30, 41, 39, 37, 35, 32}, minDifference, maxDifference))
}

func TestPart2(t *testing.T) {
	answer := Part2("input.txt")
	assert.Equal(t, 301, answer)
}
