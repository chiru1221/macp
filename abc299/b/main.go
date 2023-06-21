package main

import (
	"fmt"
)

type Player struct {
	User  int
	Value int
}

func main() {
	var (
		n, t         int
		c, r         []int
		tInfo, fInfo *Player
	)
	fmt.Scanf("%d %d", &n, &t)
	c = make([]int, n)
	r = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &c[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &r[i])
	}
	tInfo = &Player{User: -1, Value: 0}
	fInfo = &Player{User: -1, Value: 0}
	for i := 0; i < n; i++ {
		if c[i] == t {
			if tInfo.Value < r[i] {
				tInfo.User = i
				tInfo.Value = r[i]
			}
		} else if c[i] == c[0] {
			if fInfo.Value < r[i] {
				fInfo.User = i
				fInfo.Value = r[i]
			}
		}
	}
	if tInfo.User != -1 {
		fmt.Println(tInfo.User + 1)
	} else {
		fmt.Println(fInfo.User + 1)
	}
}
