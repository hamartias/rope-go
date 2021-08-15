package main

import (
	"fmt"
)

// RopeNode is a tree node
// If left == right == nil, it is a leaf
type RopeNode struct {
	piece  string
	weight int
	left   *RopeNode
	right  *RopeNode
}

// NewRopeNode creates a rope node from a string
func NewRopeNode(p string) *RopeNode {
	return &RopeNode{piece: p, weight: len(p)}
}

// isLeaf returns true if the node is a leaf
func (rn *RopeNode) isLeaf() bool {
	return (rn.left == nil) && (rn.right == nil)
}

// Split breaks a rope into two ropes
//    let s be a string, i be the given index
//    returns s[:i], s[i:]
func (rn *RopeNode) Split(index int) (rnl, rnr *RopeNode) {
	if rn.isLeaf() {
		if index == 0 {
			return &RopeNode{piece: ""}, &RopeNode{piece: rn.piece, weight: rn.weight}
		}
		p1 := rn.piece[:index]
		p2 := rn.piece[index:]
		return &RopeNode{piece: p1, weight: len(p1)}, &RopeNode{piece: p2, weight: len(p2)}
	}
	if rn.weight > index {
		r1, r2 := rn.left.Split(index)
		newRight := &RopeNode{weight: r2.leafWeight(), left: r2, right: rn.right}
		return r1, newRight
	}
	r1, r2 := rn.right.Split(index - rn.weight)
	newLeft := &RopeNode{weight: rn.left.leafWeight(), left: rn.left, right: r1}
	return newLeft, r2
}

// Print shows node contents, mainly for dev use
func (rn *RopeNode) Print() {
	fmt.Printf("piece: %s, weight: %d\n", rn.piece, rn.weight)
}

// PrintRec calls print recursively and indents to give a visual indication of
// tree structure
func (rn *RopeNode) PrintRec() {
	var _helper func(prefix string, rn *RopeNode)
	_helper = func(prefix string, rn *RopeNode) {
		fmt.Printf("%s", prefix)
		rn.Print()
		if rn.left != nil {
			_helper(prefix+" ", rn.left)
		}
		if rn.right != nil {
			_helper(prefix+" ", rn.right)
		}
	}
	_helper("", rn)
}

// Index returns the character at the given index as a rune
func (rn *RopeNode) Index(i int) rune {
	if rn.weight <= i && rn.right != nil {
		return rn.right.Index(i - rn.weight)
	}
	if rn.left != nil {
		return rn.left.Index(i)
	}
	return rune(rn.piece[i])
}

// leafWeight returns the sum of len(leaf.piece) for all leaf nodes descendent
// of the given rope node.
func (rn *RopeNode) leafWeight() int {
	ret := len(rn.piece)
	if rn.left != nil {
		ret += rn.left.leafWeight()
	}
	if rn.right != nil {
		ret += rn.right.leafWeight()
	}
	return ret
}
