package tree

import (
	"io"
)

type Node interface {
	GetId() int
}

type Tree interface {
	Insert(int32, interface{})
	Delete(int32) interface{}
	Find(int32) (interface{}, bool)
	Count() int32
	Min() interface{}
	Max() interface{}
	Asc(func(int32))
	Desc(func(int32))
	Validate() (bool, error)
	Print(io.Writer)
	PrintFile(string)
}
