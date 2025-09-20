// quadA.go
package quad

import "strings"

func QuadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var builder strings.Builder
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 || i == y) && (j == 1 || j == x) {
				builder.WriteRune('o')
			} else if i == 1 || i == y {
				builder.WriteRune('-')
			} else if j == 1 || j == x {
				builder.WriteRune('|')
			} else {
				builder.WriteRune(' ')
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}
