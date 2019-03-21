package storage

import (
	"strconv"
)

type strlist struct {
	head, tail *node
	len        int
}

type node struct {
	pre, next *node
	value     string
}

func (list *strlist) getType() int {
	return 2
}
func newList(value string) *strlist {
	n := node{value: value}
	return &strlist{&n, &n, 1}
}

func (list *strlist) pushBack(value string) int {
	n := node{value: value, pre: list.tail}
	list.tail.next, list.tail = &n, &n
	list.len++
	return list.len

}
func (list *strlist) pushFront(value string) int {
	n := node{value: value, next: list.head}
	list.head.pre, list.head = &n, &n
	list.len++
	return list.len

}
func (list *strlist) popBack() (string) {
	if list.len > 1 {

		value := list.tail.value
		list.tail, list.tail.pre.next = list.tail.pre, nil
		list.len--
		return value
	} else {
		value := list.tail.value
		list.tail, list.head = nil, nil
		list.len--
		return value
	}
}
func (list *strlist) popFront() (string) {
	if list.len > 1 {
		value := list.head.value
		list.head, list.head.next.pre = list.head.next, nil
		list.len--
		return value
	} else {
		value := list.tail.value
		list.tail, list.head = nil, nil
		list.len--
		return value
	}
}
func lpush(params ...string) (string, error) {
	key := params[0]
	value := params[1]
	err := setLock(key)
	if err != nil {
		db.storage[key].val = newList(value)
		setUnLock(key)
		return "1", nil
	} else {
		result := strconv.Itoa(db.storage[key].val.(*strlist).pushBack(value))
		setUnLock(key)
		return result, nil
	}
}
func rpush(params ...string) (string, error) {
	key := params[0]
	value := params[1]
	err := setLock(key)
	if err != nil {
		db.storage[key].val = newList(value)
		setUnLock(key)
		return "1", nil
	} else {
		result := strconv.Itoa(db.storage[key].val.(*strlist).pushFront(value))
		setUnLock(key)
		return result, nil
	}
}
func lpop(params ...string) (string, error) {
	key := params[0]
	err := setLock(key)
	if err != nil {
		setUnLock(key)
		return "", err
	} else {
		list := db.storage[key].val.(*strlist)
		result := list.popBack()
		if list.len == 0 {
			db.storage[key].val = nil
		}
		setUnLock(key)
		return result, nil
	}

}
func rpop(params ...string) (string, error) {
	key := params[0]
	err := setLock(key)
	if err != nil {
		setUnLock(key)
		return "", err
	} else {
		list := db.storage[key].val.(*strlist)
		result := list.popFront()
		if list.len == 0 {
			db.storage[key].val = nil
		}
		setUnLock(key)
		return result, nil
	}

}
func lrange(params ...string) (string, error) {
	key := params[0]
	err := getLock(key)
	if err != nil {
		getUnLock(key)
		return "", err
	} else {
		return "",nil

	}
}
