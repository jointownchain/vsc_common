package utils

import "strconv"

func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Power(base int64, exponent int) int64 {
	if exponent <= 0 {
		return 0
	}
	ans := int64(1)
	for exponent != 0 {
		ans *= base
		exponent--
	}
	return ans
}
