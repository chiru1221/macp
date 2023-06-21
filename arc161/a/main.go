package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n     int
		count map[int]int
		ai    int
	)
	fmt.Scanf("%d", &n)
	if n == 1 {
		fmt.Println("Yes")
		return
	}
	count = make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ai)
		count[ai] += 1
	}
	keys := make([]int, len(count))
	idx := 0
	for k, _ := range count {
		keys[idx] = k
		idx++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	sum := 0
	stock := 0
	th := (n - 1) / 2
	for _, key := range keys {
		sum += count[key]
		if sum >= th {
			rest := count[key] - (th - stock)
			if rest <= stock {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
			break
		}
		stock = sum
	}
}
