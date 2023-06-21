package main

import (
	"fmt"
)

func main() {
	var (
		n int
		w string
	)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &w)
	var (
		start int
		end   int
		ast   int
	)

	start = -1
	end = -1
	for i := 0; i < n; i++ {
		if w[i] == '|' {
			if start == -1 {
				start = i
			} else {
				end = i
			}
		} else if w[i] == '*' {
			ast = i
		}
	}

	if start < ast && ast < end {
		fmt.Println("in")
	} else {
		fmt.Println("out")
	}
}
