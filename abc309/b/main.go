package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		n     int
		a, b  [][]int
		input string
	)
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	a = make([][]int, n)
	b = make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		b[i] = make([]int, n)
		fmt.Fscan(r, &input)
		for j := range a {
			a[i][j] = int(input[j] - '0')
		}
		copy(b[i], a[i])
	}
	for i := range a {
		if i == 0 {
			for j := range a {
				if j == 0 {
					b[i][j] = a[i+1][j]
				} else {
					b[i][j] = a[i][j-1]
				}
			}
		} else if i == n-1 {
			for j := range a {
				if j == n-1 {
					b[i][j] = a[i-1][j]
				} else {
					b[i][j] = a[i][j+1]
				}
			}
		} else {
			b[i][0] = a[i+1][0]
			b[i][n-1] = a[i-1][n-1]
		}
	}
	for i := range b {
		for j := range b {
			fmt.Print(b[i][j])
		}
		fmt.Println()
	}
}
