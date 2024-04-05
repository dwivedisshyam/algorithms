package sort

// Insertion sort virtually splits the array into a sorted and an unsorted part.
// Values from the unsorted part are picked and placed in the correct
// position in the sorted part.
func Insertion(arr []int, order Order) []int {
	n := len(arr)
	comp := comparer(order)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for ; j >= 0 && comp(arr[j], key); j-- {
			arr[j+1] = arr[j]
		}

		arr[j+1] = key
	}

	return arr
}

// 11,12,13,5
