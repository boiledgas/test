package test2

import (
	"fmt"
	"testing"
)

type tree_stat struct {
	count_node1     byte
	count_node2     byte
	count_aggregate byte
}

var tree_count int = 100
var req string = "{\"name\":\"name1\",\"node_type\":3,\"childs\":[{\"name\":\"name21\",\"node_type\":2,\"color\":\"black\"},{\"name\":\"name22\",\"node_type\":1,\"enabled\":false},{\"name\":\"name23\",\"node_type\":3,\"childs\":[{\"name\":\"name31\",\"node_type\":3,\"childs\":[{\"name\":\"name41\",\"node_type\":2,\"color\":\"black\"},{\"name\":\"name42\",\"node_type\":1,\"enabled\":true}]},{\"name\":\"name32\",\"node_type\":2,\"color\":\"black\"},{\"name\":\"name33\",\"node_type\":1,\"enabled\":false}]},{\"name\":\"name24\",\"node_type\":2,\"color\":\"green\"},{\"name\":\"name25\",\"node_type\":1,\"enabled\":true}]}"
var test_stat tree_stat = tree_stat{4, 4, 3}

var trees []Tree

func init() {
	trees = make([]Tree, tree_count)
	builder := NewTreeBuilder()
	for i := 0; i < len(trees); i++ {
		builder.Parse(req)
		tree, _ := builder.Build()
		trees[i] = tree
	}
}

// парсинг входной строки
func Benchmark_Parse(b *testing.B) {
	builder := NewTreeBuilder()
	for i := 0; i < b.N; i++ {
		if err := builder.Parse(req); err != nil {
			b.Fail()
		}
		if _, err := builder.Build(); err != nil {
			b.Fail()
		}
	}
}

// скорость обхода дерева
func Benchmark_Access(b *testing.B) {
	l := len(trees)
	stat := tree_stat{}
	for i := 0; i < b.N; i++ {
		stat.count_node1 = 0
		stat.count_node2 = 0
		stat.count_aggregate = 0

		index := i % l
		tree := trees[index]
		root := tree.GetRoot()
		expandTree(tree, root, &stat)

		if stat != test_stat {
			panic(fmt.Sprintf("fail %v %v", test_stat, stat))
		}
	}
}

func expandTree(tree Tree, root Node, stat *tree_stat) {
	switch n := root.(type) {
	case *node1:
		stat.count_node1++
	case *node2:
		stat.count_node2++
	case *node_aggregate:
		stat.count_aggregate++
		for _, id := range n.childs {
			child := tree.GetNode(id)
			expandTree(tree, child, stat)
		}
	}
}
