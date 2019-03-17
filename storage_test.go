package main

import (
	"fmt"
	"github.com/flycarry/redis-like/storage"
	"testing"
)

var data *storage.Data
func init() {
	data=storage.NewData()
}
func TestStorageStr(t *testing.T) {
	data.SetString("liufei","test")
	fmt.Println(data.GetString("liufei"))
	data.SetString("liufei","nihao")
	fmt.Println(data.GetString("liufei"))
}

func TestStorageStrList(t *testing.T) {
	data.RPush("zhangrui","nihao")
	fmt.Println(data.RPop("zhangrui"))
	//data.RPush("zhuangrui","nihao")
	//fmt.Println(data.RPop("zhangrui"))

}
