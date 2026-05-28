package grayCode

import "math"

func GrayCode(n int) []int {
	ans := []int{}

	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		ans = append(ans, i^(i>>1))
	}

	return ans
}
