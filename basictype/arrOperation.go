package basictype

import (
	"fmt"
)

func Arrtest() {
	//一维数组初始化
	var a [4]int           //元素自动初始化为0
	b := [4]int{2, 5}      //[2 5 0 0]
	c := [...]int{1, 2, 3} //根据初始化值数量确定长度
	d := [2]string{"Tiger", "Peggy"}
	e := [...]int{10, 3: 100} //支持索引初始化，数组的长度与此有关 [10 0 0 100]
	fmt.Println(a, b, c, d, e)
	arr := []int{0, 1, 2, 3, 4}
	//删除第i个元素
	i := 2
	arr = append(arr[:i], arr[i+1:]...)

}

//通过索引删除
func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func ArrayRemoveRepeated(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
