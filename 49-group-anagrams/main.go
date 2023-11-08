package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams(([]string{"eat", "tea", "tan", "ate", "nat", "bat"})))
}

func groupAnagrams(strs []string) [][]string {

	maps := make(map[string][]string)

	for _, str := range strs {
		s := []byte(str)

		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})

		strr := string(s)

		maps[strr] = append(maps[strr], str)
	}

	res := [][]string{}
	for _, v := range maps {
		res = append(res, v)
	}

	return res
}
