package dp

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/

// dp programs require memoization
// climbing up to n steps happens two ways - 1 + (n-1) and 2 + (n-2)
// because you might need an extra 1 step and an extra 2 step
// n = (n-1) + (n-2)
func climbStairs(n int) int {
	// curr is named that way bc our answer will be in curr
	prev, curr := 1, 2 // prev and curr store n-2, n-1,

	// general "base cases" but not really
	if n == 1 {
		return prev
	}
	if n == 2 {
		return curr
	}
	// calculate up to n
	for i := 3; i <= n; i++ {
		next := prev + curr // calculate n
		prev = curr         // assign new values to a and
		curr = next
	}
	return curr
}
