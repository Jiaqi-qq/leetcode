package leetcode

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for i, num := range nums {
		if idx, ok := set[num]; ok {
			return []int{idx, i}
		}
		set[target-num] = i
	}
	return nil
}
