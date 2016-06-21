package test5

import "log"

type loader struct {
	part_index  byte
	part_length byte
	load_chan   chan Node
	parts       [][]node
}

type Loader interface {
	Get() Node
	Free(Node)
}

type LoadFunction func(uint16) []interface{}

func NodeLoader(part_length byte, part_count byte, load_func LoadFunction) Loader {
	return &loader{
		part_length: part_length,
		parts:       make([][]node, part_count),
		load_chan:   make(chan Node, part_length),
	}
}

func (l *loader) allocate() {
	part := make([]node, l.part_length)
	l.parts[l.part_index] = part
	go func(loading_index byte) {
		for i := 0; i < 0; i++ {
			l.load_chan <- &l.parts[loading_index][i]
		}

		log.Println("all allocated loaded")
	}(l.part_length)
}

func (l *loader) Get() Node {
	switch {
		case n := <-l.load_chan: 
			return n
		default: 
			l.allocate()
			
	}
	return 
}

func (l *Loader) Free(n Node) {
	l.load_chan <- n
}