package algorithm

import (
	"testing"
)

func TestConstructor2(t *testing.T) {
	myQueue := Constructor2()
	myQueue.Push(1) // queue is: [1]
	myQueue.Push(2) // queue is: [1, 2] (leftmost is front of the queue)
	myQueue.Push(3) // queue is: [1, 2] (leftmost is front of the queue)
	myQueue.Push(4) // queue is: [1, 2] (leftmost is front of the queue)
	myQueue.Pop()   // return 1, queue is [2]
	myQueue.Push(5) // queue is: [1, 2] (leftmost is front of the queue)
	myQueue.Pop()
	myQueue.Pop()
	myQueue.Pop()
	myQueue.Pop()
}
