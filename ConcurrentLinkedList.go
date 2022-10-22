package LinkedList

import "sync"

// ConcurrentLinkedList is a thread safety two-way loop linked list.
type ConcurrentLinkedList[T interface{}] struct {
	linkedList LinkedList[T]
	mutex      sync.RWMutex
}

// Start gets the value of the first element.
func (cl *ConcurrentLinkedList[T]) Start() T {
	cl.mutex.RLock()
	defer cl.mutex.RUnlock()
	return cl.linkedList.Start()
}

// End gets the value of the last element.
func (cl *ConcurrentLinkedList[T]) End() T {
	cl.mutex.RLock()
	defer cl.mutex.RUnlock()
	return cl.linkedList.End()
}

// Get the value of the current element.
func (cl *ConcurrentLinkedList[T]) Get() T {
	cl.mutex.RLock()
	defer cl.mutex.RUnlock()
	return cl.linkedList.Get()
}

// Set the value of the current element.
func (cl *ConcurrentLinkedList[T]) Set(v T) {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	cl.linkedList.Set(v)
}

// IsStart checks if it is the first element.
func (cl *ConcurrentLinkedList[T]) IsStart() bool {
	cl.mutex.RLock()
	defer cl.mutex.RUnlock()
	return cl.linkedList.IsStart()
}

// IsEnd checks if it is the last element.
func (cl *ConcurrentLinkedList[T]) IsEnd() bool {
	cl.mutex.RLock()
	defer cl.mutex.RUnlock()
	return cl.linkedList.IsEnd()
}

// Forward to the next element.
func (cl *ConcurrentLinkedList[T]) Forward() {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	cl.linkedList.Forward()
}

// Back to the last element.
func (cl *ConcurrentLinkedList[T]) Back() {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	cl.linkedList.Back()
}

// Length get the LinkedList's length.
func (cl *ConcurrentLinkedList[T]) Length() int {
	cl.mutex.RLock()
	defer cl.mutex.RUnlock()
	return cl.linkedList.Length()
}

// Index get the current index.
func (cl *ConcurrentLinkedList[T]) Index() int {
	cl.mutex.RLock()
	defer cl.mutex.RUnlock()
	return cl.linkedList.Index()
}

// Insert a value before this element
func (cl *ConcurrentLinkedList[T]) Insert(v T) {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	cl.linkedList.Insert(v)
}

// Append a value to the end of the LinkedList.
func (cl *ConcurrentLinkedList[T]) Append(v T) {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	cl.linkedList.Append(v)
}

// Delete the current value and back to the last value.
func (cl *ConcurrentLinkedList[T]) Delete() {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	cl.linkedList.Delete()
}

// Add the value to the item corresponding to position.
func (cl *ConcurrentLinkedList[T]) InsertTo(pos int, v T) {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	cl.linkedList.InsertTo(pos, v)
}

// Goto the item corresponding to the position(in var 'START','END','BEFORE' and 'AFTER').
func (cl *ConcurrentLinkedList[T]) Goto(pos int) {
	cl.linkedList.Goto(pos)
}

// Get the value of the item by selecting an position(in var 'START','END','BEFORE' and 'AFTER').
func (cl *ConcurrentLinkedList[T]) GetByPosition(pos int) T {
	cl.mutex.RLock()
	defer cl.mutex.RUnlock()
	return cl.linkedList.GetByPosition(pos)
}
