
package main

import (
	"bytes"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello World!")

	//Basic variables
	address := "127.0.0.1:8080"
	protocol := "udp"

	//Build the address
	udpAddr, err := net.ResolveUDPAddr(protocol, address)
	if err != nil {
		fmt.Println("Wrong Address")
		return
	}

	//Output
	fmt.Println("ready " + protocol + " from " + udpAddr.String())

	//Create the connection
	udpConn, err := net.ListenUDP(protocol, udpAddr)
	if err != nil {
		fmt.Println(err)
	}

	//Keep calling this function
	for {
		display(udpConn)
	}

}

func display(conn *net.UDPConn) {
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	stroka := string(bytes.Trim(buf[0:n], "\x00"))
	if err != nil {
		fmt.Println("Error Reading")
		return
	} else {
		fmt.Println(stroka)
		fmt.Println("ok")
	}
}

