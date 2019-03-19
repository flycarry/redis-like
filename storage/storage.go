package storage

import (
	"errors"
	"sync"
)

type dict struct {
	storage map[string]value
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
	function func(...string) func(...string) string
}

var mapFunc map[string]funcStruct

var db dict

func init() {
	db = dict{make(map[string]value)}
	mapFunc = make(map[string]funcStruct, 10)
	mapFunc["isexsit"] = funcStruct{true, 1, isExist}
}

func (d *dict) exist(key string) bool {

	d.storage[key].mutex.RLock()
	defer d.storage[key].mutex.RUnlock()
	_, ok := d.storage[key]
	if ok {
		return true
	} else {
		return false
	}
}
func isExist(_ ...string) func(...string) string {
	return func(_ ...string) string {
		return "exist"
	}
}

func (mapfunc *funcStruct) checkFunc(subparams ...string) error {
	if mapfunc.params > 0 {
		if mapfunc.params != len(subparams) {
			return errors.New("params number not right")
		}
	}
	if mapfunc.checkKey {
		if !db.exist(subparams[0]) {
			return errors.New("key not exist")
		}
	}
	return nil
}
func GetResult(params ...string) (result string, err error) {
	if len(params) < 2 {
	}
}
func GetFunc(params ...string) func(...string) string {
	if len(params) < 2 {
		return errResponse(errors.New("params too short"))
	}
	mapstruct, ok := mapFunc[params[0]]
	if ok {
		if err := mapstruct.checkFunc(params[1:]...); err != nil {
			return errResponse(err)
		}
		return mapstruct.function(params[1:]...)
	} else {
		return func(_ ...string) string { return "no such function" }
	}
}

func errResponse(err error) func(...string) string {
	return func(_ ...string) string {
		return err.Error()
	}
}
