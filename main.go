package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
)

func main() {
	fmt.Println("listening on port :6379")
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	conn, err := l.Accept()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	defer conn.Close()
	for {
		resp := NewResp(conn)
		value, err := resp.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			slog.Error("error reading from client: " + err.Error())
			panic(err)
		}
		fmt.Println(fmt.Sprintf("%+v", value))

		conn.Write([]byte("+OK\r\n"))
	}

}
