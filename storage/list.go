package storage

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
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
	DataLock.RLock()
	if value, ok := BigData[strKey]; ok {
		value.Lock.Lock()
		DataLock.RUnlock()
		tempList := value.Princess.(*strongList)
		for _, arg := range args[1:] {
			tempList.PushFront(arg)
		}
		value.Lock.Unlock()
		return "+OK", nil
	}
	DataLock.RUnlock()
	DataLock.Lock()
	if value, ok := BigData[strKey]; ok {
		value.Lock.Lock()
		DataLock.Unlock()
		tempList := value.Princess.(*strongList)
		for _, arg := range args[1:] {
			tempList.PushFront(arg)
		}
		value.Lock.Unlock()
		return "+OK", nil
	}
	value := &Value{
		Princess: newStrongList(),
	}
	value.Lock.Lock()
	BigData[strKey] = value
	DataLock.Unlock()
	tempList := value.Princess.(*strongList)
	for _, arg := range args[1:] {
		tempList.PushFront(arg)
	}
	value.Lock.Unlock()
	return "+OK", nil
}

// Rpush means push value into list, a data structure in table, by right
func Rpush(args []string) (string, error) {
	if len(args) < 2 {
		return "", ErrInvalidNumPara
	}
	strKey := args[0]
	DataLock.RLock()
	if value, ok := BigData[strKey]; ok {
		value.Lock.Lock()
		DataLock.RUnlock()
		tempList := value.Princess.(*strongList)
		for _, arg := range args[1:] {
			tempList.PushBack(arg)
		}
		value.Lock.Unlock()
		return "+OK", nil
	}
	DataLock.RUnlock()
	DataLock.Lock()
	if value, ok := BigData[strKey]; ok {
		value.Lock.Lock()
		DataLock.Unlock()
		tempList := value.Princess.(*strongList)
		for _, arg := range args[1:] {
			tempList.PushBack(arg)
		}
		value.Lock.Unlock()
		return "+OK", nil
	}
	value := &Value{
		Princess: newStrongList(),
	}
	value.Lock.Lock()
	BigData[strKey] = value
	DataLock.Unlock()
	tempList := value.Princess.(*strongList)
	for _, arg := range args[1:] {
		tempList.PushBack(arg)
	}
	value.Lock.Unlock()
	return "+OK", nil
}

// Lpop means pop value from list, a data structure in table, by left
func Lpop(args []string) (result string, err error) {
	if len(args) != 1 {
		return "", ErrInvalidNumPara
	}
	strKey := args[0]
	DataLock.RLock()
	if value, ok := BigData[strKey]; ok {
		value.Lock.Lock()
		DataLock.RUnlock()
		tempList := value.Princess.(*strongList)
		if tempList.Len() == 0 {
			value.Lock.Unlock()
			return "", ErrKeyNotExist
		}
		result = tempList.Remove(tempList.Front()).(string)
		value.Lock.Unlock()
	} else {
		DataLock.RUnlock()
	}
	return
}

// Rpop means pop value from list, a data structure in table, by right
func Rpop(args []string) (result string, err error) {
	if len(args) != 1 {
		return "", ErrInvalidNumPara
	}
	strKey := args[0]
	DataLock.RLock()
	if value, ok := BigData[strKey]; ok {
		value.Lock.Lock()
		DataLock.RUnlock()
		tempList := value.Princess.(*strongList)
		if tempList.Len() == 0 {
			value.Lock.Unlock()
			return "", ErrKeyNotExist
		}
		result = tempList.Remove(tempList.Back()).(string)
		value.Lock.Unlock()
	} else {
		DataLock.RUnlock()
	}
	return
}

// Lrange list the value from start to end(start and end should in "args")
// when start is more than end, the return is "(empty list)" without error
func Lrange(args []string) (string, error) {
	if len(args) != 3 {
		return "", ErrInvalidNumPara
	}
	start, err := ifstrwmod(args[1], len(args))
	if err != nil {
		return "", err
	}
	end, err := ifstrwmod(args[2], len(args))
	if err != nil {
		return "", err
	}
	if start >= end {
		return "(empty list)", nil
	}
	strKey := args[0]
	DataLock.RLock()
	if value, ok := BigData[strKey]; ok {
		value.Lock.RLock()
		DataLock.RUnlock()
		tempList := value.Princess.(*strongList)
		if tempList.Len() == 0 {
			value.Lock.RUnlock()
			return "", ErrKeyNotExist
		}
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
	DataLock.RUnlock()
	return "", ErrKeyNotExist
}

func ifstrwmod(index string, len int) (ret int, err error) {
	ret, err = strconv.Atoi(index)
	if err != nil {
		return 0, ErrInvalidPara
	}
	if len < 0 {
		ret += len + 1
	}
	return
}
