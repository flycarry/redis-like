package storage


type data struct {
	hashmap map[string]value
	size int
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
func  isExist(v value) bool{
	return v!=nil
}
func (d *data) SetString(key string,value string)int{
	d.hashmap[key]=NewStr(value)
	return sucess
}

func (d *data) GetString(key string) (result string,err int){
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

func (d *data) Push(key string,value string){

}

func (d *data) Pop(key string) string{
	return ""
}
