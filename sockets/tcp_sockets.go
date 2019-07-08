package main

import "fmt"
import "net"
import "os"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service:= os.Args[1]
	tcpAddr, error := net.ResolveTCPAddr("tcp4", service)
	checkError(error)
	conn, error := net.DialTCP("tcp", nil, tcpAddr)
	checkError(error)
	_, error = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(error)
	result := make([]byte, 256)
	_, error = conn.Read(result)
	fmt.Println(string(result))
	os.Exit(0)
}


func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}