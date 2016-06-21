package test5

import "fmt"

type node struct {
	Id          uint64
	Name        [80]byte
	Description [255]byte
	Type        byte
}

type Node interface {
	String() string
}

func (n *node) String() string {
	return fmt.Sprintf("%v %v %v %v", n.Id, n.Name, n.Description, n.Type)
}
