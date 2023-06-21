package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		n int
		s string
	)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &s)
	if len(s) == 1 {
		fmt.Println(-1)
		return
	}
	if !strings.Contains(s, "-") || !strings.Contains(s, "o") {
		fmt.Println(-1)
		return
	}
	ans := 1
	for _, dango := range strings.Split(s, "-") {
		if l := len(dango); ans < l {
			ans = l
		}
	}
	fmt.Println(ans)
}
