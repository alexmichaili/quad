package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func QuadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var builder strings.Builder
	for i := 1; i <= y; i++ { // rows
		for j := 1; j <= x; j++ { // columns
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

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./quadA width height")
		return
	}
	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Width and height must be integers")
		return
	}
	fmt.Print(QuadA(x, y))
}
