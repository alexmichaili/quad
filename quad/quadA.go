package main

import "strings"

func QuadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var builder strings.Builder
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if i == 1 || i == x {
				if j == 1 || j == y {
					builder.WriteRune('o')
				} else {
					builder.WriteRune('-')
				}
			} else {
				if j == 1 || j == y {
					builder.WriteRune('|')
				} else {
					builder.WriteRune(' ')
				}
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}
