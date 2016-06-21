package test2

// полиморфизм на основе интерфейса
type node struct {
	id   byte
	name string
}

type node1 struct {
	node
	enabled bool
}

type node2 struct {
	node
	color string
}

type node_aggregate struct {
	node
	childs []byte
}

func (n *node) GetId() byte {
	return n.id
}

func (n *node) GetName() string {
	return n.name
}
