package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		io.WriteString(conn, fmt.Sprint("Ok\n", time.Now(), "\n"))

		conn.Close()
	}
}