package test2

type tree struct {
	Nodes []Node
	Root  byte
}

func NewTree(root byte) Tree {
	return &tree{make([]Node, 10), root}
}

func (t *tree) GetRoot() Node {
	return t.Nodes[t.Root]
}

func (t *tree) GetNode(id byte) Node {
	return t.Nodes[id]
}
