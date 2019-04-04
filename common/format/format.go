package format

import "strconv"

func ToInt(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}

func ToInt64(str string) int64 {
	value, _ := strconv.ParseInt(str, 10, 64)
	return value
}

func ToFloat64(str string) float64 {
	value, _ := strconv.ParseFloat(str, 64)
	return value
}
