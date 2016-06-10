package main

import (
	"encoding/json"
	"log"
)

type avl_node struct {
	Id int

	Balance int

	parent *avl_node // это поле
	Left   *avl_node `json:",omitempty"`
	Right  *avl_node `json:",omitempty"`
}

type avl_tree struct {
	Len  uint16
	Root *avl_node
}

func NewAvlTree() Tree {
	return &avl_tree{}
}

func (t *avl_tree) Insert(id int) {
	log.Printf("insert %v", id)
	parent := find(t.Root, id)
	switch {
	case parent == nil:
		log.Printf("- root node")
		t.Root = &avl_node{Id: id}
	case parent.Id == id:
		log.Printf("- key exists")
		return
	case parent != nil:
		t.Len++
		node := &avl_node{Id: id, parent: parent}
		switch {
		case parent.Id > id:
			log.Printf("- set %v left %v", id, parent.Id)
			parent.Left = node
		case parent.Id < id:
			log.Printf("- set %v right %v", id, parent.Id)
			parent.Right = node
		}

		find_support_node(t, node)
		t.Print()
	}
}

func (t *avl_tree) Delete(int) {
}

func (t *avl_tree) Find(id int) (result interface{}, ok bool) {
	n := find(t.Root, id)
	ok = n != nil
	result = n.Id
	return
}

func (t *avl_tree) Count() uint16 {
	return t.Len
}

func (t *avl_tree) Min() int {
	node := t.Root
	for {
		if node.Left == nil {
			break
		}
		node = node.Left
	}
	return node.Id
}

func (t *avl_tree) Max() int {
	node := t.Root
	for {
		if node.Right == nil {
			break
		}
		node = node.Right
	}
	return node.Id
}

func (t *avl_tree) Asc(func(int)) {
}

func (t *avl_tree) Desc(func(int)) {
}

func (t *avl_tree) Print() {
	json, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(json))
}

func find(root *avl_node, id int) (node *avl_node) {
	for {
		switch {
		case root == nil:
			return
		case root.Id == id:
			node = root
			return
		case root.Id > id:
			node = root
			root = root.Left
		case root.Id < id:
			node = root
			root = root.Right
		}
	}
	return
}

// mode:
// 1) вставка в левое поддерево левого сына
// 2) вставка в правое поддерево правого сына
// 3) вставка в правое поддерево левого сына
// 4) вставка в левое поддерево правого сына
func find_support_node(tree *avl_tree, node *avl_node) {
	var parent, son, subtree *avl_node
	son, parent = node, node.parent
	var delta int
	for {
		delta = 0
		switch {
		case parent == nil:
			return
		case parent.Left != nil && parent.Left.Id == son.Id:
			delta = -1
		case parent.Right != nil && parent.Right.Id == son.Id:
			delta = 1
		}

		parent.Balance += delta
		log.Printf("-- node:%v; balance:%v; delta:%v", parent.Id, parent.Balance, delta)
		if parent.Balance > 1 || parent.Balance < -1 {
			switch {
			case subtree.Id == son.Left.Id && son.Id == parent.Left.Id:
				parent.Balance, son.Balance = 0, 0
				son.parent = parent.parent
				parent.parent = son
				son.Right = parent
				parent.Left = nil
				if tree.Root == parent {
					tree.Root = son
				}
			case subtree.Id == son.Right.Id && son.Id == parent.Right.Id:
				parent.Balance, son.Balance = 0, 0
				son.parent = parent.parent
				parent.parent = son
				son.Left = parent
				parent.Right = nil
				if tree.Root == parent {
					tree.Root = son
				}
			case subtree.Id == son.Right.Id && son.Id == parent.Left.Id:
				//3
			case subtree.Id == son.Left.Id && son.Id == parent.Right.Id:
				//4
			}

			log.Printf("rotate: support %v son %v subtree %v", parent.Id, son.Id, subtree.Id)
			log.Printf("-- after: node:%v; balance:%v; delta:%v", parent.Id, parent.Balance, delta)
		}

		subtree = son
		son = parent
		parent = parent.parent
	}
	return
}
