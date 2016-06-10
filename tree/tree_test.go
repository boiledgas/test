package main

import "testing"

func Test_TreeFind(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(20)
	tree.Insert(8)
	tree.Insert(15)
	tree.Insert(25)

	if tree.Count() != 7 {
		t.Fail()
	}
	if _, ok := tree.Find(17); ok {
		t.Error("found 17")
	}
	if _, ok := tree.Find(7); ok {
		t.Error("found 7")
	}

	tree.Insert(7)
	tree.Insert(17)
	if _, ok := tree.Find(7); !ok {
		t.Error("couldn't find 7")
	}
	if _, ok := tree.Find(17); !ok {
		t.Error("couldn't find 17")
	}

	if _, ok := tree.Find(10); !ok {
		t.Error("couldnt found 10")
	}
	tree.Delete(10)
	if _, ok := tree.Find(10); ok {
		t.Error("found 10")
	}

	if tree.Min() != 3 {
		t.Error("min not found")
	}

	if tree.Max() != 25 {
		t.Error("max not found")
	}

	asc_sum := 0
	tree.Asc(func(id int) {
		asc_sum += id
	})

	desc_sum := 0
	tree.Desc(func(id int) {
		desc_sum += id
	})

	if asc_sum != desc_sum || asc_sum != 100 {
		t.Error("sum not match")
	}
}
