/*
 * @Name: Queue Go
 * @Author: Max Base
 * @Date: 2022-11-07
 * @Class: Data Structure, Dr. Mahsa Soheil Shamaee
 * @Repository: https://github.com/basemax/QueueGo
 */
package main

// type for queue
type Queue struct {
	rear  int
	front int
	items []int
}

// create new queue
func NewQueue() *Queue {
	return &Queue{rear: 0, front: 0, items: []int{}}
}

// add item to queue
// O(1)
func (q *Queue) Enqueue(i int) {
	if q.IsFull() {
		return
	}
	q.items = append(q.items, i)
	q.rear++
}

// remove item from queue
// O(1)
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}
	item := q.items[q.front]
	q.front++
	return item
}

// check if queue is empty
// O(1)
func (q *Queue) IsEmpty() bool {
	return q.rear == q.front
}

// check if queue is full
// O(1)
func (q *Queue) IsFull() bool {
	return q.rear == len(q.items)
}

// get size of queue
// O(1)
func (q *Queue) Size() int {
	return q.rear - q.front
}

// get front item of queue
// O(1)
func (q *Queue) Front() int {
	if q.rear == q.front {
		return -1
	}
	return q.items[q.front]
}

// get rear item of queue
// O(1)
func (q *Queue) Rear() int {
	if q.rear == q.front {
		return -1
	}
	return q.items[q.rear-1]
}
