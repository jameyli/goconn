package main

import (
	"fmt"
	"time"
)

// type Dispatch struct {
//     id int64
// }

func Dispatch(sessions *map[string]Session, client_msg, server_in_msg, server_out_msg chan string) {
	for {
		select {
		case msg := <-client_msg:
			fmt.Println("client_msg:", msg)
			server_in_msg <- msg

		case msg := <-server_out_msg:
			for key, value := range *sessions {
				fmt.Println("connection is connected from ...", key)
				_, err := value.conn.Write([]byte(msg))
				if err != nil {
					fmt.Println(err.Error())
					delete(*sessions, key)
				}
			}
		default:
			time.Sleep(time.Second)
		}
	}
}
