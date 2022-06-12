package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Solution struct {
	rects [][]int
	sum   []int // 前缀
}

func Constructor(rects [][]int) Solution {
	sum := make([]int, len(rects)+1)
	for i, r := range rects {
		a, b, x, y := r[0], r[1], r[2], r[3]
		sum[i+1] = sum[i] + (x-a+1)*(y-b+1)
	}

	return Solution{rects, sum}
}

func (this *Solution) Pick() []int {
	k := rand.Intn(this.sum[len(this.sum)-1]) + 1 // 从全部点中选取一个数 k:[1,n]
	rectIndex := sort.SearchInts(this.sum, k) - 1 // 这个数在哪个矩阵中
	r := this.rects[rectIndex]
	a, b, x, _ := r[0], r[1], r[2], r[3]
	curIdx := k - this.sum[rectIndex] - 1 // 选取的点在当前矩阵中的第几个 [0,cur_N-1]
	da := curIdx % (x - a + 1)            // 11 % 3 == 2
	db := curIdx / (x - a + 1)            // 11 / 3 == 3

	return []int{a + da, b + db}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */

func main() {
	obj := Constructor([][]int{
		{-2, -2, 1, 1},
		{2, 2, 4, 6},
	})
	for i := 0; i < 3; i++ {
		x := obj.Pick()
		fmt.Printf("%+v\n", x)
	}
}
