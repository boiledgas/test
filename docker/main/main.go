package main

import (
	"container/list"
	"test/test3"
	"time"
)

func main() {
	save := false
	const l = 5024
	var buf [l]*test3.Data
	list := list.New()
	i := 0
	d := test3.GenerateData(5, 5)
	for {
		buf[i] = d
		time.Sleep(10)
		i++
		if i == l {
			i = 0
		}

		if save {
			list.PushBack(d)
		}
	}
}
