package _189_rotate_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	require.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, nums)
}

// https://leetcode.cn/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[k:])
	reverse(nums[:k])
}

func reverse(arr []int) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		temp := arr[i]
		arr[i] = arr[length-1-i]
		arr[length-1-i] = temp
	}
}
