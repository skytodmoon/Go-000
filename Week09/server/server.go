package main

import (
	"fmt"
	"net"
)

func main() {

	mesgChan := make(chan string)
	fmt.Println("Start listening...")
	listener, err := net.Listen("tcp", "localhost:50000")

	if err != nil {
		fmt.Println("Error listening", err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accept", err.Error())
			return
		}
		go readConn(conn, mesgChan)
		go sendConn(conn, mesgChan)
	}
}

func readConn(conn net.Conn, inputChan chan<- string) {
	defer conn.Close()
	var buf [1024]byte
	for {

		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from %s msg faild err:[%v]\n", conn.RemoteAddr().String(), err)
			break
		}
		fmt.Printf("rev data from %s msg:%s\n", conn.RemoteAddr().String(), string(buf[:n]))
		inputChan <- string(buf[:n])
	}
}

func sendConn(conn net.Conn, outputChan <-chan string) {
	defer conn.Close()
	for {
		msg := <-outputChan
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("send data error to %s msg:%s\n", conn.RemoteAddr().String(), msg)
		}
	}
}
