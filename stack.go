package main

import (
	"fmt"
)

// RopeNodeStack wraps a slice of RopeNode refs for stack ops
type RopeNodeStack struct {
	stack []*RopeNode
}

// NewRopeNodeStack creates a RopeNodeStack and returns a reference
func NewRopeNodeStack() *RopeNodeStack {
	return &RopeNodeStack{stack: make([]*RopeNode, 0)}
}

// Len returns the number of elements in the stack
func (s *RopeNodeStack) Len() int {
	return len(s.stack)
}

// Push adds an element to the top of the stack
func (s *RopeNodeStack) Push(rn *RopeNode) {
	s.stack = append(s.stack, rn)
}

// Pop removes an element from the top of the stack
func (s *RopeNodeStack) Pop() (*RopeNode, error) {
	if len(s.stack) == 0 {
		return nil, fmt.Errorf("Trying to pop from empty stack")
	}
	li := len(s.stack) - 1
	ret := (s.stack)[li]
	s.stack = (s.stack)[:li]
	return ret, nil
}

// Print shows stack contents for dev use
func (s *RopeNodeStack) Print() {
	for i, e := range s.stack {
		fmt.Printf("%d: ", i)
		fmt.Println(*e)
	}
}
