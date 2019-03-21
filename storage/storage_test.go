package storage

import (
	"log"
	"testing"
)

var tt [][]string = [][]string{
	{"get", "liufei"},
	{"set", "liufei", "liufei"},
	{"get", "liufei"},
	{"del", "liufei"},
	{"get", "liufei"},
	{"lpush", "liufeil","nihao"},
	{"lpush","liufeil","test"},
	{"lpop", "liufeil"},
	{"lpop", "liufeil"},
	{"lpop", "liufeil"},
}

func TestDojob(t *testing.T) {
	for _, s := range tt {
		r, err := GetResult(s...)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(r,s)
		}
	}
}
func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range tt {
			_,_ =GetResult(s...)
		}
	}
}
