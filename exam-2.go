package main

import (
	"fmt"
)

// convert encoded string L,R,= to numbers
func decode(encoded string) string {
	n := len(encoded)
	nums := make([]int, n+1)

	// iterate from left to right
	for i := 0; i < n; i++ {
		if encoded[i] == 'L' {
			nums[i] = nums[i+1] + 1
		} else if encoded[i] == 'R' {
			nums[i+1] = nums[i] + 1
		} else { // =
			nums[i+1] = nums[i]
		}
	}

	// iterate from right to left
	for i := n - 1; i >= 0; i-- {
		if encoded[i] == 'L' {
			if nums[i] <= nums[i+1] {
				nums[i] = nums[i+1] + 1
			}
		} else if encoded[i] == 'R' {
			if nums[i] >= nums[i+1] {
				nums[i+1] = nums[i] + 1
			}
		} else { // =
			nums[i] = nums[i+1]
		}
	}

	// find the minimum value
	min := nums[0]
	for i := 1; i <= n; i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	// subtract the minimum value from all elements
	for i := 0; i <= n; i++ {
		nums[i] = nums[i] - min
	}

	// convert to string
	ans := ""
	for i := 0; i <= n; i++ {
		ans = ans + fmt.Sprint(nums[i])
	}
	return ans
}
