package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	root := &List[string]{val: "Hello"}
	root.next = &List[string]{val: "World"}
	root.next.next = &List[string]{val: "Awesome!"}

	p := root
	for p != nil {
		fmt.Println(p.val)
		p = p.next
	}
}
