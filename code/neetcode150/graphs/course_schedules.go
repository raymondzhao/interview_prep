package graphs

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
 You are given an array prerequisites where prerequisites[i] = [a_i, b_i] indicates
 that you must take course bi first if you want to take course a_i.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return true if you can finish all courses. Otherwise, return false.
*/

// we can convert this into a graph if we think each edge representing pre-req -> course you can take
func canFinish(numCourses int, prerequisites [][]int) bool {
	// track visited state
	// 0 - not visited
	// 1 - visiting
	// 2 - visited
	visited := make([]int, numCourses)

	// build map with k: prereq, v: what course you can take
	// this creates a graph
	graph := make(map[int][]int)
	for _, pre := range prerequisites {
		a, b := pre[0], pre[1]
		graph[b] = append(graph[b], a)
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i, graph, visited) {
			return false
		}
	}
	return true
}

// recursion helper
func dfs(node int, graph map[int][]int, visited []int) bool {
	// if we are visiting as part of recursion, it's a cycle
	if visited[node] == 1 {
		return false
	}
	// if node has been visited/processed, return true since we dont want to repeat work
	if visited[node] == 2 {
		return true
	}

	// set current node to visiting
	visited[node] = 1

	// iterate through the "neighbors"
	for _, v := range graph[node] {
		// if anything reaches base case, we need to return false since it's a cycle
		if !dfs(v, graph, visited) {
			return false
		}
	}
	// if we can iterate through then we mark node as processed
	// this lets base case 2 work since we dont need to do redundant processing
	visited[node] = 2
	return true
}
