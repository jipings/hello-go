
package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	var m = len(matrix)
	if m == 0 {
		return false
	}
	var n = len(matrix[0])
	if n == 0 {
		return false
	}
	var i,j = m - 1, 0

	for i >=0 && j < n {
		if matrix[i][j] == target {
			return true
		} 
		if matrix[i][j] < target {
			j = j + 1
		}else {
			i = i - 1
		}
	}
	return false
}

func main() {
	var matrix = [][]int{
        {1,   4,  7, 11, 15},
		{2,   5,  8, 12, 19},
		{3,   6,  9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	fmt.Println(searchMatrix(matrix,5))
}