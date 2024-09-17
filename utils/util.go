package utils

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func Power(base uint32, pow uint32) uint32 {
	var result uint32 = 1
	for pow > 0 {
		if pow%2 == 0 {
			pow /= 2
			base *= base
		} else {
			pow -= 1
			result *= base
			pow /= 2
			base *= base
		}
	}
	return result
}
