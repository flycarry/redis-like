package storage

import (
	"log"
	"strconv"
	"testing"
)

var m *myMap = newMap()

type keyvalue struct {
	key, value string
}

var kv []keyvalue = []keyvalue{
	keyvalue{"liufei", "nihao"},
	keyvalue{"liufei", "test"},
	keyvalue{"hello", "nihao"},
	keyvalue{"bushi", "nihao"},
}
var kv2 []keyvalue = []keyvalue{}
var mm map[string]string = make(map[string]string)

func init() {
	for i := 0; i < 100; i++ {
		kv2 = append(kv2, keyvalue{strconv.Itoa(i), strconv.Itoa(i)})
	}
}
func TestMapInit(t *testing.T) {
	for _, keyval := range kv2 {
		m.set(keyval.key, keyval.value)
		result, err := m.get(keyval.key)
		if err != nil {
			log.Fatal(keyval.key, keyval.value, result, "not equal")
		}
	}

}

func BenchmarkMyMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keyval := kv2[b.N%len(kv2)]
		m.set(keyval.key, keyval.value)
		result, err := m.get(keyval.key)
		if err != nil {
			log.Fatal(keyval.key, keyval.value, result, "not equal")
		}
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keyval := kv2[b.N%len(kv2)]
		mm[keyval.key] = keyval.value
		result, ok := mm[keyval.key]
		if ok != true {
			log.Fatal(keyval.key, keyval.value, result, "not equal")
		}
	}
}
