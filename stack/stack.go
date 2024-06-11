package stack

import "errors"

const defaultCapacity = 2 << 8

var ErrEmptyStack error = errors.New("empty stack")

// Stack is a LIFO data structure
type Stack[T any] struct {
	underlying []T
}

func New[T any]() *Stack[T] {
	return NewWithOptions(
		WithCapacity[T](defaultCapacity),
	)
}

type Option[T any] func(s *Stack[T])

func WithCapacity[T any](capacity int) Option[T] {
	return func(s *Stack[T]) {
		s.underlying = make([]T, 0, capacity)
	}
}

func NewWithOptions[T any](opts ...Option[T]) *Stack[T] {
	s := new(Stack[T])
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Pop removes the element from the top of the stack,
// i.e the most recently pushed item
func (s *Stack[T]) Pop() (T, error) {
	if s.Len() < 1 {
		return *new(T), ErrEmptyStack
	}
	p := s.underlying[len(s.underlying)-1]
	s.underlying = s.underlying[:len(s.underlying)-1]
	return p, nil
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(elem T) {
	s.underlying = append(s.underlying, elem)
}

// Peek returns the top element from the stack, without
// removing it
func (s *Stack[T]) Peek() T {
	return s.underlying[len(s.underlying)-1]
}

// Len returns the length of the stack, i.e how many
// elements currently occupy the stack
func (s *Stack[T]) Len() int {
	return len(s.underlying)
}
