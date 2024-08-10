Create a program to collect many items based on the given target point.

Test Case 1:


Input:

Target: 80
Items: 10, 10, 5, 30, 40, 10, 5
Output:

[40, 30, 10]

Explanation: the target point is 80 which means some items that can be collected to achieve the target are [40, 30, 10]


Test Case 2:


Input:

Target: 100
Items: 50, 20, 10, 25, 25
        Output:

[50 25 25]
        

package main

import (
  "fmt"
)

func main() {
  fmt.Println(getItem([]int{10, 10, 5, 30, 40, 10, 5}, 80)) // [40 30 10]
  fmt.Println(getItem([]int{50, 20, 10, 25, 25}, 100))      // [50 25 25]
}

func getItem(items []int, target int) []int {
  //TODO: your code here
  return []int{}
}