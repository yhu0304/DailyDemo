package sorting

import "fmt"

func SelectionSort(a []int) {
	var n int = len(a)
	for i := 0; i < n; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[k] {
				k = j
			}

		}
		if k > i {
			a[i], a[k] = a[k], a[i]

		}
	}
	fmt.Println(a)
}
