package tree

import (
	"fmt"
)

type node struct {
	data   int
	left   *node
	right  *node
	parent *node
}

type BST struct {
	root *node
}

func (b *BST) IsEmpty() bool {
	return b.root == nil
}

func (b *BST) Insert(i int) {
	temp := &node{data: i, left: nil, right: nil, parent: nil}

	var y *node = nil
	var x *node = b.root

	 if b.IsEmpty() {
	 	b.root = temp
	 } else{
		for x != nil {
			y = x
			if i < x.data {
				x = x.left
			} else {
				x = x.right
			}
		}

		if i < y.data {
			y.left = temp
			fmt.Printf("nilai %d masuk sebelah kiri %d \n", i, y.data)
		} else {
			y.right = temp
			fmt.Printf("nilai %d masuk sebelah kanan %d \n", i, y.data)
		}
	}
}