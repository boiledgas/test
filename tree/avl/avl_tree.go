package avl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var Log_tree bool

type node struct {
	id    int32
	value interface{}

	balance int8

	parent *node // это поле
	left   *node
	right  *node
}

func (n *node) String() string {
	if n == nil {
		return ""
	}

	var b bytes.Buffer
	b.WriteString("{")
	if n.left != nil {
		b.WriteString(fmt.Sprintf("%v(%v):", n.left.id, n.left.balance))
		if n.left.parent != nil {
			b.WriteString(fmt.Sprintf("%v", n.left.parent.id))
		} else {
			b.WriteString("n")
		}
	} else {
		b.WriteString("null ")
	}
	b.WriteString(" < ")
	b.WriteString(fmt.Sprintf("%v(%v):", n.id, n.balance))
	if n.parent != nil {
		b.WriteString(fmt.Sprintf("%v", n.parent.id))
	} else {
		b.WriteString("n")
	}
	b.WriteString(" > ")
	if n.right != nil {
		b.WriteString(fmt.Sprintf("%v(%v):", n.right.id, n.right.balance))
		if n.right.parent != nil {
			b.WriteString(fmt.Sprintf("%v", n.right.parent.id))
		} else {
			b.WriteString("n")
		}
	} else {
		b.WriteString("null ")
	}
	b.WriteString("}")
	return b.String()
}

type Tree struct {
	Len  uint16
	Root *node
}

func (t *Tree) Insert(id int32, value interface{}) {
	parent := find(t.Root, id)
	if parent != nil && parent.id == id {
		return
	}
	var n *node
	n = new(node)
	n.id, n.parent, n.value = id, parent, value
	t.Len++
	if parent == nil {
		t.Root = n
		return
	}
	switch {
	case parent.id > id:
		parent.left = n
	case parent.id < id:
		parent.right = n
	}

	// балансировка структуры дерева
	var i uint8
	son := n
	var subtree *node
balance:
	for parent != nil {
		switch {
		case parent.left == son:
			parent.balance -= 1
		case parent.right == son:
			parent.balance += 1
		}
		switch parent.balance {
		case -2:
			switch son.balance {
			case 1: // LR
				subtree = son.right
				if parent.right == nil {
					parent.balance, son.balance = 0, 0
				} else {
					switch subtree.balance {
					case 1:
						parent.balance, son.balance, subtree.balance = 0, -1, 0
					case -1:
						parent.balance, son.balance, subtree.balance = 1, 0, 0
					default:
						panic("not implemented")
					}
				}
				node_left_rotate(son)
				son, parent = parent, node_right_rotate(parent)
			case -1: // R
				parent.balance, son.balance = 0, 0
				son, parent = parent, node_right_rotate(parent)
			default: // R
				panic("not implemented")
			}
		case 2:
			switch son.balance {
			case 1: // L
				parent.balance, son.balance = 0, 0
				parent = node_left_rotate(parent)
			case -1: // RL
				subtree = son.left
				if parent.left == nil {
					parent.balance, son.balance = 0, 0
				} else {
					switch subtree.balance {
					case 1:
						parent.balance, son.balance, subtree.balance = -1, 0, 0
					case -1:
						parent.balance, son.balance, subtree.balance = 0, 1, 0
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
		if parent.balance == 0 {
			break balance
		}
		son, parent = parent, parent.parent
		i++
	}
}

func (t *Tree) Delete(id int32) (value interface{}) {
	n := find(t.Root, id)
	if n == nil || n.id != id {
		return nil
	}

	value = n.value
	t.delete(n)
	return
}

func (t *Tree) Find(id int32) (result interface{}, ok bool) {
	n := find(t.Root, id)
	ok = n != nil
	result = n.value
	return
}

func (t *Tree) Count() uint16 {
	return t.Len
}

func (t *Tree) Min() interface{} {
	if t.Root == nil {
		return nil
	}
	n := t.Root
	for {
		if n.left == nil {
			break
		}
		n = n.left
	}
	return n.id
}

func (t *Tree) Max() interface{} {
	if t.Root == nil {
		return nil
	}
	n := t.Root
	for {
		if n.right == nil {
			break
		}
		n = n.right
	}
	return n.id
}

func (t *Tree) Asc(func(int32)) {
}

func (t *Tree) Desc(func(int32)) {
}

func (t *Tree) Validate() (err error) {
	_, err = node_height(t.Root)
	return
}

func node_height(n *node) (height int8, err error) {
	if n == nil {
		return
	}

	if n.left != nil && n.left.id >= n.id || n.right != nil && n.right.id <= n.id {
		err = errors.New(fmt.Sprintf("%v node structure not valid", n))
		return
	}

	if n.left != nil && (n.left.parent == nil || n.left.parent.id != n.id) {
		err = errors.New(fmt.Sprintf("%v left parent wrong right(%v)-left(%v)", n, n.right, n.left))
		return
	}
	if n.right != nil && (n.right.parent == nil || n.right.parent.id != n.id) {
		err = errors.New(fmt.Sprintf("%v right parent wrong right(%v)-left(%v)", n, n.right, n.left))
		return
	}

	var left_height int8
	if left_height, err = node_height(n.left); err != nil {
		return
	}
	var right_height int8
	if right_height, err = node_height(n.right); err != nil {
		return
	}

	if n.balance != right_height-left_height {
		err = errors.New(fmt.Sprintf("%v balance not valid right(%v)-left(%v)", n, right_height, left_height))
	}

	if height = left_height; left_height < right_height {
		height = right_height
	}
	height++
	return
}

func find(root *node, id int32) (n *node) {
	for {
		switch {
		case root == nil:
			return
		case root.id == id:
			n = root
			return
		case root.id > id:
			n = root
			root = root.left
		case root.id < id:
			n = root
			root = root.right
		}
	}
	return
}

func (t *Tree) delete(n *node) {
	if Log_tree {
		t.PrintFile("tree_init.jpg")
	}
	parent := n.parent
	// определение направления движения
	// удаление узла из структуры дерева
	switch {
	case n.left == nil && n.right == nil:
		if t.Root == n {
			t.Root = nil
		}
		switch {
		case parent == nil:
		case parent.left == n:
			parent.left = nil
			parent.balance += 1
		case parent.right == n:
			parent.right = nil
			parent.balance -= 1
		}
	case n.left != nil && n.right == nil:
		if t.Root == n {
			t.Root = n.left
		}
		n.left.parent = n.parent
		switch {
		case parent == nil:
			return
		case parent.left == n:
			parent.left = n.left
			parent.balance += 1
		case parent.right == n:
			parent.right = n.left
			parent.balance -= 1
		}
	case n.left == nil && n.right != nil:
		if t.Root == n {
			t.Root = n.right
		}
		n.right.parent = n.parent
		switch {
		case parent == nil:
			return
		case parent.left == n:
			parent.left = n.right
			parent.balance += 1
		case parent.right == n:
			parent.right = n.right
			parent.balance -= 1
		}
	case n.left != nil && n.right != nil:
		nearest := next_node(n)
		parent = nearest.parent
		// copy
		n.id, n.value = nearest.id, nearest.value
		if parent.id != n.id {
			parent.left = nearest.right
			parent.balance += 1
		} else {
			parent.right = nearest.right
			parent.balance -= 1
		}
		if nearest.right != nil {
			nearest.right.parent = nearest.parent
		}

		nearest.parent = nil
		nearest.left = nil
		nearest.right = nil
		nearest.value = nil
		//free(nearest)
	}
	//free(node)
	var son *node
	var subtree *node
	i := 0
balance:
	for parent != nil {
		if Log_tree {
			t.PrintFile(fmt.Sprintf("tree_%v.jpg", i))
			i++
		}
		switch parent.balance {
		case -1:
			break balance
		case 1:
			break balance
		case -2:
			son = parent.left
			switch son.balance {
			case 1:
				if parent.right == nil {
					parent.balance, son.balance = 0, 0
				} else {
					subtree = son.right
					switch subtree.balance {
					case -1:
						parent.balance, son.balance, subtree.balance = 1, 0, 0
					case 0:
						parent.balance, son.balance, subtree.balance = 0, 0, 0
					case 1:
						parent.balance, son.balance, subtree.balance = 0, -1, 0
					}
				}
				node_left_rotate(son)
				if Log_tree {
					t.PrintFile(fmt.Sprintf("tree_%v.jpg", i))
					i++
				}
				parent = node_right_rotate(parent)
			case 0:
				parent.balance, son.balance = -1, 1
				parent = node_right_rotate(parent)
			case -1:
				parent.balance, son.balance = 0, 0
				parent = node_right_rotate(parent)
			}
		case 2:
			son = parent.right
			switch son.balance {
			case 1:
				parent.balance, son.balance = 0, 0
				parent = node_left_rotate(parent)
			case 0:
				parent.balance, son.balance = 1, -1
				parent = node_left_rotate(parent)
			case -1:
				if parent.left == nil {
					parent.balance, son.balance = 0, 0
				} else {
					subtree = son.left
					switch subtree.balance {
					case -1:
						parent.balance, son.balance, subtree.balance = 0, 1, 0
					case 0:
						parent.balance, son.balance, subtree.balance = 0, 0, 0
					case 1:
						parent.balance, son.balance, subtree.balance = -1, 0, 0
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
			n, parent = parent, parent.parent
			switch {
			case parent == nil:
			case parent.left == n:
				parent.balance += 1
			case parent.right == n:
				parent.balance -= 1
			}
		}
	}

	if parent == nil {
		t.Root = n
	}
	if Log_tree {
		t.PrintFile("tree_result.jpg")
	}
}

func node_right_rotate(n *node) *node {
	parent := n.parent
	son := n.left
	subtree := son.right
	if parent != nil {
		switch {
		case parent.left == n:
			parent.left = son
		case parent.right == n:
			parent.right = son
		}
	}
	son.parent = n.parent
	son.right = n
	n.parent = son
	n.left = subtree
	if subtree != nil {
		subtree.parent = n
	}
	return son
}

func node_left_rotate(n *node) *node {
	parent := n.parent
	son := n.right
	subtree := son.left
	if parent != nil {
		switch {
		case parent.left == n:
			parent.left = son
		case parent.right == n:
			parent.right = son
		}
	}
	son.parent = n.parent
	son.left = n
	n.parent = son
	n.right = subtree
	if subtree != nil {
		subtree.parent = n
	}
	return son
}

func next_node(n *node) (next *node) {
	next = n.right
	for next.left != nil {
		next = next.left
	}
	return
}

func (t *Tree) Print(w io.Writer) {
	if t.Root == nil {
		return
	}

	n := t.Root
	var doc bytes.Buffer

	doc.WriteString("digraph AvlTree {")
	color := "black"
	if n.balance == -2 || n.balance == 2 {
		color = "red"
	}
	doc.WriteString(fmt.Sprintf("%v [shape=circle, xlabel=%v, color=%v];", n.id, n.balance, color))
	n.write(&doc)
	doc.WriteString("}")

	w.Write(doc.Bytes())
}

func (n *node) write(doc *bytes.Buffer) {
	color := "green"
	if n.balance == -2 || n.balance == 2 {
		color = "black"
	}

	if n.left != nil {
		if n.left.balance == -2 || n.left.balance == 2 {
			color = "red"
		}
		pid := "n"
		if n.left.parent != nil {
			pid = fmt.Sprintf("%v", n.left.parent.id)
		}
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v(%v)", color=%v];`, n.left.id, pid, n.left.balance, color))
		doc.WriteString(fmt.Sprintf("%v -> %v [color=black];", n.id, n.left.id))
	} else {
		doc.WriteString(fmt.Sprintf("nl%v [shape=point];", n.id))
		doc.WriteString(fmt.Sprintf("%v -> nl%v [color=blue];", n.id, n.id))
	}
	if n.right != nil {
		if n.right.balance == -2 || n.right.balance == 2 {
			color = "red"
		}
		pid := "n"
		if n.right.parent != nil {
			pid = fmt.Sprintf("%v", n.right.parent.id)
		}
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v(%v)", color=%v];`, n.right.id, pid, n.right.balance, color))
		doc.WriteString(fmt.Sprintf("%v -> %v [color=blue];", n.id, n.right.id))
	} else {
		doc.WriteString(fmt.Sprintf("nr%v [shape=point];", n.id))
		doc.WriteString(fmt.Sprintf("%v -> nr%v [color=blue];", n.id, n.id))
	}

	if n.left != nil {
		n.left.write(doc)
	}
	if n.right != nil {
		n.right.write(doc)
	}
}

func (t *Tree) PrintFile(path string) {
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
