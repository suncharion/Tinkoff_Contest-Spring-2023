package main

import (
	"fmt"
)

func sumArr(nums []int, k int) int {
	seen := make(map[int]int)
	seen[0] = 1
	pref, res := 0, 0
	for _, item := range nums {
		pref += item
		curTarget := pref - k
		fromSeen := seen[curTarget]
		res += fromSeen
		seen[pref]++
	}
	return res
}

func answer(dists [][]int, lenNums int) int {
	parm := make(map[[2]int]bool)
	for _, dist := range dists {
		start, end := dist[0], dist[1]
		for i := 0; i <= start; i++ {
			parm[[2]int{i, end}] = true
		}
		for i := end; i < lenNums; i++ {
			parm[[2]int{start, i}] = true
		}
	}
	return len(parm)
}

func analysis(nums []int, k int) int {
	seen := make(map[int][]int)
	seen[0] = []int{-1}
	pref := 0
	good := [][]int{}
	for i := 0; i < len(nums); i++ {
		pref += nums[i]
		if _, ok := seen[pref]; !ok {
			seen[pref] = []int{i}
		}
		if _, ok := seen[pref-k]; ok {
			for _, option := range seen[pref-k] {
				if i-option > 0 {
					good = append(good, []int{option + 1, i})
					seen[pref-k] = append(seen[pref-k], i)
				}
			}
		}
	}
	return answer(good, len(nums))
}

func main() {

	var length int
	fmt.Println("Enter the length of the slice:")
	fmt.Scan(&length)

	slice := make([]int, length)

	fmt.Println("Enter", length, "int values:")
	for i := 0; i < length; i++ {
		fmt.Scan(&slice[i])
	}
	fmt.Println(analysis(slice, 0))

}
