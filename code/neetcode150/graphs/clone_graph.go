package graphs

/*
Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}

Test case format:

For simplicity, each node's value is the same as the node's index (1-indexed).
For example, the first node with val == 1, the second node with val == 2, and so on.
The graph is represented in the test case using an adjacency list.

An adjacency list is a collection of unordered lists used to represent a finite graph.
Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1.
You must return the copy of the given node as a reference to the cloned graph.
*/

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// create a map storing original node to the clone
	origClone := make(map[*Node]*Node)
	graph := []*Node{node} // track what we have seen so we do not infinitely traverse

	// create clone of node and add it to our map
	origClone[node] = &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}

	// traverse the graph with BFS
	for len(graph) > 0 {
		curr := graph[0]
		graph = graph[1:] //dequeue

		currClone := origClone[curr]

		// iterate through the neighbors
		if len(curr.Neighbors) > 0 {
			for _, neighbor := range curr.Neighbors {
				// check if we have a clone of the neighbor
				neighborClone, exists := origClone[neighbor]
				if !exists {
					// no clone so we gotta make one
					neighborClone = &Node{
						Val:       neighbor.Val,
						Neighbors: []*Node{},
					}
					// pair the neighbor with its clone
					origClone[neighbor] = neighborClone
					graph = append(graph, neighbor) // add neighbor so we can visit this node via BFS
				}
				// add the neighbor clone to the curr clone's neighbors
				currClone.Neighbors = append(currClone.Neighbors, neighborClone)
			}
		}
	}
	return origClone[node]
}
