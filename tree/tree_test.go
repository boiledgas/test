package main

import (
	"math/rand"
	"testing"
)

func _Test_Avl_OnlyLeftLeft(t *testing.T) {
	tree := NewAvlTree()
	for i := 15; i > 0; i-- {
		tree.Insert(i)
	}

	filename := "c:\\temp\\graph_left.jpg"
	tree.PrintFile(filename)
}

func _Test_Avl_OnlyRightRight(t *testing.T) {
	tree := NewAvlTree()
	for i := 0; i < 15; i++ {
		tree.Insert(i)
	}

	filename := "c:\\temp\\graph_right.jpg"
	tree.PrintFile(filename)
}

// subtree.Balance == 0
func _Test_Avl_LeftRight1(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(25)
	tree.Insert(8)
	tree.Insert(16)
}

// subtree.Balance == -1 (left)
func _Test_Avl_LeftRight11(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(16)

	tree.Insert(8)
	tree.Insert(25)

	tree.Insert(3)
	tree.Insert(12)
	tree.Insert(23)
	tree.Insert(27)

	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(10)
	tree.Insert(14)

	tree.Insert(9)

	filename := "c:\\temp\\graph_leftright.jpg"
	tree.PrintFile(filename)
}

// subtree.Balance == -1 (right)
func _Test_Avl_LeftRight12(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(16)
	tree.Insert(8)
	tree.Insert(25)
	tree.Insert(3)
	tree.Insert(12)
	tree.Insert(23)
	tree.Insert(27)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(10)
	tree.Insert(14)

	tree.Insert(11)

	filename := "c:\\temp\\graph_leftright.jpg"
	tree.PrintFile(filename)
}

// subtree.Balance == 1 (left)
func _Test_Avl_LeftRight21(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(16)
	tree.Insert(8)
	tree.Insert(25)
	tree.Insert(3)
	tree.Insert(12)
	tree.Insert(23)
	tree.Insert(27)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(10)
	tree.Insert(14)

	tree.Insert(13)

	filename := "c:\\temp\\graph_leftright.jpg"
	tree.PrintFile(filename)
}

// subtree.Balance == 1 (right)
func _Test_Avl_LeftRight22(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(16)
	tree.Insert(8)
	tree.Insert(25)
	tree.Insert(3)
	tree.Insert(12)
	tree.Insert(23)
	tree.Insert(27)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(10)
	tree.Insert(14)

	tree.Insert(15)

	filename := "c:\\temp\\graph_leftright.jpg"
	tree.PrintFile(filename)
}

// subtree.Balance == 0
func _Test_Avl_RightLeft1(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(15)

	filename := "c:\\temp\\graph_rightleft.jpg"
	tree.PrintFile(filename)
}

// subtree.Balance == -1 (left)
func Test_Avl_RightLeft11(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(20)

	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(15)
	tree.Insert(25)

	tree.Insert(12)
	tree.Insert(17)
	tree.Insert(23)
	tree.Insert(27)

	tree.Insert(11)

	filename := "c:\\temp\\graph_rightleft.jpg"
	tree.PrintFile(filename)
}

// subtree.Balance == -1 (right)
func _Test_Avl_RightLeft12(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(20)

	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(15)
	tree.Insert(25)

	tree.Insert(12)
	tree.Insert(17)
	tree.Insert(23)
	tree.Insert(27)

	tree.Insert(13)

	filename := "c:\\temp\\graph_rightleft.jpg"
	tree.PrintFile(filename)
}

// subtree.Balance == 1 (left)
func Test_Avl_RightLeft21(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(20)

	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(15)
	tree.Insert(25)

	tree.Insert(12)
	tree.Insert(17)
	tree.Insert(23)
	tree.Insert(27)

	tree.Insert(16)

	filename := "c:\\temp\\graph_rightleft.jpg"
	tree.PrintFile(filename)
}

// subtree.Balance == 1 (right)
func _Test_Avl_RightLeft22(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(20)

	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(15)
	tree.Insert(25)

	tree.Insert(12)
	tree.Insert(17)
	tree.Insert(23)
	tree.Insert(27)

	tree.Insert(18)

	filename := "c:\\temp\\graph_rightleft.jpg"
	tree.PrintFile(filename)
}

func _Test_Avl_Random(t *testing.T) {
	tree := NewAvlTree()
	for i := 0; i < 100; i++ {
		tree.Insert(int(rand.Int31n(1000)))
	}

	filename := "c:\\temp\\graph_random.jpg"
	tree.PrintFile(filename)
}

func _Test_TreeFind(t *testing.T) {
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
