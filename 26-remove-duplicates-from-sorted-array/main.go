package main

func main() {
	removeDuplicates([]int{1, 1, 2})
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	len := len(nums)

	if len < 2 {
		return 1
	}

	i := 0
	for j := 1; j < len; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
