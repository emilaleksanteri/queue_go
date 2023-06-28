package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	queue := new(Queue)
	queue.enque(1)
	queue.enque(10)
	queue.enque(8)
	queue.enque(5)

	if queue.head.value != 1 {
		t.Error("Queue head should have been 1 but it was ", queue.head.value)
	}

	if queue.tail.value != 5 {
		t.Error("Queue tail should have been 5 but it was ", queue.tail.value)
	}

	queue.deque()
	if queue.head.value != 10 {
		t.Error("Queue head should have been 1 but it was ", queue.head.value)
	}

	queue.deque()
	if queue.head.value != 8 {
		t.Error("Queue head should have been 1 but it was ", queue.head.value)
	}

	queue.deque()
	if queue.head.value != 5 {
		t.Error("Queue head should have been 1 but it was ", queue.head.value)
	}

	queue.deque()
	empty := queue.deque()
	if empty != 0 && queue.length == 0{
		t.Error("empty queue should have returned 0 and have length 0")
	}




}
