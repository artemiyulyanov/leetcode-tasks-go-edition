package sqrt

func sqrt(x int) int {
	ans := 0

	left, right := 1, x

	for left <= right {
		mid := (left + right) / 2

		square := mid * mid

		if square == x {
			return mid
		}

		ans = mid

		if square > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
