package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s int
	r := bufio.NewReader(os.Stdin)
	pre := 0
	for i := 0; i < 8; i++ {
		fmt.Fscan(r, &s)
		if pre <= s && 100 <= s && s <= 675 && s%25 == 0 {
			pre = s
			continue
		} else {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
