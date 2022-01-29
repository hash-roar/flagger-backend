package tools

import "math"

func SwitchArrayToNum(intArray []int) uint64 {
	var result uint64 = 0
	for _, v := range intArray {
		v = int(math.Pow(2, float64(v)))
		result = result | uint64(v)
	}
	return result
}

func SwitchNumToArray(num uint64) []int {
	var result []int
	for i := 0; i < 10; i++ {
		powNum := math.Pow(2, float64(i))
		if num&uint64(powNum) != 0 {
			result = append(result, i)
		}
	}
	return result
}
