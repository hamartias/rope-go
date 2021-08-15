package main

// Rope contains the root of a tree as a RopeNode
type Rope struct {
	root *RopeNode
}

// NewRope converts a string into a Rope
func NewRope(s string) *Rope {
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
		for i := slen / 2; i < slen; i++ {
			sr += string(s[i])
		}
		root.left = &RopeNode{piece: sl, weight: 0}
		root.right = &RopeNode{piece: sr, weight: 0}
		root.weight = len(sl)
	}
	return ret
}

// Concat concatenates two ropes
func (r *Rope) Concat(rr *Rope) {
	newroot := &RopeNode{left: r.root, right: rr.root, weight: r.root.leafWeight()}
	r.root = newroot
}

// Split turns one rope node into two at the given index,'
// e.g. r.Split(i) == r[:i], r[i:]
func (r *Rope) Split(index int) (r1, r2 *Rope) {
	lr, rr := r.root.Split(index)
	return &Rope{root: lr}, &Rope{root: rr}
}

// Insert puts a string into the rope at index i
func (r *Rope) Insert(i int, s string) {
	newr := NewRope(s)
	ll, rr := r.Split(i)
	ll.Concat(newr)
	ll.Concat(rr)
	r.root = ll.root
}

// Delete removes a substring from the rope
// e.g. r.Delete(0, 5) removes the substring r[0:5]
func (r *Rope) Delete(si, ei int) {
	ll, rr := r.Split(si)
	_, ar := rr.Split(ei - si)
	ll.Concat(ar)
	r.root = ll.root
}

// Index returns the character at index i as a rune
func (r *Rope) Index(i int) rune {
	return r.root.Index(i)
}

// Report converts a Rope into a string
func (r *Rope) Report(startIndex, endIndex int) string {
	// Optimization: find the actual traversal node, not just from root
	traversalRoot := r.root
	var inOrder func(rn *RopeNode) string
	inOrder = func(rn *RopeNode) string {
		out := ""
		if rn != nil {
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
	return out[startIndex : endIndex+1]
}

// NewRopeFromSlice takes a slice of strings as leaf nodes and creates a Rope
func NewRopeFromSlice(sl []string) *Rope {
	// convert to RopeNodes
	nodes := make([]*RopeNode, len(sl))
	for i, s := range sl {
		// Leaf nodes have weight 0
		nodes[i] = &RopeNode{piece: s, weight: len(s)}
	}

	// Populate stack
	stack := NewRopeNodeStack()
	for i := len(nodes) - 1; i >= 0; i-- {
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
