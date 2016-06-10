package main

import (
	"encoding/json"
	"log"
)

type binary_node struct {
	Id int

	Left  *binary_node
	Right *binary_node
}

type binary_tree struct {
	node_count uint16
	Root       *binary_node
}

func (t *binary_tree) Insert(v int) {
	node, parent := find_node(t.Root, v)
	if node != nil {
		return
	}

	t.node_count++
	node = &binary_node{Id: v}
	if parent == nil {
		t.Root = node
	} else {
		var child *binary_node
		if parent.Id > v {
			child = parent.Left
			parent.Left = node
		} else if parent.Id < v {
			child = parent.Right
			parent.Right = node
		}

		if child != nil {
			if child.Id > node.Id {
				node.Right = child
			} else if child.Id < node.Id {
				node.Left = child
			}
		}
	}
}

func (t *binary_tree) Delete(id int) {
	log.Printf("delete node %v", id)
	node, parent := find_node(t.Root, id)
	if node == nil {
		return
	}

	delete_node(parent, node)
	t.node_count--
	log.Printf("deleted %v", node.Id)
}

func (t *binary_tree) Find(id int) (res interface{}, ok bool) {
	node, _ := find_node(t.Root, id)
	ok = node != nil
	if ok {
		res = node.Id
	}

	return
}

func (t *binary_tree) Count() uint16 {
	return t.node_count
}

func (t *binary_tree) Min() int {
	n := t.Root
	for {
		if n.Left == nil {
			return n.Id
		}

		n = n.Left
	}
}

func (t *binary_tree) Max() int {
	n := t.Root
	for {
		if n.Right == nil {
			return n.Id
		}

		n = n.Right
	}
}

func (t *binary_tree) Print() {
	json, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(json))
}

func (t *binary_tree) Asc(f func(int)) {
	ord(t.Root, f, true)
}

func (t *binary_tree) Desc(f func(int)) {
	ord(t.Root, f, false)
}

func ord(n *binary_node, f func(int), asc bool) {
	var open, close *binary_node
	if asc {
		open, close = n.Left, n.Right
	} else {
		open, close = n.Right, n.Left
	}

	if open != nil {
		ord(open, f, asc)
	}

	f(n.Id)

	if close != nil {
		ord(close, f, asc)
	}
}

func delete_node(parent *binary_node, node *binary_node) {
	switch {
	case node.Left == nil && node.Right != nil:
		{
			log.Printf("- case right replace")
			if parent.Id > node.Id {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	case node.Left != nil && node.Right == nil:
		{
			log.Printf("- case left replace")
			if parent.Id > node.Id {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		}
	case node.Left == nil && node.Right == nil:
		{
			log.Printf("- case simple delete")
			if parent.Id > node.Id {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
	case node.Left != nil && node.Right != nil:
		{
			log.Printf("- case next replace")
			new_node, parent := next_node(node)
			// copy
			node.Id = new_node.Id
			delete_node(parent, new_node)
		}
	}
}

func find_node(root *binary_node, id int) (node *binary_node, parent *binary_node) {
	log.Println("--")
	log.Printf("search %v", id)
	if root == nil {
		return
	}

	node = root
	for {
		if node == nil {
			break
		}

		log.Printf("find (%v)", node.Id)
		switch {
		case node.Id == id:
			log.Printf("-- equal")
			return
		case node.Id > id:
			log.Printf("-- left")
			parent = node
			node = parent.Left
		case node.Id < id:
			log.Printf("-- right")
			parent = node
			node = parent.Right
		}
	}

	node = nil
	log.Printf("result: %v -> %v", parent.Id, node)
	return
}

func next_node(node *binary_node) (next *binary_node, parent *binary_node) {
	parent = node
	next = node.Right
	log.Printf("get next for: %v", next.Id)
	if next == nil {
		return
	}
	for {
		if next.Left == nil {
			return
		}

		parent = next
		next = next.Left
		log.Printf("- left %v", next.Id)
	}

	return
}

func NewTree() Tree {
	return &binary_tree{}
}
