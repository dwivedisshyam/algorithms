package sort

// Selection picks up an element from the unsorted portion of the list and swaps
// it with the first element of the unsorted part. This process is repeated for
// the remaining unsorted portion until the entire list is sorted.
func Selection(arr []int, order Order) []int {
	n := len(arr)
	comp := comparer(order)

	for i := 0; i <= n-1; i++ {
		for j := i + 1; j < n; j++ {
			if comp(arr[i], arr[j]) {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	return arr
}
