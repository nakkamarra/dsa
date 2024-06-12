package bst

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var ErrEmptyTree = errors.New("empty tree")
var ErrValuePresent = errors.New("value already present")
var ErrNotPresent = errors.New("value not present")

type Node[T constraints.Ordered] struct {
	Value       T
	Left, Right *Node[T]
}

func (n *Node[T]) Insert(value T) error {
	if n == nil {
		n = &Node[T]{
			Value: value,
		}
		return nil
	}
	if n.Value > value {
		return n.Left.Insert(value)
	}
	if n.Value < value {
		return n.Right.Insert(value)
	}
	return ErrValuePresent
}

func (n *Node[T]) Search(query T) (*Node[T], error) {
	if n == nil {
		return nil, ErrNotPresent
	}
	if n.Value == query {
		return n, nil
	}
	if n.Value > query {
		return n.Left.Search(query)
	}
	if n.Value < query {
		return n.Right.Search(query)
	}
	return nil, ErrNotPresent
}
