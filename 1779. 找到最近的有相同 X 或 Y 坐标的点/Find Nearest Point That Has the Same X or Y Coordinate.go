package main

func nearestValidPoint(x int, y int, points [][]int) int {
	idx, ans := -1, 0
	for i := 0; i < len(points); i++ {
		if points[i][0] == x || points[i][1] == y {
			d := abs((x - points[i][0]) + (y - points[i][1]))
			if idx == -1 || d < ans {
				idx, ans = i, d
			}
		}
	}
	return idx
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
