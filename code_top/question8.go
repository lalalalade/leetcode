package code_top

func myAtoi(s string) int {
	res, sign, i := 0, 1, 0
	// 去除前导0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	// 判定符号
	if i < len(s) {
		if s[i] == '+' || s[i] == '-' {
			if s[i] == '-' {
				sign = -1
			}
			i++
		}
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		res = res*10 + int(s[i]-'0')
		if sign == 1 && res > 1<<31-1 {
			return 1<<31 - 1
		}
		if sign == -1 && res > 1<<31 {
			return -1 << 31
		}
		i++
	}
	return res * sign
}
