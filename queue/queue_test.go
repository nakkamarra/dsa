package queue

import (
	"errors"
	"testing"
)

func TestQueue(t *testing.T) {

	t.Run("/primitive", func(t *testing.T) {
		q := New[string]()
		q.Push("a")
		q.Push("b")
		q.Push("c")

		p1, _ := q.Poll()
		p2, _ := q.Poll()
		p3, _ := q.Poll()

		if p1 != "a" {
			t.Fatalf("expected %s but got %s", p1, "a")
		}
		if p2 != "b" {
			t.Fatalf("expected %s but got %s", p2, "b")
		}
		if p3 != "c" {
			t.Fatalf("expected %s but got %s", p1, "c")
		}

	})

	t.Run("/struct", func(t *testing.T) {
		type SomeStruct struct {
			a int
			b string
		}
		q := New[SomeStruct]()

		s1 := SomeStruct{a: 0, b: "zero"}
		s2 := SomeStruct{a: 1, b: "one"}
		s3 := SomeStruct{a: 2, b: "two"}

		q.Push(s1)
		q.Push(s2)
		q.Push(s3)

		p1, _ := q.Poll()
		p2, _ := q.Poll()
		p3, _ := q.Poll()

		if p1 != s1 {
			t.Fatalf("expected %v but got %v", s1, p1)
		}
		if p2 != s2 {
			t.Fatalf("expected %v but got %v", s2, p2)
		}
		if p3 != s3 {
			t.Fatalf("expected %v but got %v", s3, p3)
		}
	})

	t.Run("/pointers", func(t *testing.T) {
		type SomeStruct struct {
			a int
			b string
		}
		q := New[*SomeStruct]()

		s1 := &SomeStruct{a: 0, b: "zero"}
		s2 := &SomeStruct{a: 1, b: "one"}
		s3 := &SomeStruct{a: 2, b: "two"}

		q.Push(s1)
		q.Push(s2)
		q.Push(s3)

		p1, _ := q.Poll()
		p2, _ := q.Poll()
		p3, _ := q.Poll()

		if p1 != s1 {
			t.Fatalf("expected %v but got %v", s1, p1)
		}
		if p2 != s2 {
			t.Fatalf("expected %v but got %v", s2, p2)
		}
		if p3 != s3 {
			t.Fatalf("expected %v but got %v", s3, p3)
		}
	})

	t.Run("/empty", func(t *testing.T) {
		q := New[int]()
		_, err := q.Poll()
		if !errors.Is(err, ErrEmptyQueue) {
			t.Fatalf("expected empty queue error after poll on empty queue")
		}
	})
}
