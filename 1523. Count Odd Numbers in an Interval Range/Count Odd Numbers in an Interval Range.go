package main

//https://leetcode.cn/problems/count-odd-numbers-in-an-interval-range/
func countOdds(low int, high int) int {
	return get(high) - get(low-1)
}

func get(x int) int {
	return (x + 1) >> 1
}
