package normal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	arr1 := []int{1, 1, 2, 2, 3, 3}
	arr2 := []int{0, 0, 1}
	assert.Equal(t, 3, removeDuplicates(arr1))
	assert.Equal(t, 2, removeDuplicates(arr2))
}
