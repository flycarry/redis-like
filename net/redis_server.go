package net

import (
	"github.com/flycarry/redis-like/storage"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)
var data *storage.Data
func init() {
	data=storage.NewData()
}
const(
	setString = iota
	getString
	rPush
	rPop
)
func Socket_server(port string){
	l,err:=net.Listen("tcp",port)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	for {
		conn,err:=l.Accept()
		if err != nil {
			log.Println(err)
		}
		log.Printf("get request from %s",conn.RemoteAddr())
		go resolve_conn(conn)
	}
}

func resolve_conn(c net.Conn){
	msg,err:=ioutil.ReadAll(c)
	if err != nil {
		c.Close()
		return
	}
	if len(msg)>1{
		pars:=parsePars(msg)
		methon,err:=strconv.Atoi(pars[0])
		if err != nil {
			c.Close()
			return
		}
		switch methon {
		case setString:
			result:=data.SetString(pars[1],pars[2])
			c.Write([]byte(strconv.Itoa(result)))
		case getString:
			resultStr,err:=data.GetString(pars[1])
			if err != -1 {
				c.Write([]byte(strconv.Itoa(err)))
				return
			}
			c.Write([]byte(resultStr))
		}
	}
	c.Close()
}

func parsePars(msg []byte)[]string{
	msgStr:=string(msg)
	return strings.Split(msgStr," ")

}