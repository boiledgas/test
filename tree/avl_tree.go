package tree

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type avl_node struct {
	Id int

	Balance int8

	parent *avl_node // это поле
	Left   *avl_node `json:",omitempty"`
	Right  *avl_node `json:",omitempty"`
}

type Avl_tree struct {
	Len  uint16
	Root *avl_node
}

func (t *Avl_tree) Insert(id int) {
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

func (t *Avl_tree) Delete(id int) (err error) {
	node := find(t.Root, id)
	if node == nil || node.Id != id {
		err = errors.New(fmt.Sprintf("id %v not found", id))
		return
	}
	t.delete(node)
	return
}

func (t *Avl_tree) delete(node *avl_node) {
	parent := node.parent
	// определение направления движения
	var delta int8
	switch {
	case node.parent == nil:
		delta = 0
	case node.parent.Left == node:
		delta = 1
	case node.parent.Right == node:
		delta = -1
	}
	// удаление узла из структуры дерева
	switch {
	case node.Left == nil && node.Right == nil:
		if t.Root == node {
			t.Root = nil
		}
		switch delta {
		case 1:
			node.parent.Left = nil
		case -1:
			node.parent.Right = nil
		}
	case node.Left != nil && node.Right == nil:
		if t.Root == node {
			t.Root = node.Left
		}
		node.Left.parent = node.parent
		switch delta {
		case 1:
			node.parent.Left = node.Left
		case -1:
			node.parent.Right = node.Left
		}
	case node.Left == nil && node.Right != nil:
		if t.Root == node {
			t.Root = node.Right
		}
		node.Right.parent = node.parent
		switch delta {
		case 1:
			node.parent.Left = node.Right
		case -1:
			node.parent.Right = node.Right
		}
	case node.Left != nil && node.Right != nil:
		nearest := next_node(node)
		parent = nearest.parent
		// copy
		node.Id = nearest.Id
		if nearest.parent != node {
			// удаление из поддерева
			nearest.parent.Left = nearest.Right
			if nearest.Right != nil {
				nearest.Right.parent = nearest.parent
			}
			delta = 1
		} else {
			// удаление правого элемента
			node.Right = nearest.Right
			if nearest.Right != nil {
				nearest.Right.parent = node
			}
			delta = -1
		}

		nearest.parent = nil
		nearest.Left = nil
		nearest.Right = nil
	}

	//free(node)
	// балансировка структуры дерева
balance:
	for parent != nil {
		parent.Balance += delta
		// -2, 1,-1 =LR> 1, 0, 0
		// -2, 1, 0 =LR> 0, 0, 0
		// -2, 1, 1 =LR>-1, 0, 0
		// -2, 0    =R> -1, 1
		// -2,-1    =R>  0 ,0
		//  2, 1    =L>  0, 0
		//  2, 0    =L>  1,-1
		//  2,-1,-1 =RL> 0, 1, 0
		//  2,-1, 0 =RL> 0, 0, 0
		//  2,-1, 1 =RL> 1, 0, 0
		switch {
		case parent.Balance == -1 || parent.Balance == 1:
			break balance
		case parent.Balance == -2:
			son := parent.Left
			switch {
			case son.Balance == 1:
				if parent.Right == nil {
					parent.Balance, son.Balance = 0, 0
				} else {
					subtree := son.Right
					switch {
					case subtree.Balance == -1:
						parent.Balance, son.Balance, subtree.Balance = 1, 0, 0
					case subtree.Balance == 0:
						parent.Balance, son.Balance, subtree.Balance = 0, 0, 0
					case subtree.Balance == 1:
						parent.Balance, son.Balance, subtree.Balance = -1, 0, 0
					}
				}
				son := node_left_rotate(parent.Left)
				son.parent, parent.Left = parent, son
				parent = node_right_rotate(parent)
			case son.Balance == 0:
				parent.Balance, son.Balance = -1, 1
				parent = node_right_rotate(parent)
			case son.Balance == -1:
				parent.Balance, son.Balance = 0, 0
				parent = node_right_rotate(parent)
			}
		case parent.Balance == 2:
			son := parent.Right
			switch {
			case son.Balance == 1:
				parent.Balance, son.Balance = 0, 0
				parent = node_left_rotate(parent)
			case son.Balance == 0:
				parent.Balance, son.Balance = 1, -1
				parent = node_left_rotate(parent)
			case son.Balance == -1:
				if parent.Left == nil {
					parent.Balance, son.Balance = 0, 0
				} else {
					subtree := son.Left
					switch {
					case subtree.Balance == -1:
						parent.Balance, son.Balance, subtree.Balance = 0, 1, 0
					case subtree.Balance == 0:
						parent.Balance, son.Balance, subtree.Balance = 0, 0, 0
					case subtree.Balance == 1:
						parent.Balance, son.Balance, subtree.Balance = 1, 0, 0
					}
				}
				son := node_right_rotate(parent.Right)
				son.parent, parent.Right = parent, son
				parent = node_left_rotate(parent)
			}
		}
		node = parent
		parent = parent.parent
		switch {
		case parent == nil:
			t.Root = node
		case parent.Left == node:
			delta = 1
		case parent.Right == node:
			delta = -1
		}
	}
}

func node_right_rotate(node *avl_node) *avl_node {
	son := node.Left
	subtree := son.Right
	son.parent = node.parent
	node.parent = son
	if subtree != nil {
		subtree.parent = node
	}
	node.Left = subtree
	son.Right = node
	return son
}

func node_left_rotate(node *avl_node) *avl_node {
	son := node.Right
	subtree := son.Left
	son.parent = node.parent
	node.parent = son
	if subtree != nil {
		subtree.parent = node
	}
	node.Right = subtree
	son.Left = node
	return son
}

func next_node(node *avl_node) (next *avl_node) {
	next = node.Right
	for next.Left != nil {
		next = next.Left
	}
	return
}

func (t *Avl_tree) Find(id int) (result interface{}, ok bool) {
	n := find(t.Root, id)
	ok = n != nil
	result = n.Id
	return
}

func (t *Avl_tree) Count() uint16 {
	return t.Len
}

func (t *Avl_tree) Min() int {
	node := t.Root
	for {
		if node.Left == nil {
			break
		}
		node = node.Left
	}
	return node.Id
}

func (t *Avl_tree) Max() int {
	node := t.Root
	for {
		if node.Right == nil {
			break
		}
		node = node.Right
	}
	return node.Id
}

func (t *Avl_tree) Asc(func(int)) {
}

func (t *Avl_tree) Desc(func(int)) {
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

func (t *Avl_tree) rotate(node *avl_node) {
	var parent, son, subtree *avl_node
	son, parent = node, node.parent
	var delta int8
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
			switch {
			case son.Left != nil && subtree.Id == son.Left.Id && parent.Left != nil && son.Id == parent.Left.Id:
				right_rotate(parent, son)
				parent.Balance, son.Balance = 0, 0
				parent, son = son, parent
			case son.Right != nil && subtree.Id == son.Right.Id && parent.Right != nil && son.Id == parent.Right.Id:
				left_rotate(parent, son)
				parent.Balance, son.Balance = 0, 0
				parent, son = son, parent
			case son.Right != nil && subtree.Id == son.Right.Id && parent.Left != nil && son.Id == parent.Left.Id:
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
				right_rotate(parent, subtree)
				if node.Id != subtree.Id {
					parent.Balance, subtree.Balance = son.Balance+1, 0
				} else {
					parent.Balance, subtree.Balance = 0, 0
				}
				parent = subtree
			case son.Left != nil && subtree.Id == son.Left.Id && parent.Right != nil && son.Id == parent.Right.Id:
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
				if subtree.Balance == 2 {
					parent.Balance, subtree.Balance = -1, 0
				} else {
					parent.Balance, subtree.Balance = 0, 0
				}
				left_rotate(parent, subtree)
				parent = subtree
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

func (t *Avl_tree) Print(w io.Writer) {
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

func (t *Avl_tree) PrintFile(path string) {
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
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v(%v)", color=%v];`, n.Left.Id, n.Id, n.Left.Balance, color))
		doc.WriteString(fmt.Sprintf("%v -> %v [color=black];", n.Id, n.Left.Id))
	} else {
		doc.WriteString(fmt.Sprintf("nl%v [shape=point];", n.Id))
		doc.WriteString(fmt.Sprintf("%v -> nl%v [color=blue];", n.Id, n.Id))
	}
	if n.Right != nil {
		if n.Right.Balance == -2 || n.Right.Balance == 2 {
			color = "red"
		}
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v(%v)", color=%v];`, n.Right.Id, n.Id, n.Right.Balance, color))
		doc.WriteString(fmt.Sprintf("%v -> %v [color=blue];", n.Id, n.Right.Id))
	} else {
		doc.WriteString(fmt.Sprintf("nr%v [shape=point];", n.Id))
		doc.WriteString(fmt.Sprintf("%v -> nr%v [color=blue];", n.Id, n.Id))
	}

	if n.Left != nil {
		print_node(n.Left, doc)
	}
	if n.Right != nil {
		print_node(n.Right, doc)
	}
}
