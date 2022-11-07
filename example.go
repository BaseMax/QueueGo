package main

import "fmt"

// main function
func main() {
	fmt.Print("Hello, world.\n")

	// create new queue
	q := Queue{}

	// add items to queue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	// print front item
	println(q.Front())

	// print rear item
	println(q.Rear())

	// print size of queue
	println(q.Size())

	// print if queue is empty
	println(q.IsEmpty())

	// print if queue is full
	println(q.IsFull())

	// remove items from queue
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())

	// print if queue is empty
	println(q.IsEmpty())
}
