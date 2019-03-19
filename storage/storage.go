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
	checkKey bool
	params   int
	function func(...string) (string,error)
}

var mapFunc map[string]funcStruct

var db dict
var mu sync.Mutex
func init() {
	db = dict{make(map[string]*value)}
	mapFunc = make(map[string]funcStruct, 10)
	mapFunc["isexsit"] = funcStruct{true, 1, isExist}
}

func (d *dict) exist(key string) bool {
	return d.storage[key]!=nil
}
func isExist(params ...string)  (string,error){
	if db.exist(params[0]){
		return "yes",nil
	}
	return "no",nil
}

func GetResult(params ...string) (result string, err error) {
	if len(params) < 2 {
		return "",errors.New("params too short")
	}
	mapfunc, ok := mapFunc[params[0]]

	if ok {
		if mapfunc.params > 0 {
			if mapfunc.params != len(params)-1 {
				return "",errors.New("params number not right")
			}
		}
		return mapfunc.function(params[1:]...)
	} else {
		return "", errors.New("no such function")
	}

}

func getLock(key string)error{
	defer func() {
		mu.Lock()
		return func() {mu.Unlock()}
}
	if db.exist(key){
		db.storage[key].mutex.RLock()
		return nil
	}else{
		return  errors.New("not exist")
	}
}
func getUnLock(key string)error{
	db.storage[key].mutex.RUnlock()
	return nil
}

func setLock(key string)error{
	if db.exist(key){
		db.storage[key].mutex.Lock()
		return nil
	}else{
		return  errors.New("not exist")
	}
}
func setUnLock(key string)error{
	db.storage[key].mutex.Unlock()
	return nil
}
