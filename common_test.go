package algorithm

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestGetSingleton(t *testing.T) {
	s1 := GetSingleton()
	s2 := GetSingleton()

	if s1 != s2 {
		t.Fatal("not equal")
	}

	assert.Equal(t, s1, s2)
}

func TestDecorator(t *testing.T) {
	coolFunc(myFunc)
}

func TestGetIota(t *testing.T) {
	GetIota()
}

func TestFac(t *testing.T) {
	fmt.Println(Sum(4))
}

func TestPriFib(t *testing.T) {
	PriFib(5)
}

func TestChanel(t *testing.T) {
	ch := make(chan int)

	close(ch)
	ch <- 1

	fmt.Println("iii")
}

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(123121))
}

func TestLargestNumber(t *testing.T) {
	str := largestNumber([]int{3, 30, 34, 5, 9})
	fmt.Println(str)
}

func TestRW(t *testing.T) {
	rw := sync.RWMutex{}
	wg := sync.WaitGroup{}

	rw.Lock()

	wg.Add(1)
	go func() {
		defer rw.RUnlock()
		defer wg.Done()
		fmt.Println("111")
		rw.RLock()
		fmt.Println("222")
	}()

	time.Sleep(time.Second)
	rw.Unlock()
	wg.Wait()

	fmt.Println()
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Println(result)
}

func TestSearchInsert(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(searchInsert(l, 5))
}

func TestCountAndSay(t *testing.T) {
	fmt.Println(countAndSay(4))
}

func TestMaxSubArray(t *testing.T) {
	l := []int{
		1,
		1,
		-3,
		1,
		1,
		1,
	}
	fmt.Println(maxSubArray(l))
}

func TestLengthOfLastWord(t *testing.T) {
	//assert.Equal(t, lengthOfLastWord("li ze yuan"), 4)
	assert.Equal(t, lengthOfLastWord("li ze yuan "), 4)
}

func TestPlusOne(t *testing.T) {
	l := []int{
		9,
		9,
		9,
	}
	fmt.Println(plusOne(l))
}

func TestAddBinary(t *testing.T) {
	fmt.Println(addBinary("1010", "1011"))
}

func TestMySqrt(t *testing.T) {
	fmt.Println(mySqrt(1))
}

func TestClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(7))
}

func TestClimbStairs2(t *testing.T) {
	fmt.Println(climbStairs2(5))
}

func TestDeleteDuplicates(t *testing.T) {
	h := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				2,
				nil,
			},
		},
	}
	fmt.Println(deleteDuplicates(h))
}

func TestMerge2(t *testing.T) {
	merge2([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func TestGenerate(t *testing.T) {
	t.Log(generate(5))
}

func TestGetRow(t *testing.T) {
	t.Log(getRow(3))
}

func TestMaxProfit(t *testing.T) {
	t.Log(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func TestMaxProfit2(t *testing.T) {
	t.Log(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}

func TestIsPalindrome2(t *testing.T) {
	t.Log(isPalindrome2("A man, a plan, a canal: Panama"))
}

func TestSingleNumber(t *testing.T) {
	t.Log(singleNumber([]int{2, 2, 1}))
}

func TestConvertToTitle(t *testing.T) {
	t.Log(convertToTitle(701))
}

func TestMajorityElement(t *testing.T) {
	t.Log(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func TestMajorityElementV2(t *testing.T) {
	t.Log(majorityElementV2([]int{2, 2, 1, 1, 1, 2, 2}))
}

func TestTitleToNumber(t *testing.T) {
	t.Log(titleToNumber("AB"))
}

func TestTrailingZeroes(t *testing.T) {
	t.Log(trailingZeroes(30))
}

func TestIsHappy(t *testing.T) {
	t.Log(isHappy(2))
}

func TestIsIsomorphic(t *testing.T) {
	t.Log(isIsomorphic("badc", "baba"))
}

func TestContainsDuplicate(t *testing.T) {
	t.Log(containsDuplicate([]int{1, 2, 3, 1}))
}

func TestContainsNearbyDuplicate(t *testing.T) {
	t.Log(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}

func TestSummaryRanges(t *testing.T) {
	t.Log(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func TestSummaryRanges2(t *testing.T) {
	t.Log(summaryRanges2([]int{0, 1, 2, 4, 5, 7}))
}

func TestIsPowerOfTwo(t *testing.T) {
	t.Log(isPowerOfTwo(2))
	fmt.Printf("==%b==", 1>>1)
}

func TestIsAnagram(t *testing.T) {
	assert.Equal(t, isAnagram("rat", "car"), false)
}

func TestIsAnagram2(t *testing.T) {
	assert.Equal(t, isAnagram2("rat", "car"), false)
}

func TestIsUgly(t *testing.T) {
	assert.T(t, isUgly(1), true)
}

func TestMissingNumber(t *testing.T) {
	assert.Equal(t, missingNumber([]int{1, 0}), 2)
}

func TestFirstBadVersion(t *testing.T) {
	assert.Equal(t, firstBadVersion(5), 3)
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}

func TestIsPowerOfFour(t *testing.T) {
	assert.Equal(t, isPowerOfFour(17), false)
	assert.Equal(t, isPowerOfFour(0), false)
	assert.Equal(t, isPowerOfFour(1), true)
	assert.Equal(t, isPowerOfFour(4), true)
	assert.Equal(t, isPowerOfFour(5), false)
	assert.Equal(t, isPowerOfFour(8), false)
	assert.Equal(t, isPowerOfFour(16), true)
}

func TestWordPattern(t *testing.T) {
	assert.Equal(t, wordPattern("abba", "dog cat cat dog"), true)
}

func TestConstructor3(t *testing.T) {
	obj := Constructor3([]int{-2, 0, 3, -5, 2, -1})
	assert.Equal(t, obj.SumRange(0, 2), 1)
	assert.Equal(t, obj.SumRange(2, 5), -1)
	assert.Equal(t, obj.SumRange(0, 5), -3)
}

func TestIsPowerOfThree(t *testing.T) {
	assert.Equal(t, isPowerOfThree(0), false)
	assert.Equal(t, isPowerOfThree(27), true)
	assert.Equal(t, isPowerOfThree(4), false)
}

func TestCountBits(t *testing.T) {
	assert.Equal(t, countBits(0), []int{0})
	assert.Equal(t, countBits(2), []int{0, 1, 1})
	assert.Equal(t, countBits(5), []int{0, 1, 1, 2, 1, 2})
}

func TestCountOneBit(t *testing.T) {
	assert.Equal(t, countOneBit(0), 0)
	assert.Equal(t, countOneBit(2), 1)
	assert.Equal(t, countOneBit(5), 2)
}

func TestIsPerfectSquare(t *testing.T) {
	assert.Equal(t, isPerfectSquare(1), true)
	assert.Equal(t, isPerfectSquare(16), true)
	assert.Equal(t, isPerfectSquare(14), false)
}

func TestGuessNumber(t *testing.T) {
	assert.Equal(t, guessNumber(2), 2)
}

func TestCanConstruct(t *testing.T) {
	assert.Equal(t, canConstruct("a", "b"), false)
	assert.Equal(t, canConstruct("aa", "ab"), false)
	assert.Equal(t, canConstruct("aa", "aab"), true)
	assert.Equal(t, canConstruct("aab", "baa"), true)
}

func TestMerge(t *testing.T) {
	merge([]int{1}, 2, []int{3}, 3)
}

func TestCanWinNim(t *testing.T) {
	_ = canWinNim(3)
}

func TestFizzBuzz(t *testing.T)  {
	assert.Equal(t, len(fizzBuzz(3)), 3)
	assert.Equal(t, fizzBuzz(3)[0], "1")
	assert.Equal(t, fizzBuzz(3)[1], "2")
	assert.Equal(t, fizzBuzz(3)[2], "Fizz")

	assert.Equal(t, len(fizzBuzz(5)), 5)
	assert.Equal(t, fizzBuzz(5)[0], "1")
	assert.Equal(t, fizzBuzz(5)[1], "2")
	assert.Equal(t, fizzBuzz(5)[2], "Fizz")
	assert.Equal(t, fizzBuzz(5)[3], "4")
	assert.Equal(t, fizzBuzz(5)[4], "Buzz")
}

func TestCountSegments(t *testing.T)  {
	assert.Equal(t, countSegments("Hello, my name is John"), 5)
	assert.Equal(t, countSegments(""), 0)
	assert.Equal(t, countSegments(" Hello, "), 1)
}

func TestArrangeCoins(t *testing.T)  {
	assert.Equal(t, arrangeCoins(5), 2)
	assert.Equal(t, arrangeCoins(8), 3)
}

func TestFindMaxConsecutiveOnes(t *testing.T)  {
	assert.Equal(t, findMaxConsecutiveOnes([]int{1,1,0,1,1,1}), 3)
	assert.Equal(t, findMaxConsecutiveOnes([]int{}), 0)
	assert.Equal(t, findMaxConsecutiveOnes([]int{0}), 0)
	assert.Equal(t, findMaxConsecutiveOnes([]int{1}), 1)
}

func TestConstructRectangle(t *testing.T)  {
	assert.Equal(t, constructRectangle(1)[0], 1)
	assert.Equal(t, constructRectangle(1)[1], 1)
	assert.Equal(t, constructRectangle(4)[0], 2)
	assert.Equal(t, constructRectangle(4)[1], 2)
	assert.Equal(t, constructRectangle(6)[0], 3)
	assert.Equal(t, constructRectangle(6)[1], 2)
	assert.Equal(t, constructRectangle(122122)[0], 427)
	assert.Equal(t, constructRectangle(122122)[1], 286)
}

func TestFindPoisonedDuration(t *testing.T)  {
	assert.Equal(t, findPoisonedDuration([]int{1,4}, 2), 4)
	assert.Equal(t, findPoisonedDuration([]int{1,2}, 2), 3)
	assert.Equal(t, findPoisonedDuration([]int{1,2,3,4,5}, 5), 9)
}
func TestNextGreaterElement(t *testing.T)  {
	t.Log(nextGreaterElement([]int{4,1,2}, []int{1,3,4,2}))
	t.Log(nextGreaterElement([]int{2,4}, []int{1,2,3,4}))
}