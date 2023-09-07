package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	length := len(nums)
	i := 0
	flag := false
	for j := 1; j < length; j++ {
		if !flag {
			if nums[i] == nums[j] {
				i++
				nums[i] = nums[j]
				flag = true
			} else {
				i++
				nums[i] = nums[j]
			}
		}

		if flag && nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
			flag = false
		}

	}

	return i + 1
}
