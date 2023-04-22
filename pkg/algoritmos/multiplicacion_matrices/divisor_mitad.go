package multiplicacion_matrices

// This function returns the total number of divisors for an integer
// Based on https://codereview.stackexchange.com/questions/120642/getting-all-divisors-from-an-integer
func getDivisorsCnt(n int) int {
	var divisors int
	mod := n
	for mod > 0 {
		if n%mod == 0 {
			divisors++
		}
		mod--
	}
	return divisors
}

// This function returns the middle divisor between all the divisors of a number
// It assumes that the number has an odd number of divisors
func getMiddleDivisor(n int) int {
	var count int
	target := (getDivisorsCnt(n) + 1) / 2 // The target position of the middle divisor
	mod := n
	for mod > 0 {
		if n%mod == 0 {
			count++
			if count == target { // Found the middle divisor
				return mod
			}
		}
		mod--
	}
	return -1 // Should not happen
}
