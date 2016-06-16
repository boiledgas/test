package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var counter int = 0

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
	counter++
	index := fmt.Sprintf("0000%v", counter)
	index = index[len(index)-4:]
	defer t.PrintFile(fmt.Sprintf("%v) insert%v(2).jpg", index, id))

	parent := find(t.Root, id)
	if parent != nil && parent.Id == id {
		return
	}
	switch {
	case parent == nil:
		t.Root = &avl_node{Id: id}
	case parent.Id == id:
		return
	case parent != nil:
		t.Len++
		node := &avl_node{Id: id, parent: parent}
		switch {
		case parent.Id > id:
			parent.Left = node
		case parent.Id < id:
			parent.Right = node
		}

		t.rotate(node)
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

func (t *avl_tree) rotate(node *avl_node) {
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
		if parent.Balance == 0 {
			return
		}
		if parent.Balance > 1 || parent.Balance < -1 {
			index := fmt.Sprintf("0000%v", counter)
			index = index[len(index)-4:]
			switch {
			case son.Left != nil && subtree.Id == son.Left.Id && parent.Left != nil && son.Id == parent.Left.Id:
				{
					t.PrintFile(fmt.Sprintf("%v) insert%v(0)-left.jpg", index, node.Id))
					right_rotate(parent, son)
					parent.Balance, son.Balance = 0, 0
					parent, son = son, parent
				}
			case son.Right != nil && subtree.Id == son.Right.Id && parent.Right != nil && son.Id == parent.Right.Id:
				{
					t.PrintFile(fmt.Sprintf("%v) insert%v(0)-right.jpg", index, node.Id))
					left_rotate(parent, son)
					parent.Balance, son.Balance = 0, 0
					parent, son = son, parent
				}
			case son.Right != nil && subtree.Id == son.Right.Id && parent.Left != nil && son.Id == parent.Left.Id:
				{
					t.PrintFile(fmt.Sprintf("%v) insert%v(0)-leftright.jpg", index, node.Id))
					left_rotate(son, subtree)
					switch {
					case subtree.Balance == 0:
						subtree.Balance = -son.Balance // 1 => -1; -1 => 1
						son.Balance = 0                // 1, -1 => 0
					case subtree.Balance == -1:
						subtree.Balance = -2 // -1 => -2
						son.Balance = 0      // 1 => 0
					case subtree.Balance == 1:
						subtree.Balance = -1 // 1 => -1
						son.Balance = -1     // 1 => -1
					}
					t.PrintFile(fmt.Sprintf("%v) insert%v(1)-leftright.jpg", index, node.Id))
					right_rotate(parent, subtree)
					if node.Id != subtree.Id {
						parent.Balance, subtree.Balance = son.Balance+1, 0
					} else {
						parent.Balance, subtree.Balance = 0, 0
					}
					parent = subtree
				}
			case son.Left != nil && subtree.Id == son.Left.Id && parent.Right != nil && son.Id == parent.Right.Id:
				{
					t.PrintFile(fmt.Sprintf("%v) insert%v(0)-rightleft.jpg", index, node.Id))
					right_rotate(son, subtree)
					switch {
					case subtree.Balance == 0:
						subtree.Balance = -son.Balance // 1 => -1; -1 => 1
						son.Balance = 0                // 1, -1 => 0
					case subtree.Balance == 1:
						subtree.Balance = 2 // 1 => 2
						son.Balance = 0     // -1 => 0
					case subtree.Balance == -1:
						subtree.Balance = 1 // -1 => 1
						son.Balance = 1     // 1 => 1
					}
					t.PrintFile(fmt.Sprintf("%v) insert%v(1)-rightleft.jpg", index, node.Id))
					if subtree.Balance == 2 {
						parent.Balance, subtree.Balance = -1, 0
					} else {
						parent.Balance, subtree.Balance = 0, 0
					}
					left_rotate(parent, subtree)
					parent = subtree
				}
			default:
				panic("invalid")
			}
			if parent.parent == nil {
				t.Root = parent
			}
			return
		}
		subtree = son
		son = parent
		parent = parent.parent
	}
}

func right_rotate(parent *avl_node, son *avl_node) {
	son.parent = parent.parent
	if parent.parent != nil {
		if parent.parent.Left != nil && parent.parent.Left.Id == parent.Id {
			parent.parent.Left = son
		} else if parent.parent.Right != nil && parent.parent.Right.Id == parent.Id {
			parent.parent.Right = son
		}
	}
	parent.parent = son
	parent.Left = son.Right
	if son.Right != nil {
		son.Right.parent = parent
	}
	son.Right = parent
}

func left_rotate(parent *avl_node, son *avl_node) {
	son.parent = parent.parent
	if parent.parent != nil {
		if parent.parent.Left != nil && parent.parent.Left.Id == parent.Id {
			parent.parent.Left = son
		} else if parent.parent.Right != nil && parent.parent.Right.Id == parent.Id {
			parent.parent.Right = son
		}
	}
	parent.parent = son
	parent.Right = son.Left
	if son.Left != nil {
		son.Left.parent = parent
	}
	son.Left = parent
}

func (t *avl_tree) Print(w io.Writer) {
	if t.Root == nil {
		return
	}

	if w == nil {
		json, err := json.Marshal(t)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(json))
		return
	}

	n := t.Root
	var doc bytes.Buffer

	doc.WriteString("digraph AvlTree {")
	color := "black"
	if n.Balance == -2 || n.Balance == 2 {
		color = "red"
	}
	doc.WriteString(fmt.Sprintf("%v [shape=circle, xlabel=%v, color=%v];", n.Id, n.Balance, color))
	print_node(n, &doc)
	doc.WriteString("}")

	w.Write(doc.Bytes())
}

func (t *avl_tree) PrintFile(path string) {
	f, err := ioutil.TempFile("", "graph")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer os.Remove(f.Name())

	t.Print(f)
	cmd := exec.Command("dot", "-Tjpg", f.Name(), "-o", path)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func print_node(n *avl_node, doc *bytes.Buffer) {
	color := "green"
	if n.Balance == -2 || n.Balance == 2 {
		color = "black"
	}

	if n.Left != nil {
		if n.Left.Balance == -2 || n.Left.Balance == 2 {
			color = "red"
		}
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v(%v)", color=%v];`, n.Left.Id, n.Left.Balance, n.Id, color))
		doc.WriteString(fmt.Sprintf("%v -> %v [color=black];", n.Id, n.Left.Id))
	} else if n.Right != nil {
		//doc.WriteString(fmt.Sprintf("nl%v [shape=point];", n.Id))
		//doc.WriteString(fmt.Sprintf("%v -> nl%v [color=blue];", n.Id))
	}
	if n.Right != nil {
		if n.Right.Balance == -2 || n.Right.Balance == 2 {
			color = "red"
		}
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v(%v)", color=%v];`, n.Right.Id, n.Right.Balance, n.Id, color))
		doc.WriteString(fmt.Sprintf("%v -> %v [color=blue];", n.Id, n.Right.Id))
	} else if n.Left != nil {
		//doc.WriteString(fmt.Sprintf("nr%v [shape=point];", n.Id))
		//doc.WriteString(fmt.Sprintf("%v -> nr%v [color=blue];", n.Id))
	}

	if n.Left != nil {
		print_node(n.Left, doc)
	}
	if n.Right != nil {
		print_node(n.Right, doc)
	}
}
