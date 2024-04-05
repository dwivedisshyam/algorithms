package sort

type tc struct {
	input    []int
	order    Order
	expected []int
}

func testCases() []tc {
	tcs := []tc{
		{input: []int{1, 3, 4, 5, 2}, order: ASC, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{1, 3, 4, 5, 2}, order: DESC, expected: []int{5, 4, 3, 2, 1}},
	}

	return tcs
}
