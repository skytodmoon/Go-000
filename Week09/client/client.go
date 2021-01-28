package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Start connect to server...")
	conn, err := net.Dial("tcp", "localhost:50000")
	defer conn.Close()
	if err != nil {
		fmt.Println("Connect error", err.Error())
		return
	}

	input := bufio.NewReader(os.Stdin)

	var buf [1024]byte
	for {
		bytes, _, err := input.ReadLine()
		if err != nil {
			fmt.Printf("read line faild err:%v\n", err)
		}

		str := string(bytes)
		if str == "Q" {
			fmt.Println("exe quit!")
			return
		}
		n, err := conn.Write(bytes)

		if err != nil {
			fmt.Printf("send data faild err:%v\n", err)
		} else {
			fmt.Printf("send data length %d\n", n)
		}
		read, err := conn.Read(buf[:])

		if err != nil {
			fmt.Printf("receive data faild err:%v\n", err)
		} else {
			fmt.Printf("receive data length %d\n", read)
		}

		fmt.Printf("rev data from %s msg:%s\n", conn.RemoteAddr().String(), string(buf[:n]))

	}

}
