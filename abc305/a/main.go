package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	if n == 100 {
		fmt.Println(100)
		return
	}
	var rest int = n % 5
	if rest == 0 {
		fmt.Println(n)
		return
	}
	if rest < 3 {
		fmt.Println(n - rest)
		return
	} else {
		fmt.Println(n + (5 - rest))
		return
	}
}
