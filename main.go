package main

import (
	"fmt"
)

func main() {
	queue := new(Queue)
	fmt.Println(queue)

	queue.enque(1)
	queue.enque(10)
	queue.enque(8)
	queue.enque(5)
	fmt.Println("Queue should be 1, 10, 8, 5 now")

	fmt.Println(queue.deque())
	fmt.Println("head:", queue.peak(), "len:", queue.length)
	fmt.Println("Queue head should be: 10, length: 3")

	fmt.Println(queue.deque())
	fmt.Println("head", queue.peak(), "len:", queue.length)
	fmt.Println("Queue head should be: 8, length: 2")
}


type Node struct {
	value uint16
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	length uint16
}

func (q *Queue) peak() uint16 {
	return q.head.value
}

func (q *Queue) enque(value uint16) {
	q.length += 1
	newTail := Node{ value: value }
	if (q.length == 1) {
		q.head = &newTail
		q.tail = *&q.head
		q.head.next = q.tail
		return
	}

	if (q.length == 2) {
		q.head.next = q.tail
		q.tail.next = &newTail
		q.tail = q.tail.next
		return
	}

	q.tail.next = &newTail
	q.tail = &newTail
}

func (q *Queue) deque() uint16 {
	if q.length == 0 {
		return 0
	}

	q.length -= 1

	value := q.head.value
	if (q.length == 0) {
		q.head = &Node{}
		q.tail = q.head
		return value
	}
	fmt.Println(q.head.next, "is next from head")
	q.head = q.head.next
	return value
}

