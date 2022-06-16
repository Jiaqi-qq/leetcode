package main

// https://leetcode.cn/problems/average-salary-excluding-the-minimum-and-maximum-salary/
func average(salary []int) float64 {
	var result float64
	var max int
	min := int(^uint(0) >> 1)
	for _, v := range salary {
		result += float64(v)
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return (result - float64(max) - float64(min)) / float64(len(salary)-2)
}
