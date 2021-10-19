package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

type Tree struct {
	head      *node
	numElmnts int
}

func (t *Tree) add(value int) {
	newNode := new(node)
	newNode.value = value
	if t.numElmnts == 0 {
		t.head = newNode
		t.numElmnts++
		return
	}
	t.head.add(newNode)
}

func (n1 *node) add(n2 *node) {
	if n1.value < n2.value {
		if n1.right == nil {
			n1.right = n2
			return
		} else {
			right := n1.right
			right.add(n2)
		}
	} else {
		if n1.left == nil {
			n1.left = n2
			return
		} else {
			left := n1.left
			left.add(n2)
		}
	}
}

func (t Tree) search(value int) *node {
	if t.head == nil {
		return nil
	}

	return t.head.search(value)
}

func (n node) search(value int) *node {
	if &n == nil {
		return nil
	}

	if value == n.value {
		return &n
	}

	if value < n.value {
		return n.left.search(value)
	} else {
		return n.right.search(value)
	}
}

func (n *node) maxRight() *node {
	if n.right == nil {
		return n
	}

	return n.right.maxRight()
}

func (n *node) remove(value int, parent *node, t *Tree) bool {
	if n.value == value {
		if n.left == nil && n.right == nil {
			n = nil
			if parent.value > value {
				parent.left = n
			} else if parent.value < value {
				parent.right = n
			}
			return true
		}

		if n.left == nil {
			n = n.right
			return true
		}

		head := n
		n = n.left
		if parent != nil {
			if parent.value > value {
				parent.left = n
			} else if parent.value < value {
				parent.right = n
			}
		} 
		maxRight := n.maxRight()
		maxRight.right = head.right
		head = n
		if parent == nil {
			t.setHead(head)
		}
		head = nil
		return true
	}

	if n.value > value {
		if n.left == nil {
			return false
		}
		return n.left.remove(value, n, t)
	}

	if n.value < value {
		if n.right == nil {
			return false
		}
		return n.right.remove(value, n, t)
	}
	return false
}

func (t *Tree) setHead(n *node) {
	t.head = n
}

func (t *Tree) remove(value int) {
	if t.head == nil {
		fmt.Println("The tree is empty!")
		return
	}

	if t.head.remove(value, nil, t) {
		fmt.Printf("the value %d is removed from the tree\n", value)
	} else {
		fmt.Printf("the value %d wasn't found in the tree\n", value)
	}
}

func (n *node) String() string {
	if n == nil {
		return ""
	}
	ret := strings.Builder{}

	ret.WriteString(fmt.Sprintf("[ "))
	ret.WriteString(fmt.Sprintf(n.left.String()))
	ret.WriteString(fmt.Sprintf("( %d ) ", n.value))
	ret.WriteString(fmt.Sprintf(n.right.String()))
	ret.WriteString(fmt.Sprintf(" ]"))

	return ret.String()
}

func (t Tree) String() string {
	if t.numElmnts == 0 {
		return "The Tree is empty"
	}
	ret := strings.Builder{}
	ret.WriteString(t.head.String())

	return ret.String()
}

func main() {
	T := new(Tree)
	T.add(5)
	T.add(3)
	T.add(4)
	T.add(1)
	T.add(7)
	T.add(6)
	T.add(8)

	fmt.Println(T)

	T.remove(8)

	fmt.Println(T)

}
