package day4

import "advent-of-code-2024/common"

func huntA(grid common.MatrixChar, startPos common.Pos) int {
	if grid[startPos.Y][startPos.X] != 'A' {
		return 0
	}
	matches := 0
	if grid.IsValidPos(common.Pos{startPos.X - 1, startPos.Y - 1}) && grid.IsValidPos(common.Pos{startPos.X + 1, startPos.Y + 1}) {
		if ((grid[startPos.Y-1][startPos.X-1] == 'M' && grid[startPos.Y+1][startPos.X+1] == 'S') || (grid[startPos.Y-1][startPos.X-1] == 'S' && grid[startPos.Y+1][startPos.X+1] == 'M')) && ((grid[startPos.Y-1][startPos.X+1] == 'M' && grid[startPos.Y+1][startPos.X-1] == 'S') || (grid[startPos.Y-1][startPos.X+1] == 'S' && grid[startPos.Y+1][startPos.X-1] == 'M')) {
			matches++
		}
	}
	return matches
}

func Part2(filename string) int {
	return search(parseFile(filename), huntA)
}
