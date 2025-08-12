package double_pointer

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := range n - 3 {
		x := nums[i]
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if x+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		// 枚举第二个数
		for j := i + 1; j < n-2; j++ {
			y := nums[j]
			if j > i+1 && y == nums[j-1] {
				continue
			}
			if x+y+nums[j+1]+nums[j+2] > target {
				break
			}
			if x+y+nums[n-2]+nums[n-1] < target {
				continue
			}
			// 双指针
			u, v := j+1, n-1
			for u < v {
				s := x + y + nums[u] + nums[v]
				if s > target {
					v--
				} else if s < target {
					u++
				} else {
					res = append(res, []int{x, y, nums[u], nums[v]})
					for u++; u < v && nums[u] == nums[u-1]; u++ {
					}
					for v--; v > u && nums[v] == nums[v+1]; v-- {
					}
				}
			}
		}
	}
	return res
}
