package main

import (
	"fmt"
	"gen/utils/file"
	"gen/utils/number"
	"sort"
	"strconv"
	"strings"
)

func parseLineToNumberPairs(line string) ([]int, error) {
	parts := strings.Fields(line)

	var nums []int
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		nums = append(nums, num)
	}

	return nums, nil
}

func countSimilarityScore(array []int, target int) int {
	count := 0
	for _, num := range array {
		if num == target {
			count++
		}
	}
	return count * target
}

func main() {
	data, _ := file.ReadAndProcessLines("./in.txt", parseLineToNumberPairs)
	columns := [2][]int{{}, {}}

	for _, pair := range data {
		for i, numchik := range pair {
			columns[i] = append(columns[i], numchik)
		}
	}

	for _, column := range columns {
		sort.Ints(column)
	}

	var totalDistance, similarityScore int
	for i := 0; i < len(columns[0]); i++ {
		totalDistance += number.Abs(columns[0][i] - columns[1][i])
		similarityScore += countSimilarityScore(columns[1], columns[0][i])
	}

	file.WriteTextToFile("out.txt", fmt.Sprint(totalDistance))   // Part I
	file.WriteTextToFile("out.txt", fmt.Sprint(similarityScore)) // Part II
}
