package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countStudents([]int32{1, 1, 3, 3, 4, 1}))
}

func countStudents(height []int32) int32 {
	var result int32 = 0
	var idx int = 0
	var soretedHeights []int32 = make([]int32, len(height))
	copy(soretedHeights, height)

	sort.Slice(soretedHeights, func(i, j int) bool {
		return soretedHeights[i] < soretedHeights[j]
	})

	for _, h := range soretedHeights {
		if height[idx] != h {
			result++
		}
		idx++
	}

	return result

}
