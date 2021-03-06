package main

import (
	"fmt"
	"net"
	"os"
)

func StartServer(front_port, back_port string) {
	service := ":" + front_port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err, "ResolveTCPAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "ListenTCP")

	client_msg := make(chan string, 10)

	server_in_msg := make(chan string, 10)
	server_out_msg := make(chan string, 10)

	sessions := make(map[uint64]*Session)

	go BackDoor(back_port, server_in_msg, server_out_msg)

	go Dispatch(&sessions, client_msg, server_in_msg, server_out_msg)

	id_gen := 1
	for {
		fmt.Println("Listening ...")
		conn, err := l.Accept()
		checkError(err, "Accept")
		fmt.Println("front Accepting ...")

		session := NewSession(id_gen, conn, client_msg)
		if nil == session {
			continue
		}

		sessions[uint64(id_gen)] = session

		id_gen++
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println(os.Args[0], "front_port back_port")
		os.Exit(0)
	}

	StartServer(os.Args[1], os.Args[2])
}

func checkError(err error, info string) (res bool) {

	if err != nil {
		fmt.Println(info + "  " + err.Error())
		return false
	}
	return true
}
