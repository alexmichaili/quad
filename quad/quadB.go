package main

import "strings"

func QuadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var builder strings.Builder
	for i := 1; i <= y; i++ { // i rows, j columns
		for j := 1; j <= x; j++ {
			if i == 1 {
				if j == 1 {
					builder.WriteRune('/')
				} else if j == x {
					builder.WriteRune('\\')
				} else {
					builder.WriteRune('*')
				}
			} else if i == y {
				if j == 1 {
					builder.WriteRune('\\')
				} else if j == x {
					builder.WriteRune('/')
				} else {
					builder.WriteRune('*')
				}
			} else {
				if j == 1 || j == x {
					builder.WriteRune('*')
				} else {
					builder.WriteRune(' ')
				}
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}
