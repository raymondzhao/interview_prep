package binarysearch

/*
You are given an m x n integer matrix matrix with the following two properties:

    - Each row is sorted in non-decreasing order.
    - The first integer of each row is greater than the last integer of the previous row.

Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

Example 1:
	Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
	Output: true

*/

func searchMatrix(matrix [][]int, target int) bool {
	// searching a matrix is just searching a long 1d array
	// we just need to flatten it or treat it flattened
	// need to convert from flattened index to matrix coordinates
	rows := len(matrix)
	cols := len(matrix[0])

	min := 0
	max := rows*cols - 1

	// need to check if min = max so we can still check the middle
	for min <= max {
		mid := (min + max) / 2
		x, y := indexToMatrixLoc(mid, cols)
		val := matrix[x][y]

		if val > target {
			// need to decrement since we just checked mid
			max = mid - 1
		} else if val < target {
			// need to increment since we just checked mid
			min = mid + 1
		} else {
			return true
		}
	}
	return false
}

func indexToMatrixLoc(num int, colSize int) (x, y int) {
	// helper function to convert flattened array index to x,y matrix coords
	x = num / colSize
	y = num % colSize
	return
}
