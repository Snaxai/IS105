package tcpserver

import (
	"fmt"
	"io/ioutil"
	"net"
)

//Tcpdial asd
func Tcpdial() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}
