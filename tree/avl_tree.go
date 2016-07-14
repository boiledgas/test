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

var Log_tree bool

type avl_node struct {
	Id int32

	Balance int8

	parent *avl_node // это поле
	Left   *avl_node `json:",omitempty"`
	Right  *avl_node `json:",omitempty"`
}

func (n *avl_node) String() string {
	if n == nil {
		return ""
	}

	var b bytes.Buffer
	b.WriteString("{")
	if n.Left != nil {
		b.WriteString(fmt.Sprintf("%v(%v):", n.Left.Id, n.Left.Balance))
		if n.Left.parent != nil {
			b.WriteString(fmt.Sprintf("%v", n.Left.parent.Id))
		} else {
			b.WriteString("n")
		}
	} else {
		b.WriteString("null ")
	}
	b.WriteString(" < ")
	b.WriteString(fmt.Sprintf("%v(%v):", n.Id, n.Balance))
	if n.parent != nil {
		b.WriteString(fmt.Sprintf("%v", n.parent.Id))
	} else {
		b.WriteString("n")
	}
	b.WriteString(" > ")
	if n.Right != nil {
		b.WriteString(fmt.Sprintf("%v(%v):", n.Right.Id, n.Right.Balance))
		if n.Right.parent != nil {
			b.WriteString(fmt.Sprintf("%v", n.Right.parent.Id))
		} else {
			b.WriteString("n")
		}
	} else {
		b.WriteString("null ")
	}
	b.WriteString("}")
	return b.String()
}

type Avl_tree struct {
	Len  uint16
	Root *avl_node
}

func (t *Avl_tree) Insert(id int32) {
	parent := find(t.Root, id)
	if parent != nil && parent.Id == id {
		return
	}
	var node *avl_node
	node = new(avl_node)
	node.Id, node.parent = id, parent
	t.Len++
	if parent == nil {
		t.Root = node
		return
	}
	switch {
	case parent.Id > id:
		parent.Left = node
	case parent.Id < id:
		parent.Right = node
	}

	// балансировка структуры дерева
	var i uint8
	son := node
	var subtree *avl_node
balance:
	for parent != nil {
		switch {
		case parent.Left == son:
			parent.Balance -= 1
		case parent.Right == son:
			parent.Balance += 1
		}
		switch parent.Balance {
		case -2:
			switch son.Balance {
			case 1: // LR
				subtree = son.Right
				if parent.Right == nil {
					parent.Balance, son.Balance = 0, 0
				} else {
					switch subtree.Balance {
					case 1:
						parent.Balance, son.Balance, subtree.Balance = 0, -1, 0
					case -1:
						parent.Balance, son.Balance, subtree.Balance = 1, 0, 0
					default:
						panic("not implemented")
					}
				}
				node_left_rotate(son)
				son, parent = parent, node_right_rotate(parent)
			case -1: // R
				parent.Balance, son.Balance = 0, 0
				son, parent = parent, node_right_rotate(parent)
			default: // R
				panic("not implemented")
			}
		case 2:
			switch son.Balance {
			case 1: // L
				parent.Balance, son.Balance = 0, 0
				parent = node_left_rotate(parent)
			case -1: // RL
				subtree = son.Left
				if parent.Left == nil {
					parent.Balance, son.Balance = 0, 0
				} else {
					switch subtree.Balance {
					case 1:
						parent.Balance, son.Balance, subtree.Balance = -1, 0, 0
					case -1:
						parent.Balance, son.Balance, subtree.Balance = 0, 1, 0
					default:
						panic("not implemented")
					}
				}
				node_right_rotate(son)
				parent = node_left_rotate(parent)
			default:
				panic("not implemented")
			}
		}
		if parent.parent == nil {
			t.Root = parent
		}
		if parent.Balance == 0 {
			break balance
		}
		son, parent = parent, parent.parent
		i++
	}
}

func (t *Avl_tree) Delete(id int32) (err error) {
	node := find(t.Root, id)
	if node == nil || node.Id != id {
		err = errors.New(fmt.Sprintf("id %v not found", id))
		return
	}
	t.delete(node)
	return
}

func (t *Avl_tree) Find(id int32) (result interface{}, ok bool) {
	n := find(t.Root, id)
	ok = n != nil
	result = n.Id
	return
}

func (t *Avl_tree) Count() uint16 {
	return t.Len
}

func (t *Avl_tree) Min() int32 {
	node := t.Root
	for {
		if node.Left == nil {
			break
		}
		node = node.Left
	}
	return node.Id
}

func (t *Avl_tree) Max() int32 {
	node := t.Root
	for {
		if node.Right == nil {
			break
		}
		node = node.Right
	}
	return node.Id
}

func (t *Avl_tree) Asc(func(int32)) {
}

func (t *Avl_tree) Desc(func(int32)) {
}

func (t *Avl_tree) Validate() (err error) {
	_, err = node_height(t.Root)
	return
}

func node_height(node *avl_node) (height int8, err error) {
	if node == nil {
		return
	}

	if node.Left != nil && node.Left.Id >= node.Id || node.Right != nil && node.Right.Id <= node.Id {
		err = errors.New(fmt.Sprintf("%v node structure not valid", node))
		return
	}

	if node.Left != nil && (node.Left.parent == nil || node.Left.parent.Id != node.Id) {
		err = errors.New(fmt.Sprintf("%v left parent wrong right(%v)-left(%v)", node, node.Right, node.Left))
		return
	}
	if node.Right != nil && (node.Right.parent == nil || node.Right.parent.Id != node.Id) {
		err = errors.New(fmt.Sprintf("%v right parent wrong right(%v)-left(%v)", node, node.Right, node.Left))
		return
	}

	var left_height int8
	if left_height, err = node_height(node.Left); err != nil {
		return
	}
	var right_height int8
	if right_height, err = node_height(node.Right); err != nil {
		return
	}

	if node.Balance != right_height-left_height {
		err = errors.New(fmt.Sprintf("%v balance not valid right(%v)-left(%v)", node, right_height, left_height))
	}

	if height = left_height; left_height < right_height {
		height = right_height
	}
	height++
	return
}

func find(root *avl_node, id int32) (node *avl_node) {
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

func (t *Avl_tree) delete(node *avl_node) {
	if Log_tree {
		t.PrintFile("tree_init.jpg")
	}
	parent := node.parent
	// определение направления движения
	// удаление узла из структуры дерева
	switch {
	case node.Left == nil && node.Right == nil:
		if t.Root == node {
			t.Root = nil
		}
		switch {
		case parent == nil:
		case parent.Left == node:
			parent.Left = nil
			parent.Balance += 1
		case parent.Right == node:
			parent.Right = nil
			parent.Balance -= 1
		}
	case node.Left != nil && node.Right == nil:
		if t.Root == node {
			t.Root = node.Left
		}
		node.Left.parent = node.parent
		switch {
		case parent == nil:
			return
		case parent.Left == node:
			parent.Left = node.Left
			parent.Balance += 1
		case parent.Right == node:
			parent.Right = node.Left
			parent.Balance -= 1
		}
	case node.Left == nil && node.Right != nil:
		if t.Root == node {
			t.Root = node.Right
		}
		node.Right.parent = node.parent
		switch {
		case parent == nil:
			return
		case parent.Left == node:
			parent.Left = node.Right
			parent.Balance += 1
		case parent.Right == node:
			parent.Right = node.Right
			parent.Balance -= 1
		}
	case node.Left != nil && node.Right != nil:
		nearest := next_node(node)
		parent = nearest.parent
		// copy
		node.Id = nearest.Id
		nearest.Id = 0
		if parent.Id != node.Id {
			parent.Left = nearest.Right
			parent.Balance += 1
		} else {
			parent.Right = nearest.Right
			parent.Balance -= 1
		}
		if nearest.Right != nil {
			nearest.Right.parent = nearest.parent
		}


		nearest.parent = nil
		nearest.Left = nil
		nearest.Right = nil
		//free(nearest)
	}
	//free(node)
	var son *avl_node
	var subtree *avl_node
	i := 0
balance:
	for parent != nil {
		if Log_tree {
			t.PrintFile(fmt.Sprintf("tree_%v.jpg", i))
			i++
		}
		switch parent.Balance {
		case -1:
			break balance
		case 1:
			break balance
		case -2:
			son = parent.Left
			switch son.Balance {
			case 1:
				if parent.Right == nil {
					parent.Balance, son.Balance = 0, 0
				} else {
					subtree = son.Right
					switch subtree.Balance {
					case -1:
						parent.Balance, son.Balance, subtree.Balance = 1, 0, 0
					case 0:
						parent.Balance, son.Balance, subtree.Balance = 0, 0, 0
					case 1:
						parent.Balance, son.Balance, subtree.Balance = 0, -1, 0
					}
				}
				node_left_rotate(son)
				if Log_tree {
					t.PrintFile(fmt.Sprintf("tree_%v.jpg", i))
					i++
				}
				parent = node_right_rotate(parent)
			case 0:
				parent.Balance, son.Balance = -1, 1
				parent = node_right_rotate(parent)
			case -1:
				parent.Balance, son.Balance = 0, 0
				parent = node_right_rotate(parent)
			}
		case 2:
			son = parent.Right
			switch son.Balance {
			case 1:
				parent.Balance, son.Balance = 0, 0
				parent = node_left_rotate(parent)
			case 0:
				parent.Balance, son.Balance = 1, -1
				parent = node_left_rotate(parent)
			case -1:
				if parent.Left == nil {
					parent.Balance, son.Balance = 0, 0
				} else {
					subtree = son.Left
					switch subtree.Balance {
					case -1:
						parent.Balance, son.Balance, subtree.Balance = 0, 1, 0
					case 0:
						parent.Balance, son.Balance, subtree.Balance = 0, 0, 0
					case 1:
						parent.Balance, son.Balance, subtree.Balance = -1, 0, 0
					}
				}
				node_right_rotate(son)
				if Log_tree {
					t.PrintFile(fmt.Sprintf("tree_%v.jpg", i))
					i++
				}
				parent = node_left_rotate(parent)
			}
		default:
			node, parent = parent, parent.parent
			switch {
			case parent == nil:
			case parent.Left == node:
				parent.Balance += 1
			case parent.Right == node:
				parent.Balance -= 1
			}
		}
	}

	if parent == nil {
		t.Root = node
	}
	if Log_tree {
		t.PrintFile("tree_result.jpg")
	}
}

func node_right_rotate(node *avl_node) *avl_node {
	parent := node.parent
	son := node.Left
	subtree := son.Right
	if parent != nil {
		switch {
		case parent.Left == node:
			parent.Left = son
		case parent.Right == node:
			parent.Right = son
		}
	}
	son.parent = node.parent
	son.Right = node
	node.parent = son
	node.Left = subtree
	if subtree != nil {
		subtree.parent = node
	}
	return son
}

func node_left_rotate(node *avl_node) *avl_node {
	parent := node.parent
	son := node.Right
	subtree := son.Left
	if parent != nil {
		switch {
		case parent.Left == node:
			parent.Left = son
		case parent.Right == node:
			parent.Right = son
		}
	}
	son.parent = node.parent
	son.Left = node
	node.parent = son
	node.Right = subtree
	if subtree != nil {
		subtree.parent = node
	}
	return son
}

func next_node(node *avl_node) (next *avl_node) {
	next = node.Right
	for next.Left != nil {
		next = next.Left
	}
	return
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
		pid := "n"
		if n.Left.parent != nil {
			pid = fmt.Sprintf("%v", n.Left.parent.Id)
		}
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v(%v)", color=%v];`, n.Left.Id, pid, n.Left.Balance, color))
		doc.WriteString(fmt.Sprintf("%v -> %v [color=black];", n.Id, n.Left.Id))
	} else {
		doc.WriteString(fmt.Sprintf("nl%v [shape=point];", n.Id))
		doc.WriteString(fmt.Sprintf("%v -> nl%v [color=blue];", n.Id, n.Id))
	}
	if n.Right != nil {
		if n.Right.Balance == -2 || n.Right.Balance == 2 {
			color = "red"
		}
		pid := "n"
		if n.Right.parent != nil {
			pid = fmt.Sprintf("%v", n.Right.parent.Id)
		}
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v(%v)", color=%v];`, n.Right.Id, pid, n.Right.Balance, color))
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
