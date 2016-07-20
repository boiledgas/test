package main

import (
	"net/http"
	_ "net/http/pprof"
	"test/test1"
	"test/tree/splay"
)

func test1Handler(w http.ResponseWriter, r *http.Request) {
	test1.WritePng(w)
}

func test2Handler(w http.ResponseWriter, r *http.Request) {
	data := []byte("<html><head><title>Hello world</title></head><body><h1>Hello World!</h1></body></html>")
	w.Write(data)
}

func main() {
	t1 := new(splay.Tree)
	t1.Insert(7, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(3, nil)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	t1.Insert(4, nil)

	if _, ok := t1.Find(6); !ok {
	}
	if err := t1.Validate(); err != nil {
		t1.PrintFile("img\\find-zig.jpg")
	}
}

//
//func main() {
//	len := 1000000
//	data := make([]int32, len)
//	added := list.New()
//	test_tree := new(tree.Avl_tree)
//	var item int32
//	for i := 0; i < len; i ++ {
//		item = rand.Int31n(len)
//		data[i] = item
//		added.PushBack(item)
//		test_tree.Insert(item)
//	}
//
//	var index int32
//	var node *list.Element
//	removed := list.New()
//	for i := 0; i < len; i ++ {
//		node = added.Front()
//		index = rand.Int31n(int32(added.Len()))
//		for j := 0; int32(j) < index; j ++ {
//			node = node.Next()
//		}
//		test_tree.Delete(node.Value.(int32))
//		added.Remove(node)
//		if err := test_tree.Validate(); err != nil {
//			log.Println(err)
//			log.Println(node.Value.(int32))
//
//			bad_tree := new(tree.Avl_tree)
//			for j := 0; j < len; j ++ {
//				bad_tree.Insert(data[j])
//			}
//			for remove_node := removed.Front(); remove_node != nil; remove_node = remove_node.Next() {
//				bad_tree.Delete(remove_node.Value.(int32))
//			}
//			tree.Log_tree = true
//			bad_tree.Delete(node.Value.(int32))
//			return
//		}
//		removed.PushBack(node.Value.(int32))
//	}
//}
