package radix

import (
	"fmt"
	"strings"
)

type leafNode struct {
	key string
	val interface{}
}

type edge struct {
	label byte
	node  *node
}

type edges []edge

type node struct {
	leaf   *leafNode
	prefix string
	edges  edges
}

func (n *node) isLeaf() bool {
	return n.leaf != nil
}

func (n *node) getEdge(label byte) *node {
	num := len(n.edges)
	fmt.Println(num)

	return nil
}

type Tree struct {
	root *node
	size int
}

func New() *Tree {
	//return &Tree{root: &node{}}
	return NewFromMap(nil)
}

func NewFromMap(input map[string]interface{}) *Tree {
	tree := &Tree{root: &node{}}
	for k, v := range input {
		tree.Insert(k, v)
	}
	return tree
}

func (t *Tree) Len() int {
	return t.size
}

func (t *Tree) Insert(s string, value interface{}) (interface{}, bool) {
	//var parent *node
	//n := t.root
	search := s

	for {
		if len(search) == 0 {
		}

		return nil, false
	}
}

func (t *Tree) Delete(s string) (interface{}, bool) {
	n := t.root
	search := s

	for {
		if len(search) == 0 {
			if !n.isLeaf() {
				break
			}
			//do a delete
		}
	}
	return nil, false
}

func (t *Tree) Get(s string) (interface{}, bool) {
	n := t.root
	search := s

	for {
		if len(search) == 0 {
			if n.isLeaf() {
				return n.leaf.val, true
			}
			break
		}
		n = n.getEdge(search[0])
		if n == nil {
			break
		}
		if strings.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]
		} else {
			break
		}
	}
	return nil, false
}

func (t *Tree) LongestPrefix() {
}

func (t *Tree) Minimum() {
}

func (t *Tree) Maximum() {
}

func (t *Tree) Walk() {
}

func (t *Tree) WalkPrefix() {
}

func (t *Tree) WalkPath() {
}
