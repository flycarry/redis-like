package storage

import (
	"errors"
	"regexp"
	"strings"
	"sync"
)

// Value represent that value in BigData
type Value struct {
	Princess interface{}
	Lock     sync.RWMutex
}

// Method represent the method that underlying data structure regsiter
type Method func([]string) (string, error)

// ErrMethodNotSupport means that method not support
var ErrMethodNotSupport = errors.New("method not support")

// ErrInvalidNumPara means that invalid number of parameters
var ErrInvalidNumPara = errors.New("invalid number of parameters")

// ErrKeyNotExist means that the key don't exist
var ErrKeyNotExist = errors.New("key do not exist")

// BigData represent a table in redis-like
var BigData map[string]*Value

// MethodMap is a set that include all supported method
var MethodMap map[string]Method

// DataLock is a lock which protect the BigData's keys
var DataLock sync.RWMutex

func init() {
	BigData = make(map[string]*Value)
	MethodMap = make(map[string]Method)
}

// RegisterMethod can let underlying data structure register their method support
func RegisterMethod(methodName string, method Method) (err error) {
	if _, ok := MethodMap[methodName]; ok {
		err = errors.New("method already exist")
		return
	}
	MethodMap[methodName] = method
	return nil
}

// Process accept all commond string, parse it and distribute to the corresponding method
func Process(command string) (result string) {
	temp := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(command), -1)
	if len(temp) < 2 {
		return "-error: invalid number of parameters\r\n"
	}
	method, ok := MethodMap[strings.ToUpper(temp[0])]
	if !ok {
		return "-error: method not support"
	}
	result, err := method(temp[1:])
	if err != nil {
		return "-error: " + err.Error()
	}
	return
}
