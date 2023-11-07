package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)

	for i, num := range nums {
		tmp := target - num
		if _, ok := maps[tmp]; ok {
			return []int{maps[tmp], i}
		}
		maps[num] = i
	}

	return nil
}
