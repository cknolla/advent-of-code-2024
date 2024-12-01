package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCountMap(t *testing.T) {
	countMap := getCountMap(
		[]int{3, 3, 3, 4, 5, 9},
	)
	assert.Equal(t, map[int]int{
		3: 3,
		4: 1,
		5: 1,
		9: 1,
	}, countMap)
}

func TestGetSimilarityScore(t *testing.T) {
	score := getSimilarityScore(
		[]int{3, 4, 2, 1, 3, 3},
		map[int]int{
			3: 3,
			4: 1,
			5: 1,
			9: 1,
		})
	assert.Equal(t, 31, score)
}

func TestPart2(t *testing.T) {
	answer := Part2("input.txt")
	assert.Equal(t, 19457120, answer)
}
