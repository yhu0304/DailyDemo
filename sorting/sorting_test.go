package sorting

import "testing"

func TestBubbleSort(t *testing.T) {
	source := []int{0,2,1,3,6,4,5}
	t.Log(BubbleSort(source))
}

func TestInsertSort(t *testing.T) {
	source := []int{0,2,1,3,6,4,5}
	t.Log(InsertSort(source))
}
