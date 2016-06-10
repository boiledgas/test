package main

type b_tree_node struct {
	keys  []int
	nodes []*b_tree_node
}

type b_tree struct {
	root *b_tree_node
	t    byte
}

func (t *b_tree) Find(id int) (res interface{}, ok bool) {
	for {

	}
}
