package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "172.20.10.1:53")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte{0x00, 0x00, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x77, 0x77, 0x77, 0x04, 0x6a, 0x70, 0x72, 0x73, 0x02, 0x63, 0x6f, 0x02, 0x6a, 0x70, 0x00, 0x00, 0x01, 0x00, 0x01})
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(os.Stdout, conn)

	// Headerは24byte
	headBuffer := make([]byte, 24)
	_, err = conn.Read(headBuffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("head:%X", headBuffer)
	// Questionは40byte
	questionBuffer := make([]byte, 40)
	_, err = conn.Read(questionBuffer)
	if err != nil {
		fmt.Println(err)
	}
	// Answerは32byte
	answerBuffer := make([]byte, 32)
	_, err = conn.Read(answerBuffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%X", answerBuffer)

}
