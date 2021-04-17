package algorithm

import (
	"fmt"
	"sync"
	"testing"

	"github.com/bmizerany/assert"
	"time"
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

func TestCountAndSay(t *testing.T)  {
	fmt.Println(countAndSay(4))
}

func TestMaxSubArray(t *testing.T)  {
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