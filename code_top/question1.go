package code_top

func twoSum(nums []int, target int) []int {
	idx := make(map[int]int)
	for j, x := range nums {
		if i, ok := idx[target-x]; ok {
			return []int{i, j}
		}
		idx[x] = j
	}
	return nil
}
