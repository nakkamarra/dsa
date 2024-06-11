package stack

import (
	"errors"
	"testing"
)

func TestStack(t *testing.T) {

	t.Run("/int", func(t *testing.T) {
		st := New[int]()
		st.Push(0)
		st.Push(1)
		st.Push(2)

		if st.Len() != 3 {
			t.Fatalf("expected length: %d got: %d", 3, st.Len())
		}
		p1, _ := st.Pop()
		if p1 != 2 {
			t.Fatalf("expected pop: %d got: %d", 2, p1)
		}
		p2, _ := st.Pop()
		if p2 != 1 {
			t.Fatalf("expected pop: %d got: %d", 1, p2)
		}
		p3, _ := st.Pop()
		if p3 != 0 {
			t.Fatalf("expected pop: %d got: %d", 0, p3)
		}
	})

	t.Run("/structs", func(t *testing.T) {
		type SomeStruct struct {
			a int
			b string
		}
		st := New[SomeStruct]()
		s1 := SomeStruct{a: 0, b: "zero"}
		s2 := SomeStruct{a: 1, b: "one"}
		s3 := SomeStruct{a: 2, b: "two"}
		st.Push(s1)
		st.Push(s2)
		st.Push(s3)
		p1, _ := st.Pop()
		if p1 != s3 {
			t.Fatalf("expected %v to equal %v", p1, s3)
		}
		p2, _ := st.Pop()
		if p2 != s2 {
			t.Fatalf("expected %v to equal %v", p2, s2)
		}
		p3, _ := st.Pop()
		if p3 != s1 {
			t.Fatalf("expected %v to equal %v", p3, s1)
		}
	})

	t.Run("/pointers", func(t *testing.T) {
		type SomeStruct struct {
			a int
			b string
		}
		st := New[*SomeStruct]()
		s1 := &SomeStruct{
			a: 0,
			b: "zero",
		}
		s2 := &SomeStruct{
			a: 1,
			b: "one",
		}
		s3 := &SomeStruct{
			a: 2,
			b: "two",
		}
		st.Push(s1)
		st.Push(s2)
		p1, _ := st.Pop()
		if p1 != s2 {
			t.Fatalf("expected %v to equal %v", p1, s2)
		}
		st.Push(s3)
		p2, _ := st.Pop()
		if p2 != s3 {
			t.Fatalf("expected %v to equal %v", p2, s3)
		}
		p3, _ := st.Pop()
		if p3 != s1 {
			t.Fatalf("expected %v to equal %v", p3, s1)
		}
	})

	t.Run("/empty", func(t *testing.T) {
		st := New[int]()
		_, err := st.Pop()
		if !errors.Is(err, ErrEmptyStack) {
			t.Fatalf("expected empty stack error when popping empty stack")
		}
	})
}
