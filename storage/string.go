package storage

type str struct {
	data string
}

func getStr(param ...string) (string, error) {
	key := param[0]
	val, err := getLock(key)
	if err != nil {
		return "", err
	}
	s := val.val.(*str).data
	getUnLock(val)
	return s, nil
}

func setStr(param ...string) (string, error) {
	key := param[0]
	v, err := setLock(key)
	defer setUnLock(v, key)
	if err != nil {
		v.val = &str{param[1]}
		return "ok", nil
	} else {
		v.val.(*str).data = param[1]
		return "ok", nil
	}
}

func (s *str) getType() int {
	return 1
}
