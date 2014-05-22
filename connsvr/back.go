package main

import (
	"fmt"
	"net"
)

func BackDoor(port string, server_in_msg, server_out_msg chan string) {
	service := ":" + port
	udpAddr, err := net.ResolveTCPAddr("tcp", service)
	if nil != err {
		return
	}

	l, err := net.ListenTCP("tcp", udpAddr)
	if nil != err {
		return
	}

	for {
		conn, err := l.Accept()
		if nil != err {
			continue
		}
		fmt.Println("back Accepting ...")

		go In(conn, server_in_msg)
		go Out(conn, server_out_msg)
	}

}

func In(conn net.Conn, server_in_msg chan string) {
	for {
		msg := <-server_in_msg

		_, err := conn.Write([]byte(msg))
		if nil != err {
			fmt.Println("ServerIn error")
		}
	}
}

func Out(conn net.Conn, server_out_msg chan string) {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	for {
		lenght, err := conn.Read(buf)
		if nil != err {
			conn.Close()
			break
		}
		if lenght > 0 {
			buf[lenght] = 0
		}
		reciveStr := string(buf[0:lenght])

		fmt.Println("SeverOut", reciveStr)
		server_out_msg <- reciveStr
	}
}
