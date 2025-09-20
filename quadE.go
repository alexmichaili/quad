package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func QuadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var builder strings.Builder
	for i := 1; i <= y; i++ { // i rows and j columns
		for j := 1; j <= x; j++ {
			if (j == 1 && i == 1) || (j == x && y != 1 && i == y && j != 1) {
				builder.WriteRune('A')
			} else if (j == 1 && i == y) || (j == x && i == 1) {
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

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./quadE width height")
		return
	}
	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Width and height must be integers")
		return
	}

	fmt.Print(QuadE(x, y))
}
