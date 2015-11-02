package main

import "fmt"

func main() {
	a := stackFactory(3, 1, 5, 2)
	b := &Stack{}

	reverseSort(a, b)

	fmt.Println("a:")
	stackPrinter(a)
	fmt.Println("b:")
	stackPrinter(b)
}

func stackFactory(i ...int) *Stack {
	s := &Stack{}
	for _, i := range i {
		s.Push(i)
	}
	return s
}

func stackPrinter(s *Stack) {
	var e interface{}
	for !s.IsEmpty() {
		e = s.Pop()
		fmt.Printf("%d ", e.(int))
	}
	fmt.Printf("\n")
}

func reverseSort(a, b *Stack) {
	var smallest int
	counter := 0

	// first make b reverse sorted first
	smallest = findSmallest(a, b)
	moveFromBtoAExceptI(a, b, smallest, false)
	fmt.Printf("pushing %d to b\n", smallest)
	b.Push(smallest)

	// now that b is reverse sorted
	for !a.IsEmpty() && counter < 20 {
		smallest = findSmallest(a, b)
		moveFromBtoAExceptI(a, b, smallest, true)
		fmt.Printf("pushing %d to b\n", smallest)
		b.Push(smallest)
		counter++
	}
}

// find smallest from A, pushing elements to B
func findSmallest(a, b *Stack) int {
	s := int(^uint(0) >> 1) // MAX INT
	fmt.Printf("in findSmallest: ")
	var temp interface{}
	for !a.IsEmpty() {
		temp = a.Pop()
		fmt.Printf("%d ", temp)
		if s > temp.(int) {
			s = temp.(int)
		}
		b.Push(temp)
	}
	fmt.Printf("found %d\n", s)
	return s
}

// move items from B to A
// except when the item is i
// when i is found, break if the stack is sorted (the next item is smaller)
// or continue to move items if not
func moveFromBtoAExceptI(a, b *Stack, i int, returnOnI bool) {
	var temp interface{}
	fmt.Printf("in moveFromBtoAExceptI: ")
	found := false
	for !b.IsEmpty() {
		temp = b.Pop()
		if temp.(int) == i && !found {
			if b.Peek().(int) <= temp.(int) { // this means the stack is ordered
				break
			} else {
				found = true
			}
		} else {
			fmt.Printf("pushing %d to A, ", temp)
			a.Push(temp)
		}
	}
	fmt.Println("")
}

// credit: https://gist.github.com/bemasher/1777766
type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *Element
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Peek() (value interface{}) {
	return s.top.value
}
