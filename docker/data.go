package test3

import "fmt"

type Data struct {
	Id    uint64
	Name  []byte
	Buf   [0]byte
	Datas []*Data `_,omitempty`
}

var i uint64

func GenerateData(l uint8, c byte) (d *Data) {
	if l == 0 {
		return nil
	}

	d = &Data{Id: i, Name: []byte(fmt.Sprintf("item %v", i))}
	i++
	var child *Data
	datas := make([]*Data, c)
	for i := 0; i < int(c); i++ {
		child = GenerateData(l-1, c)
		if child == nil {
			return
		}

		datas[i] = child
	}

	d.Datas = datas
	return
}
