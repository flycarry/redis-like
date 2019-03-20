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
}

func TestDojob(t *testing.T) {
	for _, s := range tt {
		r, err := GetResult(s...)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(r)
		}
	}
}
func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range tt {
			r, err := GetResult(s...)
			if err != nil {
				log.Println(err)
			} else {
				log.Println(r)
			}
		}
	}
}
