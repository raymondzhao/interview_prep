package heap

/*
Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane
and an integer k, return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance
(i.e., √(x1 - x2)2 + (y1 - y2)2).

You may return the answer in any order. The answer is guaranteed to be unique
(except for the order that it is in).
*/
import (
	"container/heap"
)

type Point struct {
	x, y int
	dist int
}

// setup max heap using heap pkg
type MaxHeap []Point

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].dist > h[j].dist } // max-heap
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Point)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// use a max heap to store k number of points and then return the results from the heap
func kClosest(points [][]int, k int) [][]int {
	// create max heap pointer - use heap pkg to initialize
	h := &MaxHeap{}
	heap.Init(h)

	// iterate through points and push to the heap
	for _, v := range points {
		dist := v[0]*v[0] + v[1]*v[1]
		heap.Push(h, Point{v[0], v[1], dist})
		// pop bc we only care about smallest dist values so remove higher dist
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// prepare proper return of [][]int
	res := [][]int{}
	for h.Len() > 0 {
		p := heap.Pop(h).(Point)
		res = append(res, []int{p.x, p.y})
	}
	return res
}
