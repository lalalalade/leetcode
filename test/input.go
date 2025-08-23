package test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 设置大缓冲区处理长行（关键优化）
	const maxCapacity = 1024 * 1024 * 4 // 4MB
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	// 读取单行输入
	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	// 读取并分割单词
	readWords := func() []string {
		return strings.Fields(readLine())
	}
	// 读取整数数组
	readInts := func() []int64 {
		words := readWords()
		nums := make([]int64, len(words))
		for i, s := range words {
			nums[i], _ = strconv.ParseInt(s, 10, 64)
		}
		return nums
	}

	n := 0
	res := int64(0)
	fmt.Scan(&n)
	nums := readInts()
	i, j := int64(0), int64(n-1)
	for i < j {
		res = max(res, min(nums[i], nums[j])*(j-i))
		if nums[i] <= nums[j] {
			i++
		} else {
			j--
		}
	}
}

func max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}

func min(i, j int64) int64 {
	if i < j {
		return i
	}
	return j
}
