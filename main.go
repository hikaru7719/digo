package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "103.5.140.1:53")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte{0x00, 0x00, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x77, 0x77, 0x77, 0x04, 0x6a, 0x70, 0x72, 0x73, 0x02, 0x63, 0x6f, 0x02, 0x6a, 0x70, 0x00, 0x00, 0x01, 0x00, 0x01})
	if err != nil {
		fmt.Println(err)
	}
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	packet := buffer[:length]
	fmt.Printf("%X\n", packet)
	// Headerは12byte
	fmt.Printf("%X\n", packet[:12])
	// Questionは20byte
	fmt.Printf("%X\n", packet[12:12+20])
	// Answerは16byte
	fmt.Printf("%X\n", packet[12+20:12+20+16])

}
