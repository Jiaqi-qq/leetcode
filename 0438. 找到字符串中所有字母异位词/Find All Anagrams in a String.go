package main

func findAnagrams(s string, p string) []int {
	return findAnagrams2(s, p)
}

// 滑动窗口 维护窗口内各个字母的个数 如果pCount和sCount相同说明此时异构
func findAnagrams1(s string, p string) (ans []int) {
	if len(s) < len(p) {
		return
	}

	i, j := 0, 0 // [i,j)
	var sCount, pCount [26]int
	for j < len(p) {
		sCount[s[j]-'a']++
		pCount[p[j]-'a']++
		j++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for j < len(s) {
		sCount[s[i]-'a']--
		i++
		sCount[s[j]-'a']++
		if sCount == pCount {
			ans = append(ans, i)
		}
		j++
	}

	return
}

// 滑动窗口优化 维护两个count的差differ 比较时只需O(1)判断differ==0即可 不需要对count数组进行遍历
func findAnagrams2(s string, p string) (ans []int) {
	if len(s) < len(p) {
		return nil
	}

	i, j, differ := 0, 0, 0
	count := [26]int{}
	for j < len(p) {
		count[s[j]-'a']++
		count[p[j]-'a']--
		j++
	}

	for _, c := range count {
		if c != 0 {
			differ++
		}
	}
	if differ == 0 {
		ans = append(ans, 0)
	}

	for j < len(s) {
		if count[s[i]-'a'] == 1 { // 当前字符 s-p 即将从:1->0 即:不同->相同
			differ--
		} else if count[s[i]-'a'] == 0 { // 相同变不同
			differ++
		}
		count[s[i]-'a']-- // 窗口移动 左侧移除
		i++

		if count[s[j]-'a'] == -1 {
			differ--
		} else if count[s[j]-'a'] == 0 {
			differ++
		}
		count[s[j]-'a']++ // 窗口移动 右侧新加
		j++

		if differ == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
