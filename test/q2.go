package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s1 := "1.10.9"
	s2 := "2.6"
	fmt.Println(ff2(s1, s2))
}

// 实现编写一个函数，返回比较2个版本号的大小的结果
//bool CompareVersion(string version1, string version2)

func ff2(version1, version2 string) int {
	a := strings.Split(version1, ".")
	b := strings.Split(version2, ".")
	m, n := len(a), len(b)

	for i := 0; i < max(m, n); i++ {
		num1 := 0
		if i < m {
			num1, _ = strconv.Atoi(a[i])
		}
		num2 := 0
		if i < n {
			num2, _ = strconv.Atoi(b[i])
		}
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}
	return 0
}
