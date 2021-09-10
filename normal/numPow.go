package normal

import "errors"

/*
> 给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。

前面所说的指数为负数只是边界的一种情况，学习算法，必须全面了解所有的边界

指数幂的所有边界包括:

- 指数为0的情况，不管底数是多少都应该是1
- 指数为负数的情况，求出的应该是其倒数幂的倒数
- 指数为负数的情况下，底数不能为0

正整数幂次可以用位运算来求
*/
func Pow(base float64, exp int) (float64, error) {
	if 0 == exp {
		return 1, nil
	}
	if exp < 0 && base == 0 {
		return -1, errors.New("base == 0 and exp < 0")
	}
	if exp > 0 {
		return PowNoraml(base, exp), nil
	} else { // 如果指数为负数，则应该求result的倒数
		res := PowNoraml(base, -exp)
		res = 1 / res
		return res, nil
	}
}

//递归方式，o(logn)
//当n为偶数，a^n =（a^n/2）*（a^n/2）
//当n为奇数，a^n = a^[(n-1)/2] * a^[(n-1)/2] * a
func PowNoraml(base float64, exp int) float64 {
	res, temp := 1.0, base
	for exp != 0 {
		// 如果指数n为奇数，则要再乘一次底数base
		// 奇数，二进制最后一位必定是1，与1相与得1；偶数与1相与得0
		if exp&1 == 1 { //判断奇数
			res *= temp
		}
		temp *= temp
		exp >>= 1
	}
	return res
}
