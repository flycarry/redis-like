package storage

type str struct {
	data string
}

func (s *str)getType()int{
	return 1
}
func getStr(param ...string)(string,error){
	key:=param[0]
	err:=getLock(key)
	if err != nil {
		return "",err
	}
	getUnLock(key)
	return db.storage[key].val.(*str).data,nil
}

func setStr(param ...string)(string,error){
	key:=param[0]
	err:=setLock(key)
	if err != nil {
		db.storage[key].val=&str{param[1]}
		return "ok",nil
	}else{
		db.storage[key].val.(*str).data=param[1]
		setUnLock(key)
		return "ok",nil
	}
}
