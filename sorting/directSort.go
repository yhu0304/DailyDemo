package sorting

func InsertSort(arr []int) []int{
	var j int
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		j = i - 1
		for j = i - 1; j >= 0 && arr[j] > temp; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
	return arr
}
