package test2

type Tree interface {
	GetRoot() Node
	GetNode(byte) Node
}

type Node interface {
	GetId() byte
	GetName() string
}

const (
	NT_NODE_1         byte = 1
	NT_NODE_2         byte = 2
	NT_NODE_AGGREGATE byte = 3
)
