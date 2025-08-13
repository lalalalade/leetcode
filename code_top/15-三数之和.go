package code_top

import "slices"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	slices.Sort(nums)
	n := len(nums)
	for i, x := range nums[:n-2] {
		// 跳过重复数字
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if x+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				res = append(res, []int{x, nums[j], nums[k]})
				// 跳过重复数字
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}
	return res
}
