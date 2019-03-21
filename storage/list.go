package storage

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func init() {
	RegisterMethod("lpush", Lpush)
	RegisterMethod("rpush", Rpush)
	RegisterMethod("lpop", Lpop)
	RegisterMethod("rpop", Rpop)
	RegisterMethod("lrange", Lrange)
}

type strongList struct {
	*list.List
}

func newStrongList() *strongList {
	return &strongList{List: list.New()}
}

// Lpush means push value into list, a data structure in table, by left
func Lpush(args []string) (string, error) {
	if len(args) < 2 {
		return "", ErrInvalidNumPara
	}
	strKey := args[0]
	tempList, tempLock, err := readTableHelper(strKey, false)
	if err == ErrKeyNotExist {
		tempList, tempLock, err = writeTableHelper(strKey)
		if err != nil {
			return "", err
		}
		for _, arg := range args[1:] {
			tempList.PushFront(arg)
		}
		tempLock.Unlock()
		return "+OK", nil

	}
	if err != nil {
		return "", err
	}
	for _, arg := range args[1:] {
		tempList.PushFront(arg)
	}
	tempLock.Unlock()
	return "+OK", nil
}

// Rpush means push value into list, a data structure in table, by right
func Rpush(args []string) (string, error) {
	if len(args) < 2 {
		return "", ErrInvalidNumPara
	}
	strKey := args[0]
	tempList, tempLock, err := readTableHelper(strKey, false)
	if err == ErrKeyNotExist {
		tempList, tempLock, err = writeTableHelper(strKey)
		if err != nil {
			return "", err
		}
		for _, arg := range args[1:] {
			tempList.PushBack(arg)
		}
		tempLock.Unlock()
		return "+OK", nil

	}
	if err != nil {
		return "", err
	}
	for _, arg := range args[1:] {
		tempList.PushBack(arg)
	}
	tempLock.Unlock()
	return "+OK", nil
}

// Lpop means pop value from list, a data structure in table, by left
func Lpop(args []string) (result string, err error) {
	if len(args) != 1 {
		return "", ErrInvalidNumPara
	}
	strKey := args[0]
	tempList, tempLock, err := readTableHelper(strKey, false)
	if err != nil {
		return
	}
	if tempList.Len() == 0 {
		err = ErrKeyNotExist
	} else {
		result = tempList.Remove(tempList.Front()).(string)
	}
	tempLock.Unlock()
	return
}

// Rpop means pop value from list, a data structure in table, by right
func Rpop(args []string) (result string, err error) {
	if len(args) != 1 {
		return "", ErrInvalidNumPara
	}
	strKey := args[0]
	tempList, tempLock, err := readTableHelper(strKey, false)
	if err != nil {
		return
	}
	if tempList.Len() == 0 {
		err = ErrKeyNotExist
	} else {
		result = tempList.Remove(tempList.Back()).(string)
	}
	tempLock.Unlock()
	return
}

// Lrange list the value from start to end(start and end should in "args")
// when start is more than end, the return is "(empty list)" without error
func Lrange(args []string) (string, error) {
	if len(args) != 3 {
		return "", ErrInvalidNumPara
	}

	// check if parameters is legal
	start, err := intFromStr(args[1], len(args))
	if err != nil {
		return "", err
	}
	end, err := intFromStr(args[2], len(args))
	if err != nil {
		return "", err
	}
	if start >= end {
		return "(empty list)", nil
	}

	strKey := args[0]
	tempList, tempLock, err := readTableHelper(strKey, true)
	if err != nil {
		return "", err
	}
	if tempList.Len() == 0 {
		tempLock.RUnlock()
		return "", ErrKeyNotExist
	}

	// the list is exist
	tempArray := make([]string, 0, tempList.Len())
	index := 0
Loop:
	for tempEle := tempList.Front(); tempEle != nil; tempEle = tempEle.Next() {
		switch {
		case index < start:
			index++
		case index >= start && index < end:
			tempArray = append(tempArray, fmt.Sprintf("%d) %s", index+1, tempEle.Value.(string)))
			index++
		default:
			break Loop
		}
	}

	if len(tempArray) == 0 {
		return "(empty list)", nil
	}
	return strings.Join(tempArray, "\r\n"), nil
}

func intFromStr(index string, len int) (ret int, err error) {
	ret, err = strconv.Atoi(index)
	if err != nil {
		return 0, ErrInvalidPara
	}
	if len < 0 {
		ret += len + 1
	}
	return
}

func tryTransform(value interface{}) (ret *strongList, err error) {
	ret, ok := value.(*strongList)
	if !ok {
		return nil, ErrMismatchStruct
	}
	return
}

// writeTableHelper help function Lpush/Rpush to write, and caller should call lock's Lock()
func writeTableHelper(strKey string) (tempList *strongList, lock *sync.RWMutex, err error) {
	DataLock.Lock()
	if value, ok := BigData[strKey]; ok {
		lock = &value.Lock
		lock.Lock()
		DataLock.Unlock()
		tempList, err = tryTransform(value.Princess)
	} else {
		value := &Value{
			Princess: newStrongList(),
		}
		lock = &value.Lock
		lock.Lock()
		BigData[strKey] = value
		DataLock.Unlock()
		tempList, err = tryTransform(value.Princess)
	}
	return
}

// readTableHelper check the key is exist or not, if not `err` will be ErrKeyNotExist
// else return tempList and caller should call lock's RUlock()/Ulock by read
func readTableHelper(strKey string, read bool) (tempList *strongList, lock *sync.RWMutex, err error) {
	DataLock.RLock()
	if value, ok := BigData[strKey]; ok {
		lock = &value.Lock
		if read {
			lock.RLock()
		} else {
			lock.Lock()
		}
		DataLock.RUnlock()
		tempList, err = tryTransform(value.Princess)
		// transfer succeed
	} else {
		DataLock.RUnlock()
		err = ErrKeyNotExist
	}
	return
}
