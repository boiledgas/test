package test2

import (
	"container/list"
	"encoding/json"
	"errors"
)

type TreeBuilder interface {
	Parse(string) error // разбор json-строки

	Node1(string, bool) byte       // добавление к дереву узла
	Node2(string, string) byte     // добавление к дереву узла
	Aggregate(string, []byte) byte // добавление к дереву аггрегата
	Root(byte)                     // указание главного узла дерева

	Build() (Tree, error) // построение дерева
}

type tree_builder struct {
	new_id byte
	nodes  *list.List
	root   byte
}

func NewTreeBuilder() TreeBuilder {
	b := &tree_builder{}
	b.reset()
	return b
}

func (b *tree_builder) Parse(res string) (err error) {
	var m interface{}
	if err = json.Unmarshal([]byte(res), &m); err != nil {
		return
	}

	node_map := m.(map[string]interface{})
	root_id := b.parse_node(node_map)
	b.Root(root_id)

	return
}

func (b *tree_builder) Node1(name string, enabled bool) byte {
	n := node1{node{b.new_id, name}, enabled}
	b.push_back(&n)
	return n.id
}

func (b *tree_builder) Node2(name string, color string) byte {
	n := node2{node{b.new_id, name}, color}
	b.push_back(&n)
	return n.id
}

func (b *tree_builder) Aggregate(name string, ids []byte) byte {
	n := node_aggregate{node{b.new_id, name}, ids}
	b.push_back(&n)
	return n.id
}

func (b *tree_builder) Root(id byte) {
	b.root = id
}

func (b *tree_builder) Build() (t Tree, err error) {
	if b.root == 255 {
		err = errors.New("root node not found")
		return
	}

	t, err = b.build_tree_array_interface()
	b.reset()

	return
}

func (b *tree_builder) build_tree_array_interface() (t Tree, err error) {
	var tree = &tree{make([]Node, b.nodes.Len()), b.root}
	for e := b.nodes.Front(); e != nil; e = e.Next() {
		node := e.Value.(Node)
		tree.Nodes[node.GetId()] = node
	}

	t = tree
	return
}

func (b *tree_builder) reset() {
	b.nodes = list.New()
	b.new_id = 0
	b.root = 255
}

func (b *tree_builder) parse_node(obj map[string]interface{}) byte {
	name := obj["name"].(string)
	node_type := byte(obj["node_type"].(float64))
	switch node_type {
	case NT_NODE_1:
		enabled := obj["enabled"].(bool)
		return b.Node1(name, enabled)
	case NT_NODE_2:
		color := obj["color"].(string)
		return b.Node2(name, color)
	case NT_NODE_AGGREGATE:
		childs := obj["childs"].([]interface{})
		nodeIds := make([]byte, len(childs))
		for i, child := range childs {
			node_map := child.(map[string]interface{})
			nodeId := b.parse_node(node_map)
			nodeIds[i] = nodeId
		}
		return b.Aggregate(name, nodeIds)
	default:
		panic("not found")
	}
}

func (b *tree_builder) push_back(n Node) {
	b.nodes.PushBack(n)
	b.new_id++
}
