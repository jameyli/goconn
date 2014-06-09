/**
* @file:   session.go
* @brief:
一次链接就是一个session, session 会启动一个goroutine recv()来进行网络收包, 网络发包虽然也是由session完成，
但并没有直接产生一个goroutine， 而是提供接口共外部调用。

从面向对象的角度讲，session的功能是完整的。

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
	id      int
	conn    net.Conn
	msg2svr chan<- string
}

///@brief:  创建一个新的session，并运行
func NewSession(id int, conn net.Conn, msg2svr chan string) *Session {
	session := &Session{
		id:      id,
		conn:    conn,
		msg2svr: msg2svr,
	}

	fmt.Println("new session:", session)

	go session.run()

	return session
}

func (self Session) String() string {
	return fmt.Sprintf("{id:%d; client_addr:(%s)}", self.id, self.conn.RemoteAddr())
}

func (self *Session) recv() {
	conn := self.conn
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
		self.msg2svr <- reciveStr
	}
}

func (self *Session) SendToClient(msg string) error {
	fmt.Println("send to client:", msg)
	_, err := self.conn.Write([]byte(msg))

	return err
}
