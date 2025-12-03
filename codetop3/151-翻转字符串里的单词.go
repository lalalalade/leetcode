package codetop3

func reverseWords(s string) string {
	res := ""
	s = " " + s + " "
	l, r := len(s)-1, len(s)-1
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == ' ' {
			l = i
			if r > l+1 {
				res = res + s[l+1:r] + " "
			}
			r = i
		}
	}
	return res[:len(res)-1]
}
