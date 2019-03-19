package storage

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s,err:=GetResult("set","liufei","niubi")
	if err != nil {
		t.Error(err)
	}else {
		fmt.Println(s)
	}
}

func TestDel(t *testing.T) {
	s,err:=GetResult("del","liufei")
	if err != nil {
		t.Fatal(err)
	}else {
		fmt.Println(s)
	}
}
func TestGet(t *testing.T) {
	s,err:=GetResult("get","liufei")
	if err != nil {
		fmt.Println(s)
	}else {
		fmt.Println(s)
	}
}
func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetResult("set","liufei","niubi")
		GetResult("get","liufei")
	}
}