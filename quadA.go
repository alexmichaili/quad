package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./program width height")
		return
	}
	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Width and height must be integers")
		return
	}
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= x; i++ { // i rows and j columns
		for j := 1; j <= y; j++ {
			if i == 1 || i == x {
				if j == 1 || j == y {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			} else {
				if j == 1 || j == y {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
