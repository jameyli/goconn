/**
* @file:   session.go
* @brief:  一次链接就是一个session
* @author: jameyli <lgy AT live DOT com>
*
* @date:   2014-05-22
 */
package main

import (
	"fmt"
	"net"
)

type Session struct {
	id int64
}

func Handler(conn net.Conn, messages chan string) {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	for {
		lenght, err := conn.Read(buf)
		if checkError(err, "Connection") == false {
			conn.Close()
			break
		}
		if lenght > 0 {
			buf[lenght] = 0
		}
		//fmt.Println("Rec[",conn.RemoteAddr().String(),"] Say :" ,string(buf[0:lenght]))
		reciveStr := string(buf[0:lenght])
		messages <- reciveStr

	}
}
