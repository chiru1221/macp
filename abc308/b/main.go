package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		n, m int
		c, d []string
		p    []int
	)
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	c = make([]string, n)
	d = make([]string, m)
	p = make([]int, m+1)
	for i := range c {
		fmt.Fscan(r, &c[i])
	}
	for i := range d {
		fmt.Fscan(r, &d[i])
	}
	for i := range p {
		fmt.Fscan(r, &p[i])
	}
	price := make(map[string]int)
	for i, di := range d {
		price[di] = p[i+1]
	}
	sum := 0
	for _, ci := range c {
		v, ok := price[ci]
		if ok {
			sum += v
		} else {
			sum += p[0]
		}
	}
	fmt.Println(sum)
}
