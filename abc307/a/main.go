package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		n int
		a int
	)
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < 7; j++ {
			fmt.Fscan(r, &a)
			sum += a
		}
		fmt.Print(sum)
		if i != n-1 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
}
