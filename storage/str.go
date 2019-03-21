package storage

func init() {
	RegisterMethod("set", Set)
	RegisterMethod("get", Get)
	RegisterMethod("del", Del)
}

// Set update the str
func Set(args []string) (result string, err error) {
	if len(args) != 2 {
		return "", ErrInvalidNumPara
	}
	strKey := args[0]
	DataLock.RLock()
	if strValue, ok := BigData[strKey]; ok {
		DataLock.RUnlock()
		strValue.Lock.RLock()
		strValue.Princess = args[1]
		strValue.Lock.RUnlock()
	} else {
		DataLock.RUnlock()
		DataLock.Lock()
		BigData[strKey] = &Value{
			Princess: args[1],
		}
		DataLock.Unlock()
	}
	return "+OK", nil
}

// Get query the str
func Get(args []string) (result string, err error) {
	if len(args) != 1 {
		return "", ErrInvalidNumPara
	}
	strKey := args[0]
	DataLock.RLock()
	if strValue, ok := BigData[strKey]; ok {
		result, ok = strValue.Princess.(string)
		if !ok {
			return "", ErrMismatchStruct
		}
		DataLock.RUnlock()
		err = nil
		return
	}
	DataLock.RUnlock()
	return "", ErrKeyNotExist
}

// Del delete the str
func Del(args []string) (result string, err error) {
	if len(args) != 1 {
		return "", ErrInvalidNumPara
	}
	strKey := args[0]
	DataLock.RLock()
	if _, ok := BigData[strKey]; ok {
		DataLock.RUnlock()
		DataLock.Lock()
		delete(BigData, strKey)
		DataLock.Unlock()
		return "+OK", nil
	}
	return "", ErrKeyNotExist
}
