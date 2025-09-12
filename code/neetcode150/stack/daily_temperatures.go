package stack

/*
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you
have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.
*/

// the trick is that we can push the index of the temperature to a stack
// we iterate through the temps keeping track of our current index
// that index minus the index on the stack gives you how many days waited
// until you found a higher temperature
func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	res := make([]int, len(temperatures)) // each index is how many times that temp had to wait
	for i, temp := range temperatures {
		// if stack has values, need to check the values
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if temperatures[top] < temp {
				stack = stack[:len(stack)-1] // pop
				res[top] = i - top           // this is how many days we had to wait
			} else {
				break
			}
		}

		stack = append(stack, i)
	}
	return res
}
