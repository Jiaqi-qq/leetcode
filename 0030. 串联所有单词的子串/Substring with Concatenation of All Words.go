package main

func findSubstring(s string, words []string) (ans []int) {
	ls, m, n := len(s), len(words), len(words[0])
	for i := 0; i < n && i+m*n <= ls; i++ { // i:第一个单词起始位置
		differ := map[string]int{}
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}

		if len(differ) == 0 {
			ans = append(ans, i)
		}

		for start := i + n; start < ls-m*n+1; start += n { // 窗口移动 start:第一个单词位置
			word := s[start-n : start] // 窗口左侧移除
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}

			word = s[start+(m-1)*n : start+m*n] // 窗口右侧新增
			differ[word]++
			if differ[word] == 0 {
				delete(differ, word)
			}

			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}

	return ans
}
