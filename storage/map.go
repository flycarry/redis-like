package storage

import (
	"errors"
	"hash/crc32"
)

const defaultCapaity = 256

var keyNotFound = errors.New("key not found")

type myMap struct {
	len    int
	values []*entry
}

type entry struct {
	key   string
	value string
	next  *entry
}

func newMap() *myMap {
	return initMap(defaultCapaity)
}
func initMap(cap int) *myMap {
	return &myMap{
		0,
		make([]*entry, cap, cap),
	}
}
func (m *myMap) set(key, value string) {
	index := hash(key) % cap(m.values)

	entrys := m.values[index]
	entryTemp := entry{key, value, nil}
	if entrys == nil {
		m.values[index] = &entryTemp
		m.len++
		return
	}
	for ; entrys.next != nil; entrys = entrys.next {
		if entrys.key == key {
			entrys.value = value
			return
		}
	}
	if entrys.key == key {
		entrys.value = value
	} else {
		entrys.next = &entryTemp
		m.len++
	}
}
func (m *myMap) get(key string) (string, error) {
	index := hash(key) % cap(m.values)
	entrys := m.values[index]
	for ; entrys != nil; entrys = entrys.next {
		if entrys.key == key {
			return entrys.value, nil
		}
	}
	return "", keyNotFound
}
func (m *myMap) remap() {
	if m.len*100 > (cap(m.values)*75){

	}
}
func hash(s string) int {
	return hashtool(s)
}
func hashWithcrc(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}
func hashtool(s string) int {
	sum := 0
	for b := range []byte(s) {
		temp := sum & (255 << 24)
		temp = temp >> 24
		sum = sum << 8
		temp ^= b
		sum &= temp
	}
	return sum
}
