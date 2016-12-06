package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("up4", service)
	if err != nil {
		fmt.Println("Sender: Address Resolution error: ", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Sender: UDP Dial error: ", err)
		os.Exit(1)
	}

	_, err = conn.Write([]byte("Test"))
	if err != nil {
		fmt.Println("Sender: Connection write error: ", err)
		os.Exit(1)
	}

	buf := make([]byte, 512, 512)

	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Println("Sender: Connection read error: ", err)
		os.Exit(1)
	}

	fmt.Println(string(buf[0:n]))

	os.Exit(0)

}
