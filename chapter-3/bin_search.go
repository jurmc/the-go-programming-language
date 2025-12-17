package main

import "fmt"

func search(nums []int, target int) int {
	p1, p2 := 0, len(nums)-1

	for true {
		if nums[p1] == target { return p1 }
		if nums[p2] == target { return p2 }
		if p1 + 1 >= p2 { return -1 }

		mid := (p2 - p1)/2 + p1
		if nums[mid] >= target {
			p2 = mid
		} else {
			p1 = mid
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 2, 4, 6, 8}
	fmt.Println(search(nums, 4))
}
