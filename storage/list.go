package storage

import "strconv"

type strlist struct {
	head,tail *node
	len int
}

type node struct{
	value string
	pre,next *node
}

func (list *strlist)getType()int{
	return 2
}
func (list *strlist)pushBack(value string)int{
	n:=node{value,list.tail,nil}
	list.tail,list.tail.next=&n,&n
	list.len++
	return list.len

}

func (list *strlist) popBack() string {
	n:=list.tail
	list.tail.pre.next,list.tail=nil,list.tail.pre
	list.len--
	return n.value
}

func newList(value string)*strlist{
	n:=node{value:value}
	return &strlist{&n,&n,1}
}

func lpush(params ...string)(string, error){
    key:=params[0]
    value:=params[1]
    err:=setLock(key)
	if err != nil {
		db.storage[key].val=newList(value)
		setUnLock(key)
		return "1",nil
	}else {
		i:=db.storage[key].val.(*strlist).pushBack(value)
		setUnLock(key)
		return strconv.Itoa(i),nil
	}
}

func lpop(params ...string)(result string,err error){
	key:=params[0]
	err=setLock(key)
	if err != nil {
		setUnLock(key)
		return "",err
	}else {
		list:=db.storage[key].val.(*strlist)
		if list.len==1{
			result=list.tail.value
			db.storage[key].val=nil
		}else {
			result=list.popBack()
		}
		setUnLock(key)
		return
	}
}

