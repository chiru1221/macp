package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &s)
	for _, si := range s {
		fmt.Print(string(si) + string(si))
	}
	fmt.Println()
}
