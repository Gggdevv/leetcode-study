package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

// https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func merge(nums1 []int, m int, nums2 []int, n int) {

	len1 := m + n

	end1 := m - 1
	end2 := n - 1

	for i := len1 - 1; i >= 0; i-- {
		if end1 >= 0 && (end2 < 0 || nums1[end1] > nums2[end2]) {
			nums1[i] = nums1[end1]
			end1--
		} else {
			nums1[i] = nums2[end2]
			end2--
		}
	}

	fmt.Println(nums1)

}
