package tree

import (
	"io"
)

type Node interface {
	GetId() int
}

type Tree interface {
	Insert(int32)
	Delete(int32)
	Find(int32) (interface{}, bool)
	Count() uint16
	Min() int32
	Max() int32
	Asc(func(int32))
	Desc(func(int32))
	Validate() (bool, error)
	Print(io.Writer)
	PrintFile(string)
}
