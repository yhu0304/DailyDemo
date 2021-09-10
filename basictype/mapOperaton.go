package basictype

import "fmt"

func DeleteMap(m map[int]string, key int) {
	delete(m, key)
	for k, v := range m {
		fmt.Printf("len(m)=%d, %d ----> %s\n", len(m), k, v)
	}
}

func Maptest() {
	//赋值-->修改和追加
	var m1 = map[int]string{1: "ck_god", 2: "god_girl"}
	m1[1] = "xxx"
	m1[3] = "lily" //追加，go底层自动为map分配空间
	fmt.Println(m1)

	m2 := make(map[int]string, 10)
	m2[0] = "aaa"
	m2[1] = "bbb"
	fmt.Println(m2)
	fmt.Println(m2[0], m2[1])
	fmt.Println("")

	//遍历  forr
	m3 := map[int]string{1: "ck_god", 2: "god_girl"}
	for k, v := range m3 {
		fmt.Printf("%d--->>>%s\n", k, v)
	}

	//遍历获得 key
	for k := range m3 {
		fmt.Printf("%d--->%s\n", k, m3[k])
	}
	fmt.Println("")

	//获取map的值，ok表示是否成功
	value, ok := m3[1]
	fmt.Println("value = ", value, ", ok = ", ok)

	value2, ok2 := m3[3]
	fmt.Println("value2 = ", value2, ", ok2 = ", ok2)
	fmt.Println("")

	//删除
	m4 := map[int]string{1: "ck_god", 2: "god_girl"}
	delete(m4, 2) //删除key值为2的map
	for k, v := range m4 {
		fmt.Printf("%d--->%s\n", k, v)
	}
	fmt.Println("")

	//map作为参数传递
	//在函数间传递映射并不会制造出该映射的一个副本，不是值传递，而是引用传递：
	m := map[int]string{1: "ck_god", 2: "god_girl", 3: "girl_angle"}

	DeleteMap(m, 3)
	for k, v := range m {
		fmt.Printf("len(m)=%d, %d ----> %s\n", len(m), k, v)
	}
}
