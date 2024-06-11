package queue

import "errors"

const defaultCapacity = 2 << 8

var ErrEmptyQueue error = errors.New("queue empty")

// Queue is a FIFO data structure
type Queue[T any] struct {
	underlying []T
}

func New[T any]() *Queue[T] {
	return NewWithOptions(
		WithCapacity[T](defaultCapacity),
	)
}

func NewWithOptions[T any](opts ...Option[T]) *Queue[T] {
	q := new(Queue[T])
	for _, opt := range opts {
		opt(q)
	}
	return q
}

type Option[T any] func(q *Queue[T])

func WithCapacity[T any](capacity int) Option[T] {
	return func(q *Queue[T]) {
		q.underlying = make([]T, 0, capacity)
	}
}

// Poll returns the element at the head of the queue,
// or returns an error if the queue is empty.
func (q *Queue[T]) Poll() (T, error) {
	if q.Len() < 1 {
		return *new(T), ErrEmptyQueue
	}
	poll := q.underlying[0]
	q.underlying = q.underlying[1:]
	return poll, nil
}

// Push adds a new element to the end of the queue
func (q *Queue[T]) Push(elem T) {
	q.underlying = append(q.underlying, elem)
}

// Len returns the current number of elements occupying
// the queue.
func (q *Queue[T]) Len() int {
	return len(q.underlying)
}
