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
type errorKeyNotExits int

func (err *errorKeyNotExits) Error() string {
	return "not exist"
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
	mapFunc["lindex"] = funcStruct{params: 2, function: lindex}
	mapFunc["linsert"] = funcStruct{params: 3, function: linsert}
}

func (d *dict) exist(key string) bool {
	val, err := getLock(key)
	if err != nil {
		return false
	}
	defer getUnLock(val)
	return val == nil
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

func getLock(key string) (*value, error) {
	mu.RLock()
	defer mu.RUnlock()
	if db.storage[key] != nil {
		db.storage[key].mutex.RLock()
		return db.storage[key], nil
	} else {
		return db.storage[key], errors.New("not exist")
	}
}
func getUnLock(value *value) {
	value.mutex.RUnlock()
}

func setLock(key string) (*value, error) {
	mu.Lock()
	defer mu.Unlock()
	if db.storage[key] != nil {
		v := db.storage[key]
		v.mutex.Lock()
		return v, nil
	} else {
		db.storage[key] = &value{sync.RWMutex{}, nil}
		v := db.storage[key]
		v.mutex.Lock()
		return v, errors.New("not exist")
	}
}
func setUnLock(v *value, key string) {
	mu.Lock()
	defer mu.Unlock()
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
	v, err := setLock(key)
	if err != nil {
		setUnLock(v, key)
		return "", err
	} else {
		v.val = nil
		setUnLock(v, key)
		return "delete success", nil
	}
}
