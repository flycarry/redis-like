package storage


type Data struct {
	hashmap map[string]value
}
type value interface {
	gettype()int
}
const (
	isstr =iota
	islist
	ismap
	isset
)

const (
	sucess =iota
	notExist
	typyError
)

func isRegex(key,p string)bool{
	return true
}
func  isExist(v value) bool{
	return v!=nil
}
func NewData()*Data{
	return &Data{make(map[string]value)}
}
func (d *Data) SetString(key string,value string)int{
	d.hashmap[key]=NewStr(value)
	return sucess
}

func (d *Data) GetString(key string) (result string,err int){
	value:=d.hashmap[key]
	if isExist(value){
		if v,ok:=value.(*str);ok{
			return v.Value(),sucess
		}else {
			return "",typyError
		}
	}else {
		return "", notExist
	}
}

func (d *Data) RPush(key string,value string)(err int){
	v:=d.hashmap[key]
	if isExist(v){
		if v,ok:=v.(*strList);ok{
			v.rpush(value)
			return sucess
		}else {
			return typyError
		}
	}else {
		list:=NewStrList()
		list.rpush(value)
		d.hashmap[key]=list
		return sucess
	}
}

func (d *Data) RPop(key string) (string,int){
	value:=d.hashmap[key]
	if isExist(value){
		if v,ok:=value.(*strList);ok{
			return v.rpop(),sucess
		}else {
			return "",typyError
		}
	}else {
		return "",notExist
	}
}

func (d *Data)Keys(pattern string)[]string{
	keys:=make([]string,0,16)
	for key,_ :=range d.hashmap{
		if isRegex(key,pattern){
			keys=append(keys, key)
		}
	}
	return keys
}


