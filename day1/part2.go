package day1

func getCountMap(locationIds []int) map[int]int {
	countMap := make(map[int]int)
	for _, locationId := range locationIds {
		countMap[locationId]++
	}
	return countMap
}

func getSimilarityScore(locationIds []int, countMap map[int]int) int {
	similarityScore := 0
	for _, locationId := range locationIds {
		similarityScore += locationId * countMap[locationId]
	}
	return similarityScore
}

func Part2(filename string) int {
	leftList, rightList := parseFile(filename)
	countMap := getCountMap(rightList)
	answer := getSimilarityScore(leftList, countMap)
	return answer
}
