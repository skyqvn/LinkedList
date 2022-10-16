package LinkedList

// Value LinkedList's value struct.
type Value[T interface{}] struct {
	next *Value[T]
	last *Value[T]
	v    T
}

// LinkedList LinkedList struct.
type LinkedList[T interface{}] struct {
	start  *Value[T]
	end    *Value[T]
	now    *Value[T]
	length int
	index  int
}

// Start Get the value of the first element.
func (ll *LinkedList[T]) Start() T {
	return ll.start.v
}

// End Get the value of the last element.
func (ll *LinkedList[T]) End() T {
	return ll.end.v
}

// Get Get the value of the current element.
func (ll *LinkedList[T]) Get() T {
	return ll.now.v
}

// Set Set the value of the current element.
func (ll *LinkedList[T]) Set(v T) {
	if ll.now == nil || ll.length == 0 {
		panic("the LinkedList's length is 0")
	}
	ll.now.v = v
}

// IsStart Check if it is the first element.
func (ll *LinkedList[T]) IsStart() bool {
	if ll.index == 0 {
		return true
	} else {
		return false
	}
}

// IsEnd Check if it is the last element.
func (ll *LinkedList[T]) IsEnd() bool {
	if ll.index == ll.length-1 {
		return true
	} else {
		return false
	}
}

// Forward Forward to the next element.
func (ll *LinkedList[T]) Forward() {
	ll.now = ll.now.next
	if ll.index == ll.length-1 {
		ll.index = 0
	} else {
		ll.index += 1
	}
}

// Back Back to the last element.
func (ll *LinkedList[T]) Back() {
	ll.now = ll.now.last
	if ll.index == 0 {
		ll.index = ll.length - 1
	} else {
		ll.index -= 1
	}
}

// Length Get the LinkedList's length.
func (ll *LinkedList[T]) Length() int {
	return ll.length
}

// Index Get the current index.
func (ll *LinkedList[T]) Index() int {
	return ll.index
}

// Insert Insert a value before this element
func (ll *LinkedList[T]) Insert(v T) {
	appe := &Value[T]{
		next: ll.now,
		last: ll.now.last,
		v:    v,
	}
	if ll.now == nil {
		ll.now = appe
		appe.next = appe
		appe.last = appe
		ll.end = appe
		ll.start = appe
	}
	now := ll.now
	las := ll.now.last
	now.last = appe
	las.next = appe
	if now == ll.start {
		ll.start = appe
	}
	ll.index += 1
	ll.length += 1
}

// Append Append a value to the end of the LinkedList.
func (ll *LinkedList[T]) Append(v T) {
	appe := &Value[T]{
		last: ll.end,
		next: ll.start,
		v:    v,
	}
	if ll.now == nil {
		ll.now = appe
		appe.next = appe
		appe.last = appe
		ll.end = appe
		ll.start = appe
	}
	end := ll.end
	sta := ll.start
	end.next = appe
	sta.last = appe
	ll.end = appe
	ll.length += 1
}

// Delete Delete the current value and back to the last value.
func (ll *LinkedList[T]) Delete() {
	now := ll.now
	las := now.last
	nex := now.next
	las.next = nex
	nex.last = las
	if now == ll.end {
		ll.end = las
	}
	if now == ll.start {
		ll.start = nex
	}
	ll.length -= 1
	ll.now = las
	ll.index -= 1
}

// Positions:Some positions.
const (
	// START Represents the first value of the LinkedList.
	START = iota
	// END Represents the last value's pointer of the LinkedList.
	END
	// BEFORE Represents the value before 'now'.
	BEFORE
	// AFTER Represents the value after 'now'.
	AFTER
)

// InsertTo Add the value to the item corresponding to position.
func (ll *LinkedList[T]) InsertTo(idx int, v T) {
	if idx == START {
		appe := &Value[T]{
			last: ll.end,
			next: ll.start,
			v:    v,
		}
		ll.end.next = appe
		ll.start.last = appe
		ll.start = appe
		ll.length += 1
		ll.index += 1
	} else if idx == END {
		appe := &Value[T]{
			last: ll.end,
			next: ll.start,
			v:    v,
		}
		ll.end.next = appe
		ll.start.last = appe
		ll.end = appe
		ll.length += 1
	} else if idx == BEFORE {
		appe := &Value[T]{
			last: ll.now.last,
			next: ll.now,
			v:    v,
		}
		ll.now.last.next = appe
		ll.now.last = appe
		if ll.now == ll.start {
			ll.start = appe
		}
		ll.length += 1
		ll.index += 1
	} else if idx == AFTER {
		appe := &Value[T]{
			last: ll.now,
			next: ll.now.next,
			v:    v,
		}
		ll.now.next.last = appe
		ll.now.next = appe
		if ll.now == ll.end {
			ll.end = appe
		}
		ll.length += 1
	}
}

// Goto Jump to the item corresponding to the position(in var 'START','END','BEFORE' and 'AFTER').
func (ll *LinkedList[T]) Goto(idx int) {
	if idx == START {
		ll.now = ll.start
		ll.index = 0
	} else if idx == END {
		ll.now = ll.end
		ll.index = ll.length - 1
	} else if idx == BEFORE {
		ll.now = ll.now.last
		ll.index -= 1
	} else if idx == AFTER {
		ll.now = ll.now.next
		ll.index += 1
	}
}

// GetByPosition Get the value of the item by selecting an position(in var 'START','END','BEFORE' and 'AFTER').
func (ll *LinkedList[T]) GetByPosition(idx int) T {
	if idx == START {
		return ll.start.v
	} else if idx == END {
		return ll.end.v
	} else if idx == BEFORE {
		return ll.now.last.v
	} else if idx == AFTER {
		return ll.now.next.v
	} else {
		panic("could not find this position, please select one in 'START','END','BEFORE' and 'AFTER'")
	}
}
