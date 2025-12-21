package main

import "fmt"

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	return comma(s[:len(s)-3]) + "," + s[len(s)-3:] 
}

func main() {
	nums := []string{
		"123456",
		"87623",
		"22",
	}
	for i := range len(nums) {
		fmt.Printf("n: %s, n with comma: %s\n", nums[i], comma(nums[i]))
	}
}
