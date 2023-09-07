package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	n := len(nums)

	numsMap := make(map[int]int)
	for _, num := range nums {
		numsMap[num]++
		if numsMap[num] > n/2 {
			return num
		}
	}
	return 0
}
