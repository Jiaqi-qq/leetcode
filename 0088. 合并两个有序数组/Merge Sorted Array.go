package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, idx := m-1, n-1, m+n-1; i >= 0 || j >= 0; {
		if i < 0 || (j >= 0 && nums1[i] < nums2[j]) { // nums1空了 or 比较. 注意 j>=0是防止j==-1 nums2已经空了 应该放nums1 但是判断在下面else if
			nums1[idx] = nums2[j]
			j--
		} else if j < 0 || (i >= 0 && nums1[i] >= nums2[j]) { // 其余情况都走这里 可以改成else
			nums1[idx] = nums1[i]
			i--
		}
		idx--
	}
}
