package source

func MySqrt(x int) int {
	delta := 0.1
	high := float64(x)
	low := float64(0)
	for {
		mid := (low+high)/2
		if mid*mid-float64(x) > delta {
			high = mid
		} else if mid*mid-float64(x) < 0 {
			low = mid
		} else {
			break
		}
	}

	return int(low+high)/2
}
