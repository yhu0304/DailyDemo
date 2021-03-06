package normal

/*
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。


示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0
*/

func reverse(x int) int {
	res := 0
	for x != 0 {
		tmp := x % 10
		res = res*10 + tmp
		x /= 10
	}
	if int(int32(res)) != res {
		return 0
	}
	return res
}

/*
 思路：
1、将12345 % 10 得到5，之后将12345 / 10
2、将1234 % 10 得到4，再将1234 / 10
3、将123 % 10 得到3，再将123 / 10
4、将12 % 10 得到2，再将12 / 10
5、将1 % 10 得到1，再将1 / 10
需判断溢出，转成int32再转回来
*/
