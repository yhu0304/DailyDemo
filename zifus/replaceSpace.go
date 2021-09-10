package zifus

func replaceSpace(str []byte) {
	var (
		count = 0
		lens  = len(str)
	)
	//遍历字符串，统计字符出现的数目，计算替换后的字符串长度
	for i := 0; i < lens; i++ {
		if str[i] == ' ' {
			count++
		}
	}
	newLen := lens + count*2

	//两个index，一个指向lens-1,另一个指向newLen-1，遍历一边字符串，完成替换
	for l, nl := lens-1, newLen-1; l >= 0 && nl >= 0; {
		if str[l] == ' ' {
			str[nl] = '0'
			nl--
			str[nl] = '2'
			nl--
			str[nl] = '%'
			nl--
			l--
		} else {
			str[nl] = str[l]
			nl--
			l--
		}
	}
}

/*
strings包提供了Replace()方法，Replace(old,new,-1)
*/
