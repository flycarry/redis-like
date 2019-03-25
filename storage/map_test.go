package storage

import (
	"log"
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

func TestMapInit(t *testing.T) {
	for _, keyval := range kv {
		m.set(keyval.key, keyval.value)
		result, err := m.get(keyval.key)
		if err != nil {
			log.Fatal(keyval.key, keyval.value, result, "not equal")
		}
	}

}
func TestRight(t *testing.T) {
	var a uint32 =0
	a=^a
	var b byte=255
	log.Println(a>>24 )
	log.Println(a)
	log.Println(b<<24)
}
