package codetop2

import "strconv"

// addStrings 字符串相加
func addStrings(num1 string, num2 string) string {
	add := 0
	res := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		res = strconv.Itoa(result%10) + res
		add = result / 10
	}
	return res
}
