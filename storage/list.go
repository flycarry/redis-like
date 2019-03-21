package storage

import (
	"strconv"
	"strings"
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
func (list *strlist) lgets(start, end int) ([]string) {
	ss := make([]string, 0, list.len)
	if end < 0 {
		end = list.len + end
	}
	if start >= end {
		return ss
	}
	for i, node := 0, list.head; i < end; node = node.next {
		if i >= start {
			ss = append(ss, node.value)
		}
		i++
	}
	return ss

}
func (list *strlist) rgets(start, end int) ([]string) {
	ss := make([]string, 0, list.len)
	if end < 0 {
		end = list.len + end
	}
	if start < end {
		return ss
	}
	for i, node := 0, list.tail; i <= end; node = node.pre {
		ss = append(ss, node.value)
	}
	return ss
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
	start := params[1]
	end := params[2]
	from, err := strconv.Atoi(start)
	if err != nil {
		return "", err
	}
	to, err := strconv.Atoi(end)
	if err != nil {
		return "", err
	}
	err = getLock(key)
	if err != nil {
		return "", err
	} else {
		list := db.storage[key].val.(*strlist)
		ss := list.lgets(from, to)
		getUnLock(key)
		return strings.Join(ss, "\n"), nil
	}
}
