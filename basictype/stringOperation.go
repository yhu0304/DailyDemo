package basictype

import (
	"fmt"
	"strings"
)

func Stringtest() {
	s1 := " aBc"
	s2 := "100a"
	s3 := s1 + s2
	s4 := s3[2:4] //字符串截取
	fmt.Println(s4)
	fmt.Println(strings.HasPrefix(s3, "a"))               //判断前缀
	fmt.Println(strings.HasSuffix(s3, "0"))               //判断后缀
	fmt.Println(strings.Contains(s3, "9"))                //字符串包含关系
	fmt.Println(strings.Index(s3, "0"))                   //判断子字符串或字符在父字符串中第一次出现的位置（索引）
	fmt.Println(strings.LastIndex(s3, "0"))               //最后出现位置的索引
	fmt.Println(strings.Replace(s3, "0", "1", -1))        //如果 n = -1 则替换所有字符串
	fmt.Println(strings.Count(s3, "0"))                   //出现的非重叠次数
	fmt.Println(strings.Repeat(s3, 2))                    //重复字符串
	fmt.Println(strings.ToLower(s3))                      //修改字符串大小写
	fmt.Println(strings.ToUpper(s3))                      //修改字符串大小写
	fmt.Println(strings.TrimSpace(s3))                    //修剪字符串 去掉开头和结尾空格
	fmt.Println(strings.Trim(strings.TrimSpace(s3), "a")) //修剪字符串 去掉开头和结尾字符串
	fmt.Println(strings.Split(s3, "a"))                   //以"a"为分隔符，将s切分成多个子切片
}
