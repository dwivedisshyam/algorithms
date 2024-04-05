package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertion(t *testing.T) {
	tcs := testCases()

	for _, tc := range tcs {
		actual := Insertion(tc.input, tc.order)

		assert.Equal(t, tc.expected, actual)
	}
}
