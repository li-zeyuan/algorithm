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

	go func() {
		wg.Add(1)
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
