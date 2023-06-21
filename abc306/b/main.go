package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a      uint64
		result uint64
	)

	for i := 0; i < 64; i++ {
		fmt.Scan(&a)
		result += a * uint64(math.Pow(2, float64(i)))
	}
	fmt.Println(result)
}
