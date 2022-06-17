package main

// 思路：i左边有几个0 i就要往右移动几位
func duplicateZeros1(arr []int) {
	cnt0 := 0
	// 统计0的个数
	for _, v := range arr {
		if v == 0 {
			cnt0++
		}
	}
	// 逆向遍历 覆盖的都是无效数据
	for i := len(arr) - 1; i >= 0; i-- {
		// 遇到0就减去 得到的就是当前位置左边0的数量
		if arr[i] == 0 {
			cnt0--
		}
		if i+cnt0 < len(arr) {
			arr[i+cnt0] = arr[i]
			// 如果遇到0 移动后的位置下一个元素也是0
			if arr[i] == 0 && i+cnt0+1 < len(arr) {
				arr[i+cnt0+1] = 0
			}
		}
	}
}

// 双指针模拟栈
func duplicateZeros2(arr []int) {
	top, i := 0, 0
	for top < len(arr) {
		if arr[i] == 0 {
			top++
		}
		i++
		top++
	}

	// i 和 top 都是右侧开区间位置 -1 才是元素位置
	i--
	top--

	// if top == len(arr) { // 特殊判断 top越界情况 就无需下文判断 top < len(arr)
	// 	i--
	// 	top--
	// 	arr[top] = 0
	// 	top--
	// }

	for i >= 0 {
		if top < len(arr) { // top 有可能越界 因为会有遇到0移动两次出界的极端情况
			arr[top] = arr[i]
		}
		if arr[i] == 0 {
			top--
			arr[top] = 0
		}
		i--
		top--
	}
}
