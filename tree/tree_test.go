package tree

import (
	"math/rand"
	"testing"
	"os"
)

func init() {
	os.Mkdir("img", os.ModeDir)
}

func _Test_Avl_OnlyLeftLeft(t *testing.T) {
	tree := new(Avl_tree)
	for i := 15; i > 0; i-- {
		tree.Insert(i)
	}

	filename := "c:\\temp\\graph_left.jpg"
	tree.PrintFile(filename)
}

func _Test_Avl_OnlyRightRight(t *testing.T) {
	tree := new(Avl_tree)
	for i := 0; i < 15; i++ {
		tree.Insert(i)
	}

	filename := "c:\\temp\\graph_right.jpg"
	tree.PrintFile(filename)
}

// subtree.Balance == 0
func _Test_Avl_LeftRight1(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(25)
	tree.Insert(8)
	tree.Insert(16)
}

// subtree.Balance == -1 (left)
func _Test_Avl_LeftRight11(t *testing.T) {
	tree := new(Avl_tree)
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
	tree := new(Avl_tree)
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
	tree := new(Avl_tree)
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
	tree := new(Avl_tree)
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
	tree := new(Avl_tree)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(15)

	filename := "c:\\temp\\graph_rightleft.jpg"
	tree.PrintFile(filename)
}

// subtree.Balance == -1 (left)
func _Test_Avl_RightLeft11(t *testing.T) {
	tree := new(Avl_tree)
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
	tree := new(Avl_tree)
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
func _Test_Avl_RightLeft21(t *testing.T) {
	tree := new(Avl_tree)
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
	tree := new(Avl_tree)
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
	tree := new(Avl_tree)
	for i := 0; i < 100; i++ {
		tree.Insert(int(rand.Int31n(1000)))
	}

	filename := "c:\\temp\\graph_random.jpg"
	tree.PrintFile(filename)
}

func _Test_TreeFind(t *testing.T) {
	tree := new(Avl_tree)
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

func Test_Delete_NotFound(t *testing.T) {
	tree := new(Avl_tree)
	if err := tree.Delete(4); err != nil {
		if err.Error() != "id 4 not found" {
			t.Error("message not match")
		}
	} else {
		t.Error("delete not return error")
	}
	tree.Insert(3)
	if err := tree.Delete(4); err != nil {
		if err.Error() != "id 4 not found" {
			t.Error("message not match")
		}
	} else {
		t.Error("delete not return error")
	}
}

func Test_Delete_SingleRoot(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(1)
	tree.PrintFile("img\\delete(singleroot)-tree.jpg")
	tree.Delete(1)
	tree.PrintFile("img\\delete(singleroot)-result.jpg")
}

func Test_Delete_SingleLeft(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(2)
	tree.Insert(1)
	tree.PrintFile("img\\delete(singleleft)-tree.jpg")
	tree.Delete(2)
	tree.PrintFile("img\\delete(singleleft)-result.jpg")
}

func Test_Delete_SingleRight(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(1)
	tree.Insert(2)
	tree.PrintFile("img\\delete(singleright)-tree.jpg")
	tree.Delete(1)
	tree.PrintFile("img\\delete(singleright)-result.jpg")
}

func Test_Delete_Root(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(10)
	tree.PrintFile("img\\delete(root)-tree.jpg")
	tree.Delete(5)
	tree.PrintFile("img\\delete(root)-result.jpg")
}

func Test_DeleteLeaf_n21nil(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(2)
	tree.PrintFile("img\\delete(-21rnil)-leaf-tree.jpg")
	tree.Delete(4)
	tree.PrintFile("img\\delete(-21rnil)-leaf-result.jpg")
}

func Test_DeleteLeaf_n21n1(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(2)
	tree.Insert(3)
	tree.PrintFile("img\\delete(-21-1)-leaf-tree.jpg")
	tree.Delete(8)
	tree.PrintFile("img\\delete(-21-1)-leaf-result.jpg")
}

func Test_DeleteLeaf_n210(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(5)
	tree.PrintFile("img\\delete(-210)-leaf-tree.jpg")
	tree.Delete(8)
	tree.PrintFile("img\\delete(-210)-leaf-result.jpg")
}

func Test_DeleteLeaf_n211(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(2)
	tree.Insert(5)
	tree.PrintFile("img\\delete(-211)-leaf-tree.jpg")
	tree.Delete(8)
	tree.PrintFile("img\\delete(-211)-leaf-result.jpg")
}

func Test_DeleteLeaf_n20(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(8)
	tree.Insert(4)
	tree.Insert(9)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)
	tree.PrintFile("img\\delete(-20)-leaf-tree.jpg")
	tree.Delete(10)
	tree.PrintFile("img\\delete(-20)-leaf-result.jpg")
}

func Test_DeleteLeaf_n2n1(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(8)
	tree.Insert(4)
	tree.Insert(9)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(3)
	tree.PrintFile("img\\delete(-2-1)-leaf-tree.jpg")
	tree.Delete(10)
	tree.PrintFile("img\\delete(-2-1)-leaf-result.jpg")
}

func Test_DeleteLeaf_2n1nil(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(3)
	tree.PrintFile("img\\delete(2-1nil)-leaf-tree.jpg")
	tree.Delete(1)
	tree.PrintFile("img\\delete(2-1nil)-leaf-result.jpg")
}

func Test_DeleteLeaf_2n1n1(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(4)
	tree.PrintFile("img\\delete(2-1-1)-leaf-tree.jpg")
	tree.Delete(1)
	tree.PrintFile("img\\delete(2-1-1)-leaf-result.jpg")
}

func Test_DeleteLeaf_2n10(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.PrintFile("img\\delete(2-10)-leaf-tree.jpg")
	tree.Delete(1)
	tree.PrintFile("img\\delete(2-10)-leaf-result.jpg")
}

func Test_DeleteLeaf_2n11(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(6)
	tree.PrintFile("img\\delete(2-11)-leaf-tree.jpg")
	tree.Delete(1)
	tree.PrintFile("img\\delete(2-11)-leaf-result.jpg")
}

func Test_DeleteLeaf_20(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(10)
	tree.PrintFile("img\\delete(20)-leaf-tree.jpg")
	tree.Delete(1)
	tree.PrintFile("img\\delete(20)-leaf-result.jpg")
}

func Test_DeleteLeaf_21(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(8)
	tree.Insert(10)
	tree.PrintFile("img\\delete(21)-leaf-tree.jpg")
	tree.Delete(1)
	tree.PrintFile("img\\delete(21)-leaf-result.jpg")
}

func Test_DeleteSubtree_n20(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(8)
	tree.Insert(4)
	tree.Insert(9)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)
	tree.PrintFile("img\\delete(-20)-left-tree.jpg")
	tree.Delete(9)
	tree.PrintFile("img\\delete(-20)-left-result.jpg")
}

func Test_DeleteSubtree_20(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(10)
	tree.PrintFile("img\\delete(20)-subleft-tree.jpg")
	tree.Delete(2)
	tree.PrintFile("img\\delete(20)-subleft-result.jpg")
}

func Test_DeleteSubtree_n2n1(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(8)
	tree.Insert(4)
	tree.Insert(9)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(3)
	tree.PrintFile("img\\delete(-2-1)-subright-tree.jpg")
	tree.Delete(9)
	tree.PrintFile("img\\delete(-2-1)-subright-result.jpg")
}

func Test_Delete_RightWithLeftLeaf(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(3)
	tree.PrintFile("img\\delete(right_with_left)-tree.jpg")
	tree.Delete(4)
	tree.PrintFile("img\\delete(right_with_left)-result.jpg")
}

func Test_Delete_LeftWithRightLeaf(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(2)
	tree.PrintFile("img\\delete(right_with_left)-tree.jpg")
	tree.Delete(1)
	tree.PrintFile("img\\delete(right_with_left)-result.jpg")
}

func Test_Delete_RightRoot(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(11)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(9)
	tree.Insert(12)
	tree.Insert(10)

	tree.PrintFile("img\\delete(rightroot)-tree.jpg")
	tree.Delete(8)
	tree.PrintFile("img\\delete(rightroot)-result.jpg")
}

func Test_Delete_RightRootNode(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(8)

	tree.PrintFile("img\\delete(rightrootnode)-tree.jpg")
	tree.Delete(6)
	tree.PrintFile("img\\delete(rightrootnode)-result.jpg")
}