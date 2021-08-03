package main

import (
  "fmt"
)

type RopeNode struct {
  piece string
  weight int
  left *RopeNode
  right *RopeNode
}

func NewRopeNode(p string) *RopeNode {
  return &RopeNode{piece: p, weight: len(p)}
}

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


func (rn *RopeNode) isLeaf() bool {
  return (rn.left == nil) && (rn.right == nil)
}

type Rope struct {
  root *RopeNode
}

func MakeRope(s string) *Rope {
  slen := len(s)
  root := &RopeNode{piece: "", weight: slen}
  ret := &Rope{root: root}
  if slen == 1 || slen == 0 {
    root.left = &RopeNode{piece: s, weight: slen}
  } else {
    var sl, sr string
    for i := 0; i < slen/2; i++ {
      sl += string(s[i])
    }
    for i := slen/2; i < slen; i++ {
      sr += string(s[i])
    }
    root.left = &RopeNode{piece: sl, weight: len(sl)}
    root.right = &RopeNode{piece: sr, weight: len(sr)}
  }
  return ret
}

func (r *Rope) Report(startIndex, endIndex int) string {
  // TODO: find the actual traversal node, not just from root
  traversalRoot := r.root
  var inOrder func(rn *RopeNode) string
  inOrder = func(rn *RopeNode) string {
    out := ""
    if (rn != nil) {
      if rn.isLeaf() {
        out += rn.piece
      } else {
        out += inOrder(rn.left)
        out += inOrder(rn.right)
      }
    }
    return out
  }
  out := inOrder(traversalRoot)
  if endIndex+1 > len(out) {
    return out[startIndex:]
  }
  return out[startIndex:endIndex+1]
}


func MakeRopeFromSlice(sl []string) *Rope {
  // convert to RopeNodes
  nodes := make([]*RopeNode, len(sl))
  for i, s := range(sl) {
    nodes[i] = &RopeNode{piece: s, weight: len(s)}
  }

  // Populate stack
  stack := NewRopeNodeStack()
  for i := len(nodes)-1; i >= 0; i-- {
    stack.Push(nodes[i])
  }

  // Reduce to tree
  for {
    if stack.Len() == 1 {
      break
    }
    temp := NewRopeNodeStack()
    for done := false; !done; done = stack.Len() < 2 {
      left, _ := stack.Pop()
      right, _ := stack.Pop()
      parent := &RopeNode{weight: left.weight+right.weight, left: left, right: right}
      temp.Push(parent)
    }
    if stack.Len() != 0 {
      leftover, _ := stack.Pop()
      temp.Push(leftover)
    }
    count := temp.Len()
    for i := 0; i < count; i++ {
      next, _ := temp.Pop()
      stack.Push(next)
    }
  }
  root, _ := stack.Pop()
  return &Rope{root: root}
}
