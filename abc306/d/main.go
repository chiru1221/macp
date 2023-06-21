package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int64) int64 {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	var (
		n  int
		x  bool
		y  int64
		dp [][]int64
		r  *bufio.Reader
	)
	r = bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	dp = make([][]int64, n+1)
	dp[0] = make([]int64, 2)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int64, 2)
		fmt.Fscan(r, &x, &y)
		// no-eat
		dp[i][0] = dp[i-1][0]
		dp[i][1] = dp[i-1][1]
		// eat
		if x {
			dp[i][1] = max(dp[i-1][0]+y, dp[i][1])
		} else {
			dp[i][0] = max(dp[i][0], max(dp[i-1][1], dp[i-1][0])+y)
		}
	}
	fmt.Println(max(dp[n][0], dp[n][1]))
}
