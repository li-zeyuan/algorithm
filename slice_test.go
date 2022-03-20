package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversedSlice(t *testing.T) {
	strList := []string{"A", "B", "C", "D"}
	sList := ReversedSlice(strList)
	fmt.Println(sList)
}

func TestReverseString(t *testing.T) {
	var l1 []byte
	reverseString(l1)
	assert.Equal(t, l1, []byte{})

	l2 := []byte{96, 97, 98}
	reverseString(l1)
	assert.Equal(t, l2, []byte{98, 97, 96})

	l3 := []byte{96, 97, 98, 99}
	reverseString(l1)
	assert.Equal(t, l3, []byte{99, 98, 97, 96})
}

func TestIntersection(t *testing.T) {
	l1 := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	assert.Equal(t, l1[0], 2)
}

func TestIntersect(t *testing.T) {
	l1 := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	assert.Equal(t, len(l1), 2)
	assert.Equal(t, l1[0], 2)
	assert.Equal(t, l1[1], 2)

	l2 := intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	assert.Equal(t, len(l2), 2)
	assert.Equal(t, l2[0], 4)
	assert.Equal(t, l2[1], 9)
}

func TestThirdMax(t *testing.T) {
	assert.Equal(t, thirdMax([]int{3, 2, 1}), 1)
	assert.Equal(t, thirdMax([]int{1, 2}), 2)
	assert.Equal(t, thirdMax([]int{2, 2, 3, 1}), 1)
}

func TestFindDisappearedNumbers(t *testing.T) {
	assert.Equal(t, len(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})), 2)
	assert.Equal(t, findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})[0], 5)
	assert.Equal(t, findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})[1], 6)

	assert.Equal(t, len(findDisappearedNumbers([]int{1, 1})), 1)
	assert.Equal(t, minMoves([]int{1, 1}), 2)
}

func TestFindDisappearedNumbers2(t *testing.T) {
	assert.Equal(t, len(findDisappearedNumbers2([]int{1, 2, 1})), 2)
}

func TestFindContentChildren(t *testing.T) {
	assert.Equal(t, findContentChildren([]int{1, 2, 3}, []int{1, 1}), 1)
	assert.Equal(t, findContentChildren([]int{1, 2}, []int{1, 2, 3}), 2)
	assert.Equal(t, findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8}), 2)
}

func TestMaxArea(t *testing.T) {
	assert.Equal(t, maxArea([]int{1, 2}), 1)
	assert.Equal(t, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
}

func TestMaxArea2(t *testing.T) {
	assert.Equal(t, maxArea2([]int{1, 2}), 1)
	assert.Equal(t, maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
}

func TestLetterCombinations(t *testing.T) {
	t.Log(t, letterCombinations("7"))
	t.Log(t, letterCombinations("23"))
	t.Log(t, letterCombinations("233"))
}

func TestGroupAnagrams(t *testing.T) {
	t.Log(t, groupAnagrams([]string{""}))
	t.Log(t, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func TestSpiralOrder(t *testing.T) {
	t.Log(t, spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
}

func TestCanJump(t *testing.T)  {
	assert.Equal(t, canJump([]int{0,2,3}), false)
	assert.Equal(t, canJump([]int{2,3,1,1,4}), true)
	assert.Equal(t, canJump([]int{3,2,1,0,4}), false)
}