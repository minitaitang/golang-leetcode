package leetcode

func corpFlightBooking(bookings [][]int, n int) []int {
	counter := make([]int, n+1)

	for _, b := range bookings {
		i, j, k := b[0], b[1], b[2]
		counter[i-1] += k
		if j < n {
			counter[j] -= k
		}
	}
	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i-1]
	}

	return counter[:len(counter)-1]
}
