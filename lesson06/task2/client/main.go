package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	//get connection
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}
	//send greeting message
	_, err = conn.Write([]byte("Hello Server!"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message sent: Hello Server!")
	// read greetings
	buffer := make([]byte, 1400)
	dataSize, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("connection closed")
		return
	}
	data := buffer[:dataSize]
	fmt.Println("recieved massage:", string(data))

	// sending STDIN
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Your input: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Your input is: %v\n", text)
		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Fatalln(err)
		}
		// Exit if input = exit
		str := strings.TrimSpace(string(text))
		if strings.Compare(str, "exit") == 0 {
			conn.Close()
		}

		fmt.Printf("Message sent: %s", text)

		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			return
		}
		data := buffer[:dataSize]
		fmt.Println("recieved massage:", string(data))
	}
}
