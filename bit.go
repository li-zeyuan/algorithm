package algorithm

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