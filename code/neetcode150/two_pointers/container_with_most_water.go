package twopointers

/*
You are given an integer array height of length n.
There are n vertical lines drawn such that the two
endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container,
such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/

// this problem sucks to understand without a visual but
// area = height * width
// width = xn - xk so that x is the biggest
// height = min (yn, yk)) so that the most water can fit

// need to do two pointer approach and calculate area as we go along
func maxArea(height []int) int {
	// start and end pointers
	start, end := 0, len(height)-1
	maxArea := 0

	for start < end {
		// find width, height, area
		w := end - start
		h := min(height[start], height[end])
		area := w * h

		// compare with our max area
		if area > maxArea {
			maxArea = area
		}

		// the smaller height will yield the least area so move pointer toward center
		if height[start] < height[end] {
			start += 1
		} else {
			end -= 1
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
