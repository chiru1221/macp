package main

import "fmt"

func main() {
	var (
		h, w int
		s    [][]byte
	)
	fmt.Scanf("%d %d", &h, &w)
	s = make([][]byte, h)
	for i, _ := range s {
		s[i] = make([]byte, w)
		fmt.Scan(&s[i])
	}
	for i, si := range s {
		for j, sj := range si {
			count := 0
			if sj == '.' {
				if i > 0 {
					// fmt.Println(s[i-1][j])
					if s[i-1][j] == '#' {
						count++
					}
				}
				if i < h-1 {
					// fmt.Println(s[i+1][j])
					if s[i+1][j] == '#' {
						count++
					}
				}
				if j > 0 {
					// fmt.Println(s[i][j-1])
					if s[i][j-1] == '#' {
						count++
					}
				}
				if j < w-1 {
					// fmt.Println(s[i][j+1])
					if s[i][j+1] == '#' {
						count++
					}
				}
			}
			if count > 1 {
				fmt.Printf("%d %d\n", i+1, j+1)
				return
			}
		}
	}
}
