package main

import "fmt"

// 给你一个整数数组 nums ，返回全部为 0 的 子数组 数目。
//
// 子数组 是一个数组中一段连续非空元素组成的序列。
//
// 示例 1：
//
// 输入：nums = [1,3,0,0,2,0,0,4]
// 输出：6
// 解释：
// 子数组 [0] 出现了 4 次。
// 子数组 [0,0] 出现了 2 次。
// 不存在长度大于 2 的全 0 子数组，所以我们返回 6 。
// 示例 2：
//
// 输入：nums = [0,0,0,2,0,0]
// 输出：9
// 解释：
// 子数组 [0] 出现了 5 次。
// 子数组 [0,0] 出现了 3 次。
// 子数组 [0,0,0] 出现了 1 次。
// 不存在长度大于 3 的全 0 子数组，所以我们返回 9 。
// 示例 3：
//
// 输入：nums = [2,10,2019]
// 输出：0
// 解释：没有全 0 子数组，所以我们返回 0 。
func main() {
	nums := []int{2, 10, 2019}
	fmt.Println(f1(nums))
}

func f1(nums []int) int {
	res := 0
	n := len(nums)
	// dp[i][j]代表nums[i]到nums[j]是否是全0数组
	dp := make([][]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			dp[i][i] = true
			res++
		}
	}
	// 枚举长度
	for length := 2; length <= n; length++ {
		for start := 0; start+length-1 < n; start++ {
			end := start + length - 1
			if length == 2 {
				dp[start][end] = (nums[start] == 0) && (nums[end] == 0)
			} else {
				// 长度》= 3
				if nums[start] != 0 || nums[end] != 0 {
					dp[start][end] = false
				} else {
					dp[start][end] = dp[start+1][end-1]
				}
			}
			if dp[start][end] {
				res++
			}
		}
	}

	return res
}
