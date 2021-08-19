package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/SergeiLisitsyn/gocourse/lesson07/jsonRPC/common"
)

func main() {
	var api = new(common.API)

	addy, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8082") //Resolve to Address
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", addy) //Listen on to resolved address
	if err != nil {
		log.Fatal(err)
	}

	rpc.Register(api) //Register resiver for prorduers
	for {
		conn, err := inbound.Accept() //Manage the connection for reciving call
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
