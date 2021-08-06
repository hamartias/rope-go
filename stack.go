package main

import (
  "fmt"
)
type RopeNodeStack struct {
  stack []*RopeNode
}

func NewRopeNodeStack() *RopeNodeStack {
  return &RopeNodeStack{stack: make([]*RopeNode, 0)}
}

func (s *RopeNodeStack) Len() int {
  return len(s.stack)
}

func (s *RopeNodeStack) Push(rn *RopeNode) {
  s.stack = append(s.stack, rn)
}

func (s *RopeNodeStack) Pop() (*RopeNode, error) {
  if len(s.stack) == 0 {
    return nil, fmt.Errorf("Trying to pop from empty stack")
  }
  li := len(s.stack)-1
  ret := (s.stack)[li]
  s.stack = (s.stack)[:li]
  return ret, nil
}

func (s *RopeNodeStack) Print() {
  for i, e := range(s.stack) {
    fmt.Printf("%d: ", i)
    fmt.Println(*e)
  }
}
