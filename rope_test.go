package main

import (
  "testing"
  "fmt"
  "math/rand"
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

func TestIndex(t *testing.T) {
  t.Run("Index smoke test", func(t *testing.T) {
    s := "string"
    r := MakeRope(s)
    for i := 0; i < len(s); i++ {
      got := r.Index(i)
      expected := rune(s[i])
      if got != expected {
        t.Errorf("got %b, wanted %b", got, expected)
      }
    }
  })

  t.Run("Index works for larger tree", func(t *testing.T) {
    input := make([]string, 0)
    rand.Seed(1) // replicatable, easy to change later on for complex testing
    for i := 0; i < 20; i++ {
      input = append(input, randString(i))
    }
    reference := ""
    for _, s := range input {
      reference += s
    }
    rope := MakeRopeFromSlice(input)
    for i := 0; i < len(reference); i++ {
      got := rope.Index(i)
      expected := rune(reference[i])
      if got != expected {
        t.Errorf("got %s, wanted %s for index %d", string(got), string(expected), i)
      }
    }
  })
}

func randString(n int) string {
  letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
  b := make([]rune, n)
  for i := range b {
    b[i] = letters[rand.Intn(len(letters))]
  }
  return string(b)
}

func TestConcat(t *testing.T) {
  t.Run("Concat smoke test", func(t *testing.T) {
    s1 := "test "
    s2 := "string"
    r1 := MakeRope(s1)
    r2 := MakeRope(s2)
    r1.Concat(r2)
    got := r1.Report(0, len(s1) + len(s2)-1)
    if got != s1+s2 {
      t.Fail()
    }
  })
}

func TestSplit(t *testing.T) {
  t.Run("Split smoke test", func(t *testing.T) {
    s := []string{"test ", "this ", "op ", "with ", "a ", "rope"}
    r := MakeRopeFromSlice(s)
    expected := "test this op with a rope"
    for i := 0; i < len(expected); i++ {
      r1, r2 := r.Split(i)
      got1 := r1.Report(0, len(expected))
      got2 := r2.Report(0, len(expected))
      if got1 + got2 != expected {
        t.Errorf("1: %s, 2: %s", got1, got2)
      }
    }
  })
}
