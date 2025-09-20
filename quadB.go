package quad

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This is your pattern generator function — it takes width and height and returns the string.
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

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./quadB width height")
		return
	}
	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Width and height must be integers")
		return
	}
	fmt.Print(QuadB(x, y))
}
