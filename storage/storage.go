package storage

import (
	"errors"
	"sync"
)

type dict struct {
	storage map[string]*value
}
type value struct {
	mutex sync.RWMutex
	val   data
}
type data interface {
	getType() int
}
type funcStruct struct {
	params    int
	function  func(...string) (string, error)
	allowType int
}

var mapFunc map[string]funcStruct
var db dict
var mu sync.RWMutex

//init db and function map
func init() {
	db = dict{make(map[string]*value)}
	mapFunc = make(map[string]funcStruct, 10)
	mapFunc["exist"] = funcStruct{params: 1, function: isExist}
	mapFunc["del"] = funcStruct{params: 1, function: delKey}
	mapFunc["set"] = funcStruct{params: 2, function: setStr}
	mapFunc["get"] = funcStruct{params: 1, function: getStr}
	mapFunc["lpush"] = funcStruct{params: 2, function: lpush}
	mapFunc["lpop"] = funcStruct{params: 1, function: lpop}
	mapFunc["rpush"] = funcStruct{params: 2, function: rpush}
	mapFunc["rpop"] = funcStruct{params: 1, function: rpop}
	mapFunc["lrange"] = funcStruct{params: 3, function: lrange}
}

func (d *dict) exist(key string) bool {
	err := getLock(key)
	if err != nil {
		return false
	}
	defer getUnLock(key)
	return d.storage[key] == nil
}

func GetResult(params ...string) (result string, err error) {
	if len(params) < 2 {
		return "", errors.New("params too short")
	}
	mapfunc, ok := mapFunc[params[0]]

	if ok {
		if mapfunc.params > 0 {
			if mapfunc.params != len(params)-1 {
				return "", errors.New("params number not right")
			}
		}
		return mapfunc.function(params[1:]...)
	} else {
		return "", errors.New("no such function")
	}

}
func getType(key string) (int) {
	err := getLock(key)
	if err != nil {
		return 0
	} else {
		t := db.storage[key].val.getType()
		getUnLock(key)
		return t
	}
}
func getLock(key string) error {
	mu.RLock()
	if db.storage[key] != nil {
		db.storage[key].mutex.RLock()
		mu.RUnlock()
		return nil
	} else {
		mu.RUnlock()
		return errors.New("not exist")
	}
}
func getUnLock(key string) {
	db.storage[key].mutex.RUnlock()
}

func setLock(key string) error {
	mu.Lock()
	if db.storage[key] != nil {
		db.storage[key].mutex.Lock()
		mu.Unlock()
		return nil
	} else {
		db.storage[key] = &value{sync.RWMutex{}, nil}
		db.storage[key].mutex.Lock()
		mu.Unlock()
		return errors.New("not exist")
	}
}
func setUnLock(key string) {
	v := db.storage[key]
	if v.val == nil {
		db.storage[key] = nil
		v.mutex.Unlock()
	} else {
		v.mutex.Unlock()
	}
}
func isExist(params ...string) (string, error) {
	if db.exist(params[0]) {
		return "yes", nil
	}
	return "no", nil
}

func delKey(params ...string) (string, error) {
	key := params[0]
	err := setLock(key)
	if err != nil {
		setUnLock(key)
		return "", err
	} else {
		db.storage[key].val = nil
		setUnLock(key)
		return "delete success", nil
	}
}
