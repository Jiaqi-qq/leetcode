package main

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, 0, m*n)
	for i := 0; i < m+n-1; i++ {
		if i%2 == 1 { // 右上 -> 左下
			y := min(i, n-1) // 起点 y 坐标 一定是 i 或者 最后一列
			// x := max(i-n+1, 0)
			x := i - y
			for x < m && y >= 0 {
				ans = append(ans, mat[x][y])
				x++
				y--
			}
		} else { // 左下 -> 右上
			x := min(i, m-1) // 起点 x 坐标 一定是 i 或者 最后一行
			// y := max(i-m+1, 0)
			y := i - x
			for x >= 0 && y < n {
				ans = append(ans, mat[x][y])
				x--
				y++
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
