package main

import (
	"log"
	"sync"
	"testing"
)

type value struct {
	mutex sync.Mutex
	val   data
}
type data int
var m map[string]*value
var mu sync.Mutex
func init() {
	m=make(map[string]*value)
}
func TestValue(t *testing.T) {
	log.Println(m["liufei"])
	//m["liufei"]=&value{sync.Mutex{},0}
	m["liufei"].mutex.Lock()
	log.Println(m["liufei"])
	m["liufei"].mutex.Unlock()
}

func TestMutex(t *testing.T) {
	mu.Lock()
	log.Println("lock")
	mu.Unlock()
}