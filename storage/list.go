package storage

import (
	"container/list"
)

type strList struct {
	list.List

}
func (l *strList)gettype()int {
	return islist
}

func (l *strList)rpush(value string){
	l.PushFront(value)
}

func (l *strList)lpush(value string){
	l.PushFront(value)
}
func (l *strList)lpop()string{
	value:=l.Front().Value.(string)
	l.Remove(l.Front())
	return value
}
func (l *strList)rpop()string{
	value:=l.Front().Value.(string)
	l.Remove(l.Front())
	return value
}
func NewStrList() *strList{
	l:=list.New()
	return &(strList{*l})
}

