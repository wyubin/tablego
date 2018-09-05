package tablego

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestTsvParse(t *testing.T) {
	fTsv := bytes.NewBufferString("#col0\tcol1\tvol2\n0,0\t0,1\t0,2\n1,0\t1,1\t1,2\n")
	tabO := Input(fTsv)
	lineC := 0
	for dA := range tabO.Iter() {
		fmt.Println(strings.Join(dA, "\t"))
		if len(dA) != 3 {
			t.Errorf("line %v do not parse into array", lineC)
		}
		lineC++
	}
	if tabO.Colnames == nil {
		t.Error("Should have colnames!")
	}
}

func BenchmarkTsvParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fTsv := bytes.NewBufferString("0,0\t0,1\t0,2\n1,0\t1,1\t1,2\n")
		tabO := Input(fTsv)
		tmpI := 0
		for dA := range tabO.Iter() {
			tmpI += len(dA)
		}
	}
}
