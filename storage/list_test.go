package storage

import (
	"log"
	"testing"
)

func TestListOp(t *testing.T) {
	l := newList("test")
	l.pushFront("world")
	l.popFront()
	_ = l.insert("test", "nihao")
	l.popBack()
	log.Println(l.lgets(0, 10), l.len)
}
