package array

func SubWithoutIndex(arr []int, excludeIndex int) []int {
	result := make([]int, 0, len(arr)-1)
	result = append(result, arr[:excludeIndex]...)
	result = append(result, arr[excludeIndex+1:]...)
	return result
}
