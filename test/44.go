package main

import "unicode"

func GetEndPoint(order string) []int {
	// write code here
	res := []int{0, 0}
	x, y := 0, 0
	i := 0
	for i < len(order) {
		count := 1
		numStr := ""

		for i < len(order) && unicode.IsDigit(rune(order[i])) {
			numStr += string(order[i])
			i++
		}

		if numStr != "" {
			count := 0
			for _, ch := range numStr {
				count = count*10 + int(ch-'0')
			}
		}

		if i < len(order) {
			dir := order[i]
			i++

			switch dir {
			case 'W', 'w':
				y += count
			case 'S', 's':
				y -= count
			case 'D', 'd':
				x += count
			case 'A', 'a':
				x -= count
			}
		}
	}
	res[0], res[1] = x, y
	return res
}
