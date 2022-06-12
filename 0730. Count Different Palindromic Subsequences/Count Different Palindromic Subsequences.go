package main

func countPalindromicSubsequences(s string) int {
	const mod int = 1e9 + 7
	n := len(s)
	dp := make([][]int, n)
	next := make([][]int, n)
	prev := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
		next[i] = make([]int, 4)
		prev[i] = make([]int, 4)
	}

	pos := make([]int, 4)
	for i := 0; i < 4; i++ {
		pos[i] = -1 // 默认不存在指向 -1
	}
	for i := 0; i < n; i++ {
		for c := 0; c < 4; c++ {
			prev[i][c] = pos[c]

		}
		pos[s[i]-'a'] = i
	}

	for i := 0; i < 4; i++ {
		pos[i] = n // 默认不存在指向 n
	}
	for i := n - 1; i >= 0; i-- {
		for c := 0; c < 4; c++ {
			next[i][c] = pos[c]
		}
		pos[s[i]-'a'] = i
	}

	for sz := 2; sz <= n; sz++ {
		for i, j := 0, sz-1; j < n; i++ {
			if s[i] == s[j] {
				low, high := next[i][s[i]-'a'], prev[j][s[i]-'a']
				if low > high {
					dp[i][j] = (2 + dp[i+1][j-1]*2) % mod // 没有重复
				} else if low == high {
					dp[i][j] = (1 + dp[i+1][j-1]*2) % mod // 中间一个元素，只会有一个重复 -1 即可
				} else {
					dp[i][j] = (dp[i+1][j-1]*2 - dp[low+1][high-1] + mod) % mod
				}
			} else {
				dp[i][j] = (dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1] + mod) % mod
			}
			j++
		}
	}

	return dp[0][n-1]
}

func main() {

}
