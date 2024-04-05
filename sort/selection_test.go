package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelection(t *testing.T) {
	tcs := testCases()

	for _, tc := range tcs {
		actual := Selection(tc.input, tc.order)

		assert.Equal(t, tc.expected, actual)
	}
}
