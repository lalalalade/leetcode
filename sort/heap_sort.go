package sort

func sortArray2(nums []int) []int {
	var heapify func([]int, int, int)
	heapify = func(nums []int, root, end int) {
		for {
			// 左孩子节点索引
			child := root*2 + 1
			if child > end {
				return
			}
			// 选择左右孩子节点中的较大者
			if child < end && nums[child] < nums[child+1] {
				child++
			}
			if nums[root] > nums[child] {
				return
			}
			nums[root], nums[child] = nums[child], nums[root]
			root = child
		}
	}
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		heapify(nums, i, end)
	}
	// 依次将堆顶元素与最后一个元素交换，然后调整堆
	for i := end; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i-1)
	}
	return nums
}
