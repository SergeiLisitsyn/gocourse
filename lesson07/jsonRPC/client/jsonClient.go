package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"

	"github.com/SergeiLisitsyn/gocourse/lesson07/jsonRPC/common"
)

var reply common.Item
var list common.List

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:8082") //Resolve to connction with server
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(os.Stdin) //Create reader for input data
	//Loop to input three names with one of them - John
	for i := 0; i < 3; i++ {
		fmt.Println("Input Name to add to the List")
		name, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		//Call the procedure to input name to database
		err = client.Call("API.Register", name, &reply)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("New name: %v witht ID %v added to the list", reply.Name, reply.ID)
	}
	//Call the procedure to get database
	err = client.Call("API.GetList", "GetList", &list)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("List %v:", list)

	//Try to input the name John second time.
	err = client.Call("API.Register", []byte("John"), &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("New name: %v witht ID %v added to the list", reply.Name, reply.ID)

}
