package tools

import "strconv"

func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}
func Int2String(intArr []int) []string {
	res := make([]string, len(intArr))
	for index, val := range intArr {
		res[index] = strconv.Itoa(val)
	}
	return res
}
