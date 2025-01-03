package main

import (
	"gen/utils/array"
	"gen/utils/file"
	"gen/utils/number"
	"strconv"
	"strings"
)

func parseLineToNumbersArray(line string) ([]int, error) {
	parts := strings.Fields(line)

	var nums []int
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		nums = append(nums, num)
	}

	return nums, nil
}

func main() {
	data, _ := file.ReadAndProcessLines("in.txt", parseLineToNumbersArray)

	var safeReports = 0

	for _, report := range data {
		if number.ToBool(isReportSave(report)) {
			safeReports += 1
		} else {
			safeReports += isReportSaveWithoutSingleEl(report)
		}
	}

	file.WriteTextToFile("out.txt", strconv.Itoa(safeReports))

}

func isReportSave(report []int) int {

	isIncreasing := report[1] > report[0]
	isDecreasing := report[1] < report[0]

	for i := 1; i < len(report); i++ {
		diff := number.Abs(report[i] - report[i-1])

		if diff < 1 || diff > 3 || report[i] == report[i-1] ||
			(isIncreasing && report[i] < report[i-1]) ||
			(isDecreasing && report[i] > report[i-1]) {
			return 0
		}
	}

	return 1
}

func isReportSaveWithoutSingleEl(arr []int) int {
	for i := 0; i < len(arr); i++ {
		modifiedArray := array.SubWithoutIndex(arr, i)
		if number.ToBool(isReportSave(modifiedArray)) {
			return 1
		}
	}
	return 0
}
