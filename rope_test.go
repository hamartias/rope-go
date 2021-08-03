package main

import (
  "testing"
  "fmt"
)

func prints(s string) {
  fmt.Println(s)
}

func buildFailString(got, expected string) string {
  return fmt.Sprintf("\ngot \"%s\"\nwanted \"%s\"\n", got, expected)
}

func TestRopeNodeStack(t *testing.T) {
  t.Run("Stack ops", func(t *testing.T) {
    s := NewRopeNodeStack()
    pieces := []*RopeNode{NewRopeNode("1"), NewRopeNode("2"), NewRopeNode("3")}
    for _, p := range(pieces) {
      s.Push(p)
    }
    if s.Len() != 3 {
      t.Fail()
    }
    for i := len(pieces)-1; i >= 0; i-- {
      got, _ := s.Pop()
      expected := pieces[i]
      if got != expected {
        t.Fail()
      }
    }

    // Throws error popping from empty stack
    _, err := s.Pop()
    if err == nil {
      t.Fail()
    }
  })
}
func TestRopeNode(t *testing.T) {
  t.Run("Test isLeaf", func(t *testing.T) {
    rn := &RopeNode{piece: "s", weight: 1}
    if rn.isLeaf() == false {
      t.Fail()
    }
  })
}

func TestRopeCreation(t *testing.T) {
  t.Run("MakeRope smoke test", func(t *testing.T) {
    s := "hello world"
    rope := MakeRope(s)
    got := rope.root.left.piece + rope.root.right.piece
    if got != s {
      t.Errorf("got %s, wanted %s", got, s)
    }
  })

  t.Run("MakeRope from multiple strings", func(t *testing.T) {
    input := []string{"hello", " world", " test", " string"}
    expected := ""
    for _, s := range(input) {
      expected += s
    }
    rope := MakeRopeFromSlice(input)
    got := rope.Report(0, len(expected)-1)
    if got != expected {
      t.Errorf(buildFailString(got, expected))
    }
  })
}

func TestReport(t *testing.T) {
  testReport := func(s string, startIndex, endIndex int) bool {
    ok := true
    rope := MakeRope(s)
    got := rope.Report(startIndex, endIndex)
    expected := s[startIndex:endIndex+1]
    if got != expected {
      ok = false
      fmt.Printf("got %s, wanted %s\n", got, expected)
    }
    return ok
  }

  t.Run("Report smoke test", func(t *testing.T) {
    s := "test string"
    ok := testReport(s, 0, len(s)-1)
    if !ok { t.Fail() }
  })

  t.Run("Report part of string smoke test", func(t *testing.T) {
    s := "test string"
    ok := testReport(s, 0, len(s)-2)
    if !ok { t.Fail() }
  })

  t.Run("Report finds the correct root to traverse from", func(t *testing.T) {
  })
}
