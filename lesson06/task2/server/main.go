package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("sever listening on 8081")
	ls, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}
	defer ls.Close()

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("new connction")
		go listenConnection(conn)
	}
}

func listenConnection(conn net.Conn) {
	for {
		mess := "Hello!"
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			return
		}
		data := buffer[:dataSize]
		fmt.Println("recieved massage:", string(data))

		n, err := strconv.Atoi(string(data[:len(data)-2]))
		if err == nil {
			mess = strconv.Itoa(n * 2)

		} else {
			mess = strings.ToUpper(string(data))

		}

		_, err = conn.Write([]byte(mess + "\n"))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message sent: ", mess)
	}
}
