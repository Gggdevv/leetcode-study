package main

func main() {
	removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
}

// https://leetcode.cn/problems/remove-element/description/?utm_source=LCUS&utm_medium=ip_redirect&utm_campaign=transfer2china
func removeElement(nums []int, val int) int {
	len := len(nums)
	i := 0
	j := len - 1
	res := 0
	for i <= j {
		if i < len && j >= 0 && nums[i] == val && nums[j] != val {
			nums[i] = nums[j]
			i++
			j--
			res++
		}
		if nums[i] != val {
			i++
		}
		if nums[j] == val {
			j--
			res++
		}
	}

	return len - res
}
