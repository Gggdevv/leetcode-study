package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maps := make(map[int]bool)

	for _, num := range nums {
		maps[num] = true
	}

	long := 1
	for num, _ := range maps {
		if !maps[num-1] {
			nowLong := 1
			cnt := 1
			for maps[num+cnt] {
				nowLong += 1
				cnt += 1
			}
			if nowLong > long {
				long = nowLong
			}
		}

	}

	return long
}
