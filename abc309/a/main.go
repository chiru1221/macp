package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		check map[int][]int
		a, b  int
	)
	check = make(map[int][]int)
	check[1] = []int{2}
	check[2] = []int{1, 3}
	check[3] = []int{2}
	check[4] = []int{5}
	check[5] = []int{4, 6}
	check[6] = []int{5}
	check[7] = []int{8}
	check[8] = []int{7, 9}
	check[9] = []int{8}
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &a, &b)
	for _, value := range check[a] {
		if value == b {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
