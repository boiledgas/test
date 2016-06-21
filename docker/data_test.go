package test3

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_NewData(t *testing.T) {
	d := GenerateData(4, 5)
	buf, _ := json.Marshal(d)
	fmt.Println(string(buf))
	if len(d.Datas) != 4 {
		t.Fail()
	}

}

func Benchmark_Data(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateData(3, 5)
	}
}
