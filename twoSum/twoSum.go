package twoSum

import (
	"slices"
)

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		if slices.Contains(nums, target-nums[i]) && slices.Index(nums, target-nums[i]) != i {
			return []int{i, slices.Index(nums, target-nums[i])}
		}
	}

	return nil
}
