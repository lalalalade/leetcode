package double_pointer

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	// i < j < k
	n := len(nums)
	for i := 0; i < n-2; i++ {
		x := nums[i]
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if x+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j += 1
				for j < k && nums[j] == nums[j-1] {
					j += 1
				}
				k -= 1
				for k > j && nums[k] == nums[k+1] {
					k -= 1
				}
			}
		}
	}
	return res
}
