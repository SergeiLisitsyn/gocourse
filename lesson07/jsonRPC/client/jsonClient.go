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
	for {
		fmt.Println("If you want to add name in register type <1>.\nIf you want to resive full list of the items, type <2>. ")

		choose, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		switch procedure := string(choose); procedure {

		case "1":
			{
				fmt.Println("Please type name you want to add in list.")
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
		case "2":
			{
				//Call the procedure to get database
				err = client.Call("API.GetList", "GetList", &list)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("List :")
				var count = 1
				for _, v := range list.Database {
					fmt.Printf("%v) ID: %v Name: %v;\n", count, v.ID, v.Name)
					count++
				}
			}
		default:
			fmt.Println("You did not choose. Action 3, which doesn't exist")

		}
	}
}
