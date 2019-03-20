package net

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/flycarry/redis-like/storage"
)

var messageSeparator = []byte{'\r', '\n'}

// SocketServer listen to the local port
func SocketServer(addr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go resolveConn(conn)
	}
}

func resolveConn(conn net.Conn) {
	split := func(data []byte, atEOF bool) (int, []byte, error) {
		index := bytes.Index(data, messageSeparator)
		if index != -1 {
			return index + 2, data[:index], nil
		}
		if atEOF {
			return 0, nil, io.EOF
		}
		return 0, nil, nil
	}
	scanner := bufio.NewScanner(conn)
	scanner.Split(split)
	for scanner.Scan() {
		result := storage.Process(scanner.Text())
		io.WriteString(conn, result+string(messageSeparator))
	}
	fmt.Println("connection over")
	return
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
