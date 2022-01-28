package tools

import "math"

func GetAuthNum(authArray []int) uint64 {
	var result uint64 = 0
	for _, v := range authArray {
		v = int(math.Pow(2, float64(v)))
		result = result | uint64(v)
	}
	return result
}

func IsAuthorized(grade int, authNum uint64) bool {
	if authNum == 0 {
		return true
	}
	gradeNum := math.Pow(2, float64(grade))
	return uint64(gradeNum)^authNum != 0
}

func HasMaxNum(maxNum int, num int) bool {
	switch maxNum {
	case 0:
		return true
	case 1:
		if num <= 5 {
			return true
		}
	case 2:
		if num <= 10 {
			return true
		}
	case 3:
		if num <= 20 {
			return true
		}
	}
	return true
}
