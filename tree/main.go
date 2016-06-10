package main

type Tree interface {
	Insert(int)
	Delete(int)
	Find(int) (interface{}, bool)
	Count() uint16
	Min() int
	Max() int
	Asc(func(int))
	Desc(func(int))
	Print()
}

func main() {

}
