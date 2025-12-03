package main

import "fmt"

//定一个只包含 "()[]{}" 六种字符的字符串。
//规定它们的优先级由外至内为："{}", "[]", "()"，同一级的可以嵌套，并列。
//要求判断给定的字符串是否是合法的括号字串？
//例："()", "{((()())())[()]}()" 是合法的。"())", "([])", "())(()" 都是不合法的。

func main() {
	s1 := "())(()"
	fmt.Println(ff1(s1))
}

func ff1(s string) bool {
	stk := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stk = append(stk, s[i])
		} else {
			if s[i] == ')' {
				if len(stk) == 0 {
					return false
				}
				if stk[len(stk)-1] == '(' {
					stk = stk[:len(stk)-1]
				} else {
					return false
				}
			}
			if s[i] == ']' {
				if len(stk) == 0 {
					return false
				}
				if stk[len(stk)-1] == '[' {
					stk = stk[:len(stk)-1]
				} else {
					return false
				}
			}
			if s[i] == '}' {
				if len(stk) == 0 {
					return false
				}
				if stk[len(stk)-1] == '{' {
					stk = stk[:len(stk)-1]
				} else {
					return false
				}
			}
		}
	}
	return len(stk) == 0
}
