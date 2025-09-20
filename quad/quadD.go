package main

import "strings"

func QuadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var builder strings.Builder
	for i := 1; i <= y; i++ { // i rows and j columns
		for j := 1; j <= x; j++ {
			if j == 1 && (i == 1 || i == y) {
				builder.WriteRune('A')
			} else if j == x && (i == 1 || i == y) {
				builder.WriteRune('C')
			} else if ((i == 1 || i == y) && (j > 1 && j < x)) || (i > 1 && i < y && (j == 1 || j == x)) {
				builder.WriteRune('B')
			} else {
				builder.WriteRune(' ')
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}
