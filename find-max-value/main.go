package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{16, 6, 42, 4, 23, 15, 88, 45}

	max := 0

	//Option A : looping
	for _, v := range nums {

		if v > max {
			max = v
		}
	}

	fmt.Printf(" Option A : Max value = %v", max)

	//Option B : sorting
	sort.Ints(nums)

	max = nums[len(nums)-1]

	fmt.Printf(" Option B : Max value = %v", max)
}
