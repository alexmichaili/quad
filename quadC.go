package quad

import "strings"

func QuadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var builder strings.Builder
	for i := 1; i <= y; i++ { // i rows and j columns
		for j := 1; j <= x; j++ {
			if i == 1 && (j == 1 || j == x) {
				builder.WriteRune('A')
			} else if i == y && (j == 1 || j == x) {
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
