package backtracking

/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells,
where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once.
*/

type Point struct {
	x, y int
}

func exist(board [][]byte, word string) bool {
	// keep track of path with a map
	visited := make(map[Point]bool)
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			// we start dfs if we find a match to the first char
			if board[r][c] == word[0] {
				if dfs(board, word, visited, Point{x: r, y: c}, 0) {
					return true
				}
			}
		}
	}
	// if nothing matches or no dfs returns true, then theres no word
	return false
}

func dfs(board [][]byte, word string, visited map[Point]bool, curr Point, index int) bool {
	// put base cases to make sure recursion stops
	// bound check - if out of bounds, return false
	if curr.x > len(board)-1 || curr.x < 0 || curr.y > len(board[curr.x])-1 || curr.y < 0 {
		return false
	}
	// check if we've already visited/if the point is in the path
	if visited[curr] {
		return false
	}
	// char doesn't match the target word index so return false
	if board[curr.x][curr.y] != word[index] {
		return false
	}
	// stop recursion bc we have found the word
	if index == len(word)-1 {
		return true
	}

	// process current node
	visited[curr] = true

	// explore neighbors
	neighbors := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	// iterate through neighbors
	for _, neighbor := range neighbors {
		neighPoint := Point{
			x: curr.x + neighbor[0],
			y: curr.y + neighbor[1],
		}
		// recurse on neighbor
		if dfs(board, word, visited, neighPoint, index+1) {
			// if we found a match, bubble up this call
			return true
		}
	}
	// didn't find a match on neighbor, so unmark curr to leave it open
	visited[curr] = false
	return false
}
