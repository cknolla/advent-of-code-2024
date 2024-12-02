package day2

func isSafeWithDampener(report []int, minDifference int, maxDifference int) bool {
	if isSafe(report, minDifference, maxDifference) {
		return true
	}
	for i := 0; i < len(report); i++ {
		var shortReport []int
		shortReport = append(shortReport, report[:i]...)
		shortReport = append(shortReport, report[i+1:]...)
		if isSafe(shortReport, minDifference, maxDifference) {
			return true
		}
	}
	return false
}

func Part2(filename string) int {
	safeReports := 0
	reports := parseFile(filename)
	for _, report := range reports {
		if isSafeWithDampener(report, 1, 3) {
			safeReports++
		}
	}
	return safeReports
}
