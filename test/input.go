package main

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
	T := readInts()[0]
	for T > 0 {
		l := readInts()
		n, x, y := l[0], l[1], l[2]
		g := make([][]int, n+1)
		for i := 1; i <= int(n); i++ {
			g[i] = make([]int, 0)
		}
		for i := 0; i < int(n)-1; i++ {
			scanner.Scan()
			var u, v int
			fmt.Sscanf(scanner.Text(), "%d %d", &u, &v)
			g[u] = append(g[u], v)
			g[v] = append(g[v], u)
		}
		if n == 0 {
			fmt.Println(0)
			continue
		}
		dp := make([]int64, n+1)
		visited := make([]bool, n+1)
		var dfs func(int)
		dfs = func(u int) {
			visited[u] = true
			sum := int64(0)
			for _, v := range g[u] {
				if !visited[v] {
					dfs(v)
					sum += dp[v]
				}
			}
			dp[u] = max(y, x+sum)
		}
		dfs(1)
		fmt.Println(dp[1])
		T--
	}
}
