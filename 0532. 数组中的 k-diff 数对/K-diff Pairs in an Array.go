package main

import "sort"

func findPairs1(nums []int, k int) int { // 哈希表
	visited := map[int]struct{}{}
	res := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := visited[num-k]; ok {
			res[num-k] = struct{}{}
		}
		if _, ok := visited[num+k]; ok {
			res[num] = struct{}{}
		}
		visited[num] = struct{}{}
	}
	return len(res)
}

func findPairs2(nums []int, k int) (ans int) { // 排序 双指针
	sort.Ints(nums)
	y, n := 0, len(nums)
	for x, num := range nums {
		if x == 0 || num != nums[x-1] {
			for y < n && (nums[y] < num+k || y <= x) {
				y++
			}
			if y < n && nums[y] == num+k {
				ans++
			}
		}
	}
	return
}
