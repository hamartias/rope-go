package main

import (
  "fmt"
)

func print(s string) {
  fmt.Println(s)
}

type RopeNode struct {
  piece string
  weight int
  left *RopeNode
  right *RopeNode
}

func NewRopeNode(p string) *RopeNode {
  return &RopeNode{piece: p, weight: len(p)}
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
    root.left = &RopeNode{piece: s, weight: 0}
  } else {
    var sl, sr string
    for i := 0; i < slen/2; i++ {
      sl += string(s[i])
    }
    for i := slen/2; i < slen; i++ {
      sr += string(s[i])
    }
    root.left = &RopeNode{piece: sl, weight: 0}
    root.right = &RopeNode{piece: sr, weight: 0}
    root.weight = len(sl)
  }
  return ret
}

func (r *Rope) Concat(rr *Rope) {
  newroot := &RopeNode{left: r.root, right: rr.root, weight: r.root.leafWeight()}
  r.root = newroot
}

func (r *Rope) Split(index int) (r1, r2 *Rope) {
  lr, rr := r.root.Split(index)
  return &Rope{root: lr}, &Rope{root: rr}
}

func (r *Rope) Insert(i int, s string) {
  newr := MakeRope(s)
  ll, rr := r.Split(i)
  ll.Concat(newr)
  ll.Concat(rr)
  r.root = ll.root
}

func (r *Rope) Delete(si, ei int) {
  ll, rr := r.Split(si)
  _, ar := rr.Split(ei - si)
  ll.Concat(ar)
  r.root = ll.root
}

func (rn *RopeNode) Split(index int) (rnl, rnr *RopeNode) {
  if rn.isLeaf() {
    if index == 0 {
      return &RopeNode{piece: ""}, &RopeNode{piece: rn.piece, weight: rn.weight}
    } else {
      p1 := rn.piece[:index]
      p2 := rn.piece[index:]
      return &RopeNode{piece: p1, weight: len(p1)}, &RopeNode{piece: p2, weight: len(p2)}
    }
  } else {
    if rn.weight > index {
      r1, r2 := rn.left.Split(index)
      newRight := &RopeNode{weight: r2.leafWeight(), left: r2, right: rn.right}
      return r1, newRight
    } else {
      r1, r2 := rn.right.Split(index - rn.weight)
      newLeft := &RopeNode{weight: rn.left.leafWeight(), left: rn.left, right: r1}
      return newLeft, r2
    }
  }
}

func (rn *RopeNode) Print() {
  fmt.Printf("piece: %s, weight: %d\n", rn.piece, rn.weight)
}

func (rn *RopeNode) PrintRec() {
  var _helper func(prefix string, rn *RopeNode)
  _helper = func(prefix string, rn *RopeNode) {
    fmt.Printf("%s", prefix)
    rn.Print()
    if rn.left != nil {
      _helper(prefix + " ", rn.left)
    }
    if rn.right != nil {
      _helper(prefix + " ", rn.right)
    }
  }
  _helper("", rn)
}

func (r *Rope) Index(i int) rune {
  return r.root.Index(i)
}

func (rn *RopeNode) Index(i int) rune {
  if rn.weight <= i && rn.right != nil {
    return rn.right.Index(i - rn.weight)
  }
  if rn.left != nil {
    return rn.left.Index(i)
  }
  return rune(rn.piece[i])
}

func (r *Rope) Report(startIndex, endIndex int) string {
  // Optimization: find the actual traversal node, not just from root
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
    // Leaf nodes have weight 0
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
    // want a final root node, otherwise can keep reducing, so go until len < 2
    for done := false; !done; done = stack.Len() < 2 {
      left, _ := stack.Pop()
      right, _ := stack.Pop()
      parent := &RopeNode{weight: left.leafWeight(), left: left, right: right}
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

func (rn *RopeNode) leafWeight() int {
  // UNTESTED
  ret := len(rn.piece)
  if rn.left != nil {
    ret += rn.left.leafWeight()
  }
  if rn.right != nil {
    ret += rn.right.leafWeight()
  }
  return ret
}
