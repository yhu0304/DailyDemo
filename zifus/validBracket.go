package zifus

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

思路：
借助栈，遇到左括号则入栈，遇到右括号需要与栈中最后一个元素匹配，匹配上则出栈
匹配可借助map实现
*/

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s)%2 != 0 {
		return false
	}
	var stack []uint8

	m := map[uint8]uint8{
		'}': '{',
		')': '(',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if m[s[i]] != stack[len(stack)-1] { // 左右括号未匹配，此时已是右括号，寻找对应的左括号
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
