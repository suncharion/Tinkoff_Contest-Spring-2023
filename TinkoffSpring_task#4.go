package main

import (
	"fmt"
)

func dict(d map[int]int, allowed int) bool {
	found := false
	for _, item := range d {
		if allowed-1 <= item && item <= allowed+1 {
			if item != allowed {
				if found {
					return false
				}
				found = true
			}
		} else {
			return false
		}
	}
	return true
}

func max(_len int, nums []int) int {
	maps := make(map[int]int)
	for _, item := range nums {
		maps[item]++
	}
	for i := _len - 1; i > 0; i-- {
		if dict(maps, maps[nums[i]]) {
			return i + 1
		}
		maps[nums[i]]--
		if maps[nums[i]] == 0 {
			delete(maps, nums[i])
		}
	}
	return 0
}

func main() {
	var len int
	fmt.Scanln(&len)

	Slices := make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Scan(&Slices[i])
	}

	fmt.Println(max(len, Slices))
}
