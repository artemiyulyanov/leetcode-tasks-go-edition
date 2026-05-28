package searchInsertPosition

func SearchInsertPosition(nums []int, target int) int {
	left, right := 1, len(nums)-1
	ans := 0

	if target <= nums[0] {
		return 0
	}

	for left <= right {
		mid := (left + right) / 2
		ans = mid

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if nums[ans] == target {
		return ans
	}

	return left
}
