package splay

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
var i int = 1

// вынести в общую структуру
type node struct {
	id     int32
	left   *node
	right  *node
	parent *node
	value  interface{}
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

type Tree struct {
	len  int32
	root *node
	min  *node
	max  *node
}

func (t *Tree) Insert(id int32, value interface{}) (err error) {
	n := find(t.root, id)
	if n == nil {
		t.root = &node{parent: n, id: id, value: value}
		t.min, t.max = t.root, t.root
		t.len++
		return
	}
	if Log_tree {
		t.PrintFile(fmt.Sprintf("img\\%v-%v-insert-0.jpg", i, id))
	}
	if n != nil && n.id == id {
		err = errors.New("key already exists")
	} else {
		node := &node{parent: n, id: id, value: value}
		switch {
		case n.id < id:
			n.right = node
			n = n.right
		case n.id > id:
			n.left = node
			n = n.left
		}
		t.len++
	}
	splay(n)
	t.root = n
	if Log_tree {
		t.PrintFile(fmt.Sprintf("img\\%v-%v-insert-1.jpg", i, id))
	}
	if Log_tree {
		i++
	}
	return
}

func (t *Tree) Delete(id int32) (err error) {
	n := find(t.root, id)
	if n == nil || n.id != id {
		return errors.New(fmt.Sprintf("not found %v", id))
	}
	if Log_tree {
		t.PrintFile(fmt.Sprintf("img\\%v-%v-delete-0.jpg", i, n.id))
	}
	splay(n)
	if Log_tree {
		t.root = n
		t.PrintFile(fmt.Sprintf("img\\%v-%v-delete-1.jpg", i, n.id))
	}
	t.root = merge(n.left, n.right)
	t.len--
	if Log_tree {
		t.PrintFile(fmt.Sprintf("img\\%v-%v-delete-2.jpg", i, n.id))
	}
	// delete n
	n.parent = nil
	if n.right != nil {
		n.right = nil
	}
	if n.left != nil {
		n.left = nil
	}
	if Log_tree {
		i++
	}
	return
}

func (t *Tree) Find(id int32) (result interface{}, ok bool) {
	ok = false
	result = nil
	if Log_tree {
		t.PrintFile(fmt.Sprintf("img\\%v-%v-find-0.jpg", i, id))
	}
	if n := find(t.root, id); n.id == id {
		splay(n)
		t.root = n
		ok = true
		result = n.value
		if Log_tree {
			t.PrintFile(fmt.Sprintf("img\\%v-%v-find-1.jpg", i, id))
		}
	}
	if Log_tree {
		i++
	}
	return
}

func (t *Tree) Count() int32 {
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

func (t *Tree) Validate() (err error) {
	err = t.root.validate()
	return
}

func splay(n *node) {
	p := n.parent
	gp := p
	for p != nil {
		gp = p.parent
		switch {
		case gp == nil && p.left == n: // zig
			right_rotate(p)
		case gp == nil && p.right == n: // zag
			left_rotate(p)
		case gp.left == p && p.left == n: // zig-zig
			right_rotate(p)
			right_rotate(gp)
		case gp.right == p && p.right == n: // zag-zag
			left_rotate(p)
			left_rotate(gp)
		case gp.left == p && p.right == n: // zig-zag
			left_rotate(p)
			right_rotate(gp)
		case gp.right == p && p.left == n: // zag-zig
			right_rotate(p)
			left_rotate(gp)
		}
		p = n.parent
	}
}

// все элементы left < right
func merge(left *node, right *node) (root *node) {
	switch {
	case left != nil && right != nil:
		max := max(left)
		left.parent = nil
		splay(max)
		max.right = right
		right.parent = max
		return max
	case left == nil && right != nil:
		return right
	case left != nil && right == nil:
		return left
	case left == nil && right == nil:
		return nil
	}
	return nil
}

func right_rotate(node *node) {
	if node.left == nil {
		return
	}
	parent := node.parent
	left := node.left

	left.parent = parent
	switch {
	case parent != nil && parent.left == node:
		parent.left = left
	case parent != nil && parent.right == node:
		parent.right = left
	}
	node.parent = node.left

	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	left.right = node
}

func left_rotate(node *node) {
	if node.right == nil {
		return
	}
	parent := node.parent
	right := node.right

	right.parent = parent
	switch {
	case parent != nil && parent.left == node:
		parent.left = right
	case parent != nil && parent.right == node:
		parent.right = right
	}
	node.parent = right

	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	right.left = node
}

func find(root *node, id int32) (node *node) {
	for {
		switch {
		case root == nil:
			return
		case root.id == id:
			node = root
			return
		case root.id > id:
			node = root
			root = root.left
		case root.id < id:
			node = root
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

func (t *Tree) Print(w io.Writer) {
	if t.root == nil {
		return
	}
	n := t.root
	var doc bytes.Buffer

	doc.WriteString("digraph SplayTree {")
	color := "black"
	doc.WriteString(fmt.Sprintf("%v [shape=circle, color=%v];", n.id, color))
	n.write(&doc)
	doc.WriteString("}")

	w.Write(doc.Bytes())
}

func (n *node) write(doc *bytes.Buffer) {
	color := "green"
	if n.left != nil {
		pid := "n"
		if n.left.parent != nil {
			pid = fmt.Sprintf("%v", n.left.parent.id)
		}
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v", color=%v];`, n.left.id, pid, color))
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
		doc.WriteString(fmt.Sprintf(`%v [shape=circle, xlabel="%v", color=%v];`, n.right.id, pid, color))
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

	return nil
}

// надо выпилить отсюда
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
