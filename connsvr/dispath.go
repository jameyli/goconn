package main

import (
	"fmt"
	"net"
	"time"
)

// type Dispatch struct {
//     id int64
// }

func Dispatch(conns *map[string]net.Conn, client_msg, server_in_msg, server_out_msg chan string) {
	for {
		select {
		case msg := <-client_msg:
			fmt.Println("client_msg:", msg)
			server_in_msg <- msg

		case msg := <-server_out_msg:
			for key, value := range *conns {
				// fmt.Println("connection is connected from ...", key)
				_, err := value.Write([]byte(msg))
				if err != nil {
					fmt.Println(err.Error())
					delete(*conns, key)
				}
			}
		default:
			time.Sleep(time.Second)
		}
	}
}
