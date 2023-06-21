package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n     int
		a     int
		keys  []int
		f     map[int]int
		count map[int]int
	)
	fmt.Scan(&n)
	f = make(map[int]int)
	count = make(map[int]int)
	keys = make([]int, n)
	for i := 1; i <= n; i++ {
		count[i] = 0
		keys[i-1] = i
	}
	for i := 0; i < 3*n; i++ {
		fmt.Scan(&a)
		count[a]++
		if count[a] == 2 {
			f[a] = i
		}
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return f[keys[i]] < f[keys[j]]
	})
	for i, key := range keys {
		fmt.Print(key)
		if i == n-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
