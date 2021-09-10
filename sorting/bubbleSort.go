package sorting

func BubbleSort(source []int) []int{
	for i := len(source) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if source[j] > source[j+1] {
				source[j], source[j+1] = source[j+1], source[j]
			}
		}
	}
	return source
}