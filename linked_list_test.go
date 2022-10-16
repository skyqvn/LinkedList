package LinkedList

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Append(1)
	ll.Append(2)
	// t1:1,2,
	fmt.Println("t1")
	for i := 0; i < ll.length; i++ {
		fmt.Print(ll.Get(), ",")
		ll.Forward()
	}
	fmt.Println("\n------")
	ll.Goto(START)
	ll.Insert(3)
	ll.Goto(START)
	// t2:3,1,2,
	fmt.Println("t2")
	for i := 0; i < ll.length; i++ {
		fmt.Print(ll.Get(), ",")
		ll.Forward()
	}
	fmt.Println("\n------")
	ll.InsertTo(END, 4)
	ll.Goto(START)
	// t3:3,1,2,4,
	fmt.Println("t3")
	for i := 0; i < ll.length; i++ {
		fmt.Print(ll.Get(), ",")
		ll.Forward()
	}
	fmt.Println("\n------")
	// t4:4,3,1,2,
	fmt.Println("t4")
	ll.Goto(END)
	for i := 0; i < ll.length; i++ {
		fmt.Print(ll.Get(), ",")
		ll.Forward()
	}
	fmt.Println("\n------")
	// t5:3,1,2,5,4,
	fmt.Println("t5")
	ll.Forward()
	ll.Forward()
	ll.Forward()
	ll.InsertTo(AFTER, 5)
	ll.Goto(START)
	for i := 0; i < ll.length; i++ {
		fmt.Print(ll.Get(), ",")
		ll.Forward()
	}
	fmt.Println()
}
