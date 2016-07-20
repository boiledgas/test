package splay

import (
	"testing"
)

func Test_Insert_Asc(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	t1.Insert(3, nil)
	t1.Insert(4, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(7, nil)
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\insert-asc.jpg")
	}
}

func Test_Insert_Desc(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Insert(6, nil)
	t1.Insert(5, nil)
	t1.Insert(4, nil)
	t1.Insert(3, nil)
	t1.Insert(2, nil)
	t1.Insert(1, nil)
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\insert-asc.jpg")
	}
}

func Test_Insert_AscSwap(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(3, nil)
	t1.Insert(4, nil)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\insert-ascswap.jpg")
	}
}

func Test_Insert_Exists(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	if err := t1.Insert(7, nil); err == nil {
		t.Error(err)
	}
}

func Test_Delete_Mergelnilr(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(3, nil)
	t1.Insert(4, nil)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	t1.Delete(1)
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\delete-lnilr.jpg")
	}
}

func Test_Delete_Mergelrnil(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(3, nil)
	t1.Insert(4, nil)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	t1.Delete(7)
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\delete-lrnil.jpg")
	}
}

func Test_Delete_Mergelr(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(3, nil)
	t1.Insert(4, nil)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	t1.Delete(4)
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\delete-lr.jpg")
	}
}

func Test_Delete_Root(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Delete(7)
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\delete-root.jpg")
	}
}

func Test_Delete_NotExist(t *testing.T) {
	t1 := new(Tree)
	if err := t1.Delete(1); err == nil {
		t.Error("delete not exist")
	}
}

func Test_Find_Zag(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(3, nil)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	t1.Insert(4, nil)

	if _, ok := t1.Find(6); !ok {
		t.Error("not found")
	}
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\find-zag.jpg")
	}
}

func Test_Find_ZagZag(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(3, nil)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	t1.Insert(4, nil)

	if _, ok := t1.Find(7); !ok {
		t.Error("not found")
	}
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\find-zagzag.jpg")
	}
}

func Test_Find_ZagZig(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(3, nil)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	t1.Insert(4, nil)

	if _, ok := t1.Find(5); !ok {
		t.Error("not found")
	}
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\find-zagzig.jpg")
	}
}

func Test_Find_Zig(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(3, nil)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	t1.Insert(4, nil)

	if _, ok := t1.Find(2); !ok {
		t.Error("not found")
	}
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\find-zig.jpg")
	}
}

func Test_Find_ZigZig(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(3, nil)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	t1.Insert(4, nil)

	if _, ok := t1.Find(1); !ok {
		t.Error("not found")
	}
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\find-zigzig.jpg")
	}
}

func Test_Find_ZigZag(t *testing.T) {
	t1 := new(Tree)
	t1.Insert(7, nil)
	t1.Insert(5, nil)
	t1.Insert(6, nil)
	t1.Insert(3, nil)
	t1.Insert(1, nil)
	t1.Insert(2, nil)
	t1.Insert(4, nil)

	if _, ok := t1.Find(3); !ok {
		t.Error("not found")
	}
	if err := t1.Validate(); err != nil {
		t.Error(err)
		t1.PrintFile("img\\find-zigzag.jpg")
	}
}
