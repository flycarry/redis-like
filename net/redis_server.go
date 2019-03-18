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
	buf:=make([]byte,512)
	l,err:=c.Read(buf)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(buf[:l]))
	io.WriteString(c,route.DoReply(string(buf[:l])))
	defer c.Close()
}

func parsePars(msg []byte)[]string{
	msgStr:=string(msg)
	return strings.Split(msgStr,"\r\n")

}