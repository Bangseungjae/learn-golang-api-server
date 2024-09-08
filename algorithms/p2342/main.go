package main

import (
	"bufio"
	"os"
	"strconv"
)

var move = [][]int{
	{0, 2, 2, 2, 2},
	{0, 1, 3, 4, 3},
	{0, 3, 1, 3, 4},
	{0, 4, 3, 1, 3},
	{0, 3, 4, 3, 1},
}

const big = 1<<30 - 1

func main() {
	defer wr.Flush()

	var dp [5][5]int
	var ndp [5][5]int
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = big
			ndp[i][j] = big
		}
	}
	dp[0][0] = 0

	for {
		x := scanInt()
		if x == 0 {
			break
		}
		for i := 0; i <= 4; i++ {
			if i != x {
				for j := 0; j <= 4; j++ {
					ndp[x][i] = min(ndp[x][i], dp[j][i]+move[j][x])
					ndp[i][x] = min(ndp[i][x], dp[i][j]+move[j][x])
				}
			}
		}

		for i := 0; i <= 4; i++ {
			for j := 0; j <= 4; j++ {
				dp[i][j] = ndp[i][j]
				ndp[i][j] = big
			}
		}
	}

	m := big
	for i := range dp {
		for j := range dp[i] {
			m = min(m, dp[i][j])
		}
	}
	wr.WriteString(strconv.Itoa(m))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	wr = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}
