package binarysearch

/*
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas.
The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k.
Each hour, she chooses some pile of bananas and eats k bananas from that pile.
If the pile has less than k bananas, she eats all of them instead
and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.
*/

// we can frame this as a binary search problem - our speed can only live between 1
// and the biggest pile

// if we pick a rate in the middle, we can calculate how many hours it'll take to clear the piles
// and then compare with the hours value given to us - if we need more time, k goes up and vice versa
func minEatingSpeed(piles []int, h int) int {
	// find the max pile
	min, max := 1, 1 // Koko's slowest eating speed is 1
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	for min < max {
		mid := middle(min, max)
		// check how many hours it takes with our k
		hNeeded := 0
		for _, pile := range piles {
			// need to round up since Koko needs extra hour to clear remaining bananas
			hNeeded += (pile + mid - 1) / mid
		}
		if hNeeded > h {
			// search right since our k is too slow
			min = mid + 1
		} else {
			// search left since our k is too big
			max = mid
		}
	}
	return min // min itself is just fast enough since we looking for smallest
}

func middle(a, b int) int {
	return (a + b) / 2
}
