package tcpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

type MemberGroup struct {
	Name  string
	Epost string
}

//TcplistenB asd
func TcplistenB() {
	group := MemberGroup{
		Name:  "Daniel",
		Epost: "Daniel@",
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

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
		Bytestostring(b)
		io.WriteString(conn, fmt.Sprint(Bytestostring(b)))
		conn.Close()
	}
}

//Bytestostring convertes bytes to string
func Bytestostring(data []byte) string {
	return string(data[:])
}
