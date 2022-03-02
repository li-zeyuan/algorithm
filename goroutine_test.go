package algorithm

import (
	"fmt"
	"testing"
)

func TestAlternatePrint(t *testing.T) {
	AlternatePrint(26)
}

func TestTickerPanic(t *testing.T) {
	TickerPanic()
}

func TestWaitTimeoutDo(t *testing.T) {
	WaitTimeoutDo()
}

func TestAlternatePrintNum(t *testing.T) {
	AlternatePrintNum(10)
}

func TestAlternatePrintNum2(t *testing.T) {
	AlternatePrintNum2(10)
}
func TestAlternatePrintNum3(t *testing.T) {
	AlternatePrintNum3(10)
}
func TestAlternatePrintNum4(t *testing.T) {
	AlternatePrintNum4(10)
}

func TestDefer(t *testing.T)  {
	fmt.Println(f1())
}

func f1() int { //匿名返回值
	var r int = 6
	defer func() {
		r *= 7
	}()
	return r
}