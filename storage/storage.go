package storage

import (
	"errors"
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

var (
	// ErrMethodNotSupport means that method not support
	ErrMethodNotSupport = errors.New("method not support")

	// ErrInvalidNumPara means that invalid number of parameters
	ErrInvalidNumPara = errors.New("invalid number of parameters")

	// ErrKeyNotExist means that the key don't exist
	ErrKeyNotExist = errors.New("key do not exist")

	// ErrInvalidPara means that the parameters is invalidPara
	ErrInvalidPara = errors.New("invalid parameters")

	// ErrMismatchStruct means that the method is not suitable for this data structure
	ErrMismatchStruct = errors.New("mismatched data structure")
)

var (
	// BigData represent a table in redis-like
	BigData = make(map[string]*Value)

	// MethodMap is a set that include all supported method
	MethodMap = make(map[string]Method)

	// DataLock is a lock which protect the BigData's keys
	DataLock sync.RWMutex
)

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
	temp := strings.Fields(command)
	if len(temp) < 2 {
		return "-error: invalid number of parameters"
	}
	method, ok := MethodMap[strings.ToLower(temp[0])]
	if !ok {
		return "-error: method not support"
	}
	result, err := method(temp[1:])
	if err != nil {
		return "-error: " + err.Error()
	}
	return
}
