package zifus

import "fmt"

//直接操作string，使用一个额外的string
func StatnumString1(s string) map[string]int {
	var word string
	m := make(map[string]int)
	for i := 0; i < len(s); {
		word = s[i : i+1] //截取字符串
		v, ok := m[word]
		if ok != false {
			m[word] = v + 1
		} else {
			m[word] = 1
		}
		i += 1
	}
	for key, value := range m {
		fmt.Printf("%s = %d\n", key, value)
	}
	return m
}

//将string转换为字符数组，map中的键为byte，但是输出为ascii对应的int值
func StatnumString2(s string) {
	var data = []byte(s)
	m := make(map[byte]int)
	for _, value := range data {
		if _, ok := m[value]; ok { //以strs为例，s，m[s]没有值，赋值失败，设置值为1.第二次出现s时候，赋值1成功，设置值+1
			m[value] = m[value] + 1
		} else {
			m[value] = 1
		}
	}
	for key, value := range m {
		fmt.Printf("%s = %d\n", string(key), value)
	}
}
