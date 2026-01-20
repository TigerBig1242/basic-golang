package main

import (
	"fmt"
)

func plusIndex(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		need := target - nums[i]
		index, ok := indexMap[need]
		if ok {
			return []int{index, i}
		}
		indexMap[nums[i]] = i
	}
	return []int{}
}

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	result := plusIndex(nums, target)
	fmt.Println("result index :", result)
}
