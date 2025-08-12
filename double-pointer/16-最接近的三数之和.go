package double_pointer

import (
	"math"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	res := 0
	slices.Sort(nums)
	n := len(nums)
	minDiff := math.MaxInt
	for i, x := range nums[:n-2] {

		if i > 0 && x == nums[i-1] {
			continue
		}
		// 优化1
		s := x + nums[i+1] + nums[i+2]
		if s > target {
			if s-target < minDiff {
				res = s
			}
			break
		}

		// 优化2
		s = x + nums[n-2] + nums[n-1]
		if s < target {
			if target-s < minDiff {
				minDiff = target - s
				res = s
			}
			continue
		}
		j, k := i+1, n-1
		for j < k {
			s = x + nums[j] + nums[k]
			if s == target {
				return target
			}
			if s > target {
				if s-target < minDiff {
					minDiff = s - target
					res = s
				}
				k--
			} else {
				if target-s < minDiff {
					minDiff = target - s
					res = s
				}
				j++
			}
		}
	}
	return res
}
