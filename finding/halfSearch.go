package finding

func HalfSearch(arr []int, key int) (mid int) {
	var min,max int
	min = 0
	max = len(arr) - 1
	mid = (max + min) / 2
	for ;arr[mid] != key; {
		if key > arr[mid] {
			min = mid + 1
		} else if key < arr[mid] {
			max = mid - 1
		}
		if max < min {
			return -1
		}
		mid = (max + min) >> 1
	}
	return
}
