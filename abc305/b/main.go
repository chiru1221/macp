package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		p, q  string
		point map[string]float64
	)
	point = map[string]float64{
		"A": 0,
		"B": 3,
		"C": 4,
		"D": 8,
		"E": 9,
		"F": 14,
		"G": 23,
	}
	fmt.Scanf("%s %s", &p, &q)
	fmt.Println(math.Abs(point[p] - point[q]))
}
