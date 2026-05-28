package removeElement

import "sort"

func RemoveElement(nums []int, val int) int {
	sort.Ints(nums)

	length := len(nums)
	k := 0

	for s := BinarySearch(nums, val); s != -1; s = BinarySearch(nums, val) {
		nums = append(nums[:s], nums[s+1:]...)
		k += 1
	}

	return length - k
}

func BinarySearch(nums []int, val int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == val {
			return mid
		}

		if nums[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
