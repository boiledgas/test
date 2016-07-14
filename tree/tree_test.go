package tree

import (
	"os"
	"testing"
)

func init() {
	os.Mkdir("img", os.ModeDir)
}

func TestAvl_Insert_Exist(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(1)
	tree.Insert(1)
	if tree.Count() != 1 {
		t.Error()
	}

	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(exists).jpg")
	}
}

func TestAvl_Insert_OnlyLeftLeft(t *testing.T) {
	tree := new(Avl_tree)
	for i := 15; i > 0; i-- {
		tree.Insert(int32(i))
	}
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(ll).jpg")
	}
}

func TestAvl_Insert_OnlyRightRight(t *testing.T) {
	tree := new(Avl_tree)
	for i := 0; i < 15; i++ {
		tree.Insert(int32(i))
	}
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(rr).jpg")
	}
}

func TestAvl_Insert_n2n1nil(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(1)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(-2-1nil).jpg")
	}
}

func TestAvl_Insert_n2n1(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(1)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(-2-1).jpg")
	}
}

func TestAvl_Insert_n21nil(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(25)
	tree.Insert(8)
	tree.Insert(16)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(-21nil).jpg")
	}
}

func TestAvl_Insert_n21n1(t *testing.T) {
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
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(-21-1).jpg")
	}
}

func TestAvl_Insert_n211(t *testing.T) {
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
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(-211).jpg")
	}
}

func TestAvl_Insert_2n1nil(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(15)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(2-1nil).jpg")
	}
}

func TestAvl_Insert_2n1n1(t *testing.T) {
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
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(2-1-1).jpg")
	}
}

func TestAvl_Insert_2n11(t *testing.T) {
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
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(2-11).jpg")
	}
}

func TestAvl_Insert_21nil(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(6)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(21nil).jpg")
	}
}

func TestAvl_Insert_21(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(5)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\insert(21).jpg")
	}
}

func TestAvl_Delete_NotFound(t *testing.T) {
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
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(notfound).jpg")
	}
}

func TestAvl_Delete_SingleRoot(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(1)
	tree.Delete(1)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(singleroot).jpg")
	}
}

func TestAvl_Delete_SingleLeft(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(2)
	tree.Insert(1)
	tree.Delete(2)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(singleleft).jpg")
	}
}

func TestAvl_Delete_SingleRight(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(1)
	tree.Insert(2)
	tree.Delete(1)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(singleright).jpg")
	}
}

func TestAvl_Delete_Root(t *testing.T) {
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
	tree.Delete(5)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(root).jpg")
	}
}

func TestAvl_Delete_n21nil(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(2)
	tree.Delete(4)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(-21nil).jpg")
	}
}

func TestAvl_Delete_n21n1(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(2)
	tree.Insert(3)
	tree.Delete(8)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(-21-1).jpg")
	}
}

func TestAvl_Delete_n210(t *testing.T) {
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
	tree.Delete(8)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(-210).jpg")
	}
}

func TestAvl_Delete_n211(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(2)
	tree.Insert(5)
	tree.Delete(8)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(-211).jpg")
	}
}

func TestAvl_Delete_n20(t *testing.T) {
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
	tree.Delete(10)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(-20).jpg")
	}
}

func TestAvl_Delete_n2n1(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(8)
	tree.Insert(4)
	tree.Insert(9)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(3)
	tree.Delete(10)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(-2-1).jpg")
	}
}

func TestAvl_Delete_2n1nil(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(3)
	tree.Delete(1)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(2-1nil).jpg")
	}
}

func TestAvl_Delete_2n1n1(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(4)
	tree.Delete(1)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(2-1-1).jpg")
	}
}

func TestAvl_Delete_2n10(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Delete(1)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(2-10).jpg")
	}
}

func TestAvl_Delete_2n11(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(6)
	tree.Delete(1)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(2-11).jpg")
	}
}

func TestAvl_Delete_20(t *testing.T) {
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
	tree.Delete(1)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(20).jpg")
	}
}

func TestAvl_Delete_21(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(8)
	tree.Insert(10)
	tree.Delete(1)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(21).jpg")
	}
}

func TestAvl_Delete_Subtree_n20(t *testing.T) {
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
	tree.Delete(9)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(subtree-20).jpg")
	}
}

func TestAvl_Delete_Subtree_20(t *testing.T) {
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
	tree.Delete(2)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(subtree20).jpg")
	}
}

func TestAvl_Delete_Subtree_n2n1(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(8)
	tree.Insert(4)
	tree.Insert(9)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(3)
	tree.Delete(9)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(subtree-2-1).jpg")
	}
}

func TestAvl_Delete_RightWithLeftLeaf(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(3)
	tree.Delete(4)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(rightwithleaf).jpg")
	}
}

func TestAvl_Delete_LeftWithRightLeaf(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(2)
	tree.Delete(1)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(leftwithrightleaf).jpg")
	}
}

func TestAvl_Delete_RightRoot(t *testing.T) {
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
	tree.Delete(8)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(rightroot).jpg")
	}
}

func TestAvl_Delete_RightRootNode(t *testing.T) {
	tree := new(Avl_tree)
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(8)
	tree.Delete(6)
	if err := tree.Validate(); err != nil {
		t.Error(err)
		tree.PrintFile("img\\delete(rightrootnode).jpg")
	}
}
