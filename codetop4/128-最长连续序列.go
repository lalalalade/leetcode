package codetop4

func longestConsecutive(nums []int) int {
	res := 0
	has := map[int]bool{}
	for _, num := range nums {
		has[num] = true
	}
	for x := range has {
		if has[x-1] {
			continue
		}
		y := x + 1
		for has[y] {
			y++
		}
		res = max(res, y-x)
	}
	return res
}
