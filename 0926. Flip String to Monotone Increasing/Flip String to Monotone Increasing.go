package main

func minFlipsMonoIncr(s string) int {
	dp, cnt := 0, 0
	for _, c := range s {
		if c == '1' {
			cnt++
		} else {
			dp = min(dp+1, cnt)
		}
	}
	return dp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
