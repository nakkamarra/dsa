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

// New returns the root node of a tree with nil subtrees,
// and a specified value.
func New[T constraints.Ordered](value T) *Node[T] {
	return &Node[T]{
		Value: value,
	}
}

// NewFrom returns the root node of a tree built using the
// values slice as inputs.
func NewFrom[T constraints.Ordered](values []T) *Node[T] {
	if len(values) < 1 {
		return nil
	}
	root := New[T](values[0])
	for i := 1; i < len(values); i++ {
		root.Insert(values[i])
	}
	return root
}

// Insert attempts to insert a value into the binary search tree
// at the node from which it is called, returning an error if the
// value is already present at the current node or in one of its
// subtrees.
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

// Search queries the binary search tree for a given value,
// returning a reference to the node containing that value if it
// is found, or an error if no node with that value exists.
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

// Delete will remove a node with the given value from the
// binary search tree, readjusting children / parent values
// if necessary
func (n *Node[T]) Delete(value T) error {
	if n == nil {
		return ErrNotPresent
	}
	if n.Value == value {
		if n.Left == nil && n.Right == nil {
			n = nil
			return nil
		}
		return nil
	}
	if n.Value > value {
		return n.Left.Delete(value)
	}
	if n.Value < value {
		return n.Right.Delete(value)
	}
	return ErrNotPresent
}

func (n *Node[T]) Rebalance() {

}
