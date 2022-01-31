package algorithm

import (
	"sort"
	"strings"
)

/*
请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。
保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。

思路
	- strings.Count的使用
*/
func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

/*
请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

思路
	- 将string转成[]rune
	- 以中心字符为轴，交换两边的字符
*/
func ReversedString(s string) string {
	if len(s) > 5000 {
		return ""
	}

	strList := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		strList[i], strList[len(s)-1-i] = strList[len(s)-1-i], strList[i]
	}

	return string(strList)
}

/*
给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，
代表两串是否重新排列后可相同。 保证两串的长度都小于等于5000。

思路
	1、strings.Count()判断s1中的某个字符，在s2中的数量是否一致
*/
func IsRegroup(s1, s2 string) bool {
	if len(s1) == 0 || len(s1) > 5000 {
		return false
	}
	if len(s2) == 0 || len(s2) > 5000 {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}

	for _, i := range s1 {
		if strings.Count(s1, string(i)) != strings.Count(s2, string(i)) {
			return false
		}
	}

	return true
}

/*
345. 反转字符串中的元音字母
https://leetcode-cn.com/problems/reverse-vowels-of-a-string/
思路：双指针
*/

var vowelMap = map[string]struct{}{
	"a": {},
	"e": {},
	"i": {},
	"o": {},
	"u": {},
}

func reverseVowels(s string) string {
	if len(s) <= 1 {
		return s
	}

	bStr := []byte(s)

	j := len(bStr) - 1
	for i := 0; i < len(s); i++ {
		_, ok := vowelMap[strings.ToLower(string(bStr[i]))]
		if ok {
			for j > i {
				_, ok2 := vowelMap[strings.ToLower(string(bStr[j]))]
				if ok2 {
					bStr[i], bStr[j] = bStr[j], bStr[i]
					j--
					break
				}
				j--
			}
		}

	}

	return string(bStr)
}

/*
387. 字符串中的第一个唯一字符
https://leetcode-cn.com/problems/first-unique-character-in-a-string/
*/
func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	sMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		sMap[string(s[i])] += 1
	}

	for j := 0; j < len(s); j++ {
		if num, ok := sMap[string(s[j])]; ok && num == 1 {
			return j
		}
	}

	return -1
}

/*
389. 找不同
https://leetcode-cn.com/problems/find-the-difference/
思路：
1、排序
2、双指针
*/
func findTheDifference(s string, t string) byte {
	if len(s)+1 != len(t) {
		return 0
	}

	sArray := []byte(s)
	tArray := []byte(t)
	sort.Slice(sArray, func(i, j int) bool {
		return sArray[j] > sArray[i]
	})
	sort.Slice(tArray, func(i, j int) bool {
		return tArray[j] > tArray[i]
	})

	for i := 0; i < len(sArray); i++ {
		if sArray[i] != tArray[i] {
			return tArray[i]
		}
	}

	return tArray[len(tArray)-1]
}

/*
392. 判断子序列
https://leetcode-cn.com/problems/is-subsequence/
*/
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	amount := 0
	j := 0
	for i := 0; i < len(s); i++ {
		for j < len(t) {
			if s[i] == t[j] {
				amount++
				j++
				break
			}
			j++
		}
	}

	return len(s) == amount
}

/*
409. 最长回文串
https://leetcode-cn.com/problems/longest-palindrome/
思路：
1、排序
2、求最大偶数字符串长度
*/
func longestPalindrome(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	sByte := []byte(s)
	sort.Slice(sByte, func(i, j int) bool {
		return sByte[i] < sByte[j]
	})

	i, sum, isAdd := 0, 0, false
	for i < len(sByte) {
		strNum := 1
		for j := i + 1; j < len(sByte); j++ {
			if sByte[i] == sByte[j] {
				strNum++
				if j == len(sByte)-1 {
					i = j
					break
				}
				continue
			}
			i = j
			break
		}

		if strNum%2 == 0 {
			sum += strNum
		} else {
			sum += strNum - 1
			isAdd = true
		}

		if i == len(sByte)-1 {
			if strNum == 1 {
				isAdd = true
			}
			i++
		}
	}

	if isAdd {
		sum++
	}

	return sum
}

/*
415. 字符串相加
https://leetcode-cn.com/problems/add-strings/
*/
func addStrings(num1 string, num2 string) string {
	max := max3(len(num1), len(num2))
	isN := false
	result := ""
	for i := 1; i <= max; i++ {
		var num1Item byte
		var num2Item byte

		if len(num1)-i >= 0 {
			num1Item = num1[len(num1)-i] - 48
		}
		if len(num2)-i >= 0 {
			num2Item = num2[len(num2)-i] - 48
		}
		sum := num1Item + num2Item
		if isN {
			sum++
		}

		isN = sum/10 > 0
		result = string((sum)%10+48) + result
	}
	if isN {
		result = "1" + result
	}

	return result
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
459. 重复的子字符串
https://leetcode-cn.com/problems/repeated-substring-pattern/
思路：枚举
1、len(s)为偶数
2、枚举子字符串长度n2，1 < n2 <= len(s)
3、s[i] == s[i + n2]
*/
func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}

	n := len(s)
	for n2 := 1; n2 <= n/2; n2++ {
		if n%n2 == 0 && s[0] == s[n2] { // 子长度需为n的倍数
			match := true
			for i := 0; i+n2 < n; i++ {
				if s[i] != s[i+n2] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}

	}

	return false
}

/*
482. 密钥格式化
https://leetcode-cn.com/problems/license-key-formatting/
*/
func licenseKeyFormatting(s string, k int) string {
	if len(s) == 0 || k == 0 {
		return ""
	}

	str := ""
	for _, i := range s {
		if string(i) != "-" {
			str += string(i)
		}
	}
	l := len(str)
	if l == 0 {
		return ""
	}
	str = strings.ToUpper(str)

	firstLen := l % k
	if l%k == 0 {
		firstLen = k
	}
	result := str[:firstLen]
	strItem := ""
	for i := firstLen; i < l; i++ {
		strItem += string(str[i])
		if len(strItem) == k {
			result = result + "-" + strItem
			strItem = ""
		}
	}
	return result
}

/*
541. 反转字符串 II
https://leetcode-cn.com/problems/reverse-string-ii/
*/
func reverseStr(s string, k int) string {
	t := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		sub := t[i:min2(i+k, len(s))]
		for j, n := 0, len(sub); j < n/2; j++ {
			sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
		}
	}
	return string(t)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
