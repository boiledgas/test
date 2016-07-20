package treap

import (
	"bytes"
	"fmt"
	"io"
	"errors"
)

var Log_tree bool = false
var i = 0

type node struct {
	id     int32
	value  interface{}
	parent *node
	left   *node
	right  *node

	priority uint32
}

func (n *node) validate() error {
	if n == nil {
		return nil
	}

	if n.left != nil && n.left.id >= n.id || n.right != nil && n.right.id <= n.id {
		return errors.New(fmt.Sprintf("%v node structure not valid", n))
	}
	if n.left != nil && (n.left.parent == nil || n.left.parent.id != n.id) {
		return errors.New(fmt.Sprintf("%v left parent wrong right(%v)-left(%v)", n, n.right, n.left))
	}
	if n.right != nil && (n.right.parent == nil || n.right.parent.id != n.id) {
		return errors.New(fmt.Sprintf("%v right parent wrong right(%v)-left(%v)", n, n.right, n.left))
	}

	panic("add priority check")

	return nil
}

func (n *node) String() string {
	if n == nil {
		return ""
	}

	var b bytes.Buffer
	b.WriteString("{")
	if n.left != nil {
		b.WriteString(fmt.Sprintf("%v:", n.left.id))
		if n.left.parent != nil {
			b.WriteString(fmt.Sprintf("%v", n.left.parent.id))
		} else {
			b.WriteString("n")
		}
	} else {
		b.WriteString("null ")
	}
	b.WriteString(" < ")
	b.WriteString(fmt.Sprintf("%v:", n.id))
	if n.parent != nil {
		b.WriteString(fmt.Sprintf("%v", n.parent.id))
	} else {
		b.WriteString("n")
	}
	b.WriteString(" > ")
	if n.right != nil {
		b.WriteString(fmt.Sprintf("%v:", n.right.id))
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

func (n *node) write(doc *bytes.Buffer) {
	color := "green"

	if n.left != nil {
		pid := "n"
		if n.left.parent != nil {
			pid = fmt.Sprintf("%v", n.left.parent.id)
		}
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v(%v)", color=%v];`, n.left.id, pid, n.left.priority, color))
		doc.WriteString(fmt.Sprintf("%v -> %v [color=black];", n.id, n.left.id))
	} else {
		doc.WriteString(fmt.Sprintf("nl%v [shape=point];", n.id))
		doc.WriteString(fmt.Sprintf("%v -> nl%v [color=blue];", n.id, n.id))
	}
	if n.right != nil {
		pid := "n"
		if n.right.parent != nil {
			pid = fmt.Sprintf("%v", n.right.parent.id)
		}
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v(%v)", color=%v];`, n.right.id, pid, n.right.priority, color))
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

type Tree struct {
	root *node
	len  int
}

func (t *Tree) Insert(int32, interface{}) {
	panic("not implemented")
}

func (t *Tree) Delete(int32) interface{} {
	panic("not implemented")
}

func (t *Tree) Find(id int32) (result interface{}, ok bool) {
	ok = false
	result = nil
	if Log_tree {
		t.PrintFile(fmt.Sprintf("img\\%v-%v-find-0.jpg", i, id))
	}
	if n := find(t.root, id); n.id == id {
		ok = true
		result = n.value
	}
	if Log_tree {
		i++
	}
	return
}

func (t *Tree) Count() int {
	return t.len
}

func (t *Tree) Min() (int32, interface{}) {
	if n := min(t.root); n != nil {
		return n.id, n.value
	}

	return -1, nil
}

func (t *Tree) Max() (int32, interface{}) {
	if n := max(t.root); n != nil {
		return n.id, n.value
	}

	return -1, nil
}

func (t *Tree) Asc(func(int32)) {

}

func (t *Tree) Desc(func(int32)) {

}

func (t *Tree) Validate() error {
	return t.root.validate()
}

func (t *Tree) Print(w io.Writer) {
	if t.root == nil {
		return
	}

	n := t.root
	var doc bytes.Buffer

	doc.WriteString("digraph AvlTree {")
	color := "black"
	doc.WriteString(fmt.Sprintf("%v [shape=circle, xlabel=%v, color=%v];", n.id, n.priority, color))
	n.write(&doc)
	doc.WriteString("}")

	w.Write(doc.Bytes())
}

func (t *Tree) PrintFile(string) {

}

// общие функции
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

func min(root *node) *node {
	if root == nil {
		return nil
	}
	n := root
	for {
		if n.left == nil {
			break
		}
		n = n.left
	}
	return n
}

func max(root *node) *node {
	if root == nil {
		return nil
	}
	n := root
	for {
		if n.right == nil {
			break
		}
		n = n.right
	}
	return n
}
