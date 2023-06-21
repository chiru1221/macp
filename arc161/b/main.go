package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var (
		t              int
		n              float64
		p2List         []float64
		searchAbleList sort.Float64Slice
		sum            float64
		selectCount    int
	)
	restDiff := map[int]float64{0: 3, 1: 1, 2: 0}
	p2List = make([]float64, 60)
	for i := 0; i < 60; i++ {
		result := float64(math.Pow(2, float64(i)))
		p2List[i] = result
	}
	searchAbleList = sort.Float64Slice(p2List)
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%f", &n)
		// fmt.Println(n)
		if n < 7 {
			fmt.Println(-1)
			continue
		}
		sum = 0
		selectCount = 0
		for j := searchAbleList.Search(n) - 1; j >= 0; j-- {
			if selectCount == 3 {
				break
			}
			if n-(p2List[j]+sum) >= restDiff[selectCount] {
				sum += p2List[j]
				selectCount++
			}
		}
		fmt.Println(uint64(sum))
	}
}
