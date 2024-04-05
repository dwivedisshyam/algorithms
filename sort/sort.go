package sort

type Order string
type compare func(int, int) bool

const (
	ASC  Order = Order("ASC")
	DESC Order = Order("DESC")
)

func asc(a, b int) bool {
	return a > b
}

func desc(a, b int) bool {
	return a < b
}

func comparer(order Order) compare {
	var comp compare = asc
	if order == DESC {
		comp = desc
	}

	return comp
}
