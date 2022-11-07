package main

import "fmt"

// main function
func main() {
	fmt.Print("Hello, world.\n")

	// create new queue
	q := NewQueue()

	// add items to queue
	q.Enqueue(1)

	// check if queue is empty
	fmt.Printf("Is queue empty? %v\n", q.IsEmpty())

	// check if queue is full
	fmt.Printf("Is queue full? %v\n", q.IsFull())

	// get size of queue
	fmt.Printf("Size of queue: %v\n", q.Size())

	// get front item of queue
	fmt.Printf("Front item of queue: %v\n", q.Front())

	// get rear item of queue
	fmt.Printf("Rear item of queue: %v\n", q.Rear())

	// remove item from queue
	fmt.Printf("Removed item from queue: %v\n", q.Dequeue())

	// check if queue is empty
	fmt.Printf("Is queue empty? %v\n", q.IsEmpty())

	// check if queue is full
	fmt.Printf("Is queue full? %v\n", q.IsFull())

	// get size of queue
	fmt.Printf("Size of queue: %v\n", q.Size())

	// get front item of queue
	fmt.Printf("Front item of queue: %v\n", q.Front())

	// get rear item of queue
	fmt.Printf("Rear item of queue: %v\n", q.Rear())

	// remove item from queue
	fmt.Printf("Removed item from queue: %v\n", q.Dequeue())

	// check if queue is empty
	fmt.Printf("Is queue empty? %v\n", q.IsEmpty())

	// check if queue is full
	fmt.Printf("Is queue full? %v\n", q.IsFull())

	// get size of queue
	fmt.Printf("Size of queue: %v\n", q.Size())

	// get front item of queue
	fmt.Printf("Front item of queue: %v\n", q.Front())
}
