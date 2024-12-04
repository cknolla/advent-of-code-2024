package day4

import (
	"advent-of-code-2024/common"
)

type huntFunc func(common.MatrixChar, common.Pos) int

func parseFile(filename string) common.MatrixChar {
	var lines []string
	for l := range common.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return common.BuildMatrixChar(lines)
}

func search(grid common.MatrixChar, fn huntFunc) int {
	totalMatches := 0
	for y := 0; y < grid.LenY(); y++ {
		for x := 0; x < grid.LenX(); x++ {
			totalMatches += fn(grid, common.Pos{x, y})
		}
	}
	return totalMatches
}

func huntX(grid common.MatrixChar, startPos common.Pos) int {
	if grid[startPos.Y][startPos.X] != 'X' {
		return 0
	}
	matches := 0
	if grid.IsValidPos(common.Pos{startPos.X, startPos.Y - 3}) {
		if grid[startPos.Y-1][startPos.X] == 'M' && grid[startPos.Y-2][startPos.X] == 'A' && grid[startPos.Y-3][startPos.X] == 'S' {
			matches++
		}
	}
	if grid.IsValidPos(common.Pos{startPos.X - 3, startPos.Y - 3}) {
		if grid[startPos.Y-1][startPos.X-1] == 'M' && grid[startPos.Y-2][startPos.X-2] == 'A' && grid[startPos.Y-3][startPos.X-3] == 'S' {
			matches++
		}
	}
	if grid.IsValidPos(common.Pos{startPos.X - 3, startPos.Y}) {
		if grid[startPos.Y][startPos.X-1] == 'M' && grid[startPos.Y][startPos.X-2] == 'A' && grid[startPos.Y][startPos.X-3] == 'S' {
			matches++
		}
	}
	if grid.IsValidPos(common.Pos{startPos.X - 3, startPos.Y + 3}) {
		if grid[startPos.Y+1][startPos.X-1] == 'M' && grid[startPos.Y+2][startPos.X-2] == 'A' && grid[startPos.Y+3][startPos.X-3] == 'S' {
			matches++
		}
	}
	if grid.IsValidPos(common.Pos{startPos.X, startPos.Y + 3}) {
		if grid[startPos.Y+1][startPos.X] == 'M' && grid[startPos.Y+2][startPos.X] == 'A' && grid[startPos.Y+3][startPos.X] == 'S' {
			matches++
		}
	}
	if grid.IsValidPos(common.Pos{startPos.X + 3, startPos.Y + 3}) {
		if grid[startPos.Y+1][startPos.X+1] == 'M' && grid[startPos.Y+2][startPos.X+2] == 'A' && grid[startPos.Y+3][startPos.X+3] == 'S' {
			matches++
		}
	}
	if grid.IsValidPos(common.Pos{startPos.X + 3, startPos.Y}) {
		if grid[startPos.Y][startPos.X+1] == 'M' && grid[startPos.Y][startPos.X+2] == 'A' && grid[startPos.Y][startPos.X+3] == 'S' {
			matches++
		}
	}
	if grid.IsValidPos(common.Pos{startPos.X + 3, startPos.Y - 3}) {
		if grid[startPos.Y-1][startPos.X+1] == 'M' && grid[startPos.Y-2][startPos.X+2] == 'A' && grid[startPos.Y-3][startPos.X+3] == 'S' {
			matches++
		}
	}
	return matches
}

func Part1(filename string) int {
	return search(parseFile(filename), huntX)
}
