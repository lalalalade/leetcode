package stack_queue

// isValid 有效的括号
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	mp := map[rune]rune{')': '(', ']': '[', '}': '{'}
	st := make([]rune, 0)
	for _, c := range s {
		if _, exists := mp[c]; !exists {
			st = append(st, c)
		} else {
			if len(st) == 0 || st[len(st)-1] != mp[c] {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}
