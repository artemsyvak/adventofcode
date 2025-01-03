package number

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ToBool(num int) bool {
	return num != 0
}
