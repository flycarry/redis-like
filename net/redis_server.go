package net

import (
	"github.com/flycarry/redis-like/route"
	"io"
	"log"
	"net"
	"strings"
)

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
	for {
		buf := make([]byte, 512)
		l, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			if err==io.EOF{
				return
			}
		}
		log.Println(string(buf[:l]))
		_,err=io.WriteString(c, route.DoReply(string(buf[:l]))+"\n")
		if err != nil {
			panic("socket error")
		}
		defer c.Close()
	}
}

func parsePars(msg []byte)[]string{
	msgStr:=string(msg)
	return strings.Split(msgStr,"\r\n")

}