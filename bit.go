package algorithm

import (
	"fmt"
)

/*
<< 右移：i<<N = i * 2^N
>> 左移：i>>N = i / 2^N
& AND：同一位上都是1，才是1；否则为0
| OR：同一位上有一个是1就1,只有都是0才0
^ 异或：二进制位不同时,当前位不同才置1 否则置0（任何数和本身异或 结果为0；0和任意数异或 结果为其本身）
*/

type A struct {
	id   int
	name string
}

func main1() {
	a := A{id: 5, name: "123"}
	b := A{id: 5, name: "123"}
	c := A{id: 5, name: "1234"}
	fmt.Println(a == b) // true
	fmt.Println(a == c) // false
}

/*
翻转二进制位
https://leetcode-cn.com/problems/reverse-bits/
*/
func reverseBits(n uint32) (rev uint32) {
	for i := 0; i < 32 && n > 0; i++ {
		rev |= n & 1 << (31 - i)
		n >>= 1
	}
	return
}

/*
位1的个数
https://leetcode-cn.com/problems/number-of-1-bits/
方式一：逐位比对
*/
func hammingWeight(num uint32) int {
	ones := 0
	for i := 0; i < 32; i++ { //uint32 有32位
		if 1<<i&num > 0 { // 1右移i位，1<<i依次从num低位开始 &运算
			ones++
		}
	}
	return ones
}

/*
461. 汉明距离
https://leetcode-cn.com/problems/hamming-distance/
思路：
1、x ^ y；^:异或，二进制位相同时为0，不相同时为1
2、求异或后1的个数
*/
func hammingDistance(x int, y int) int {
	temp := x ^ y
	ones := 0
	for i := 0; i < 64; i++ { // 有32位
		if 1<<i&temp > 0 { // 1右移i位，1<<i依次从num低位开始 &运算
			ones++
		}
	}
	return ones
}

/*
476. 数字的补数
https://leetcode-cn.com/problems/number-complement/
思路：
1、找出1最高位
2、求num的掩码
3、num ^ 掩码
*/
func findComplement(num int) int {
	highBit := 0
	for i := 30; i >= 0; i-- { // 2^i ≤num<2^i+1
		if num&(1<<i) > 0 {
			highBit = i
			break
		}
	}

	mask := 1<<(highBit+1) - 1 // 掩码mask = 2^i+1 - 1
	return num ^ mask
}

func findComplement2(num int) {
	fmt.Println(num ^ 1)
}
