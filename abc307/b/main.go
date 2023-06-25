package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		n int
		s []string
	)
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	s = make([]string, n)
	for i := range s {
		fmt.Fscan(r, &s[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			c := s[i] + s[j]
			if isLoopSentense(c) {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}

func isLoopSentense(s string) bool {
	m := len(s)
	for i := 0; i < m/2; i++ {
		if s[i] != s[m-1-i] {
			return false
		}
	}
	return true
}
