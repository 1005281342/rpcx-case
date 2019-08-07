package main

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcx-case/example"
)

func main() {

	addr := "127.0.0.1:8972"
	// #1
	d := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	// #2
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// #3
	args := &example.Args{
		A: 10,
		B: 20,
	}

	// #4
	reply := &example.Reply{}

	// #5
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}

//package client

//func ClientCase() {
//	addr := "127.0.0.1:8972"
//	d := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
//	xclient := client.NewXClient("C", client.Failtry, client.RandomSelect, d, client.DefaultOption)
//	defer xclient.Close()
//
//	a := &example.A{10, 20}
//
//	b := &example.B{}
//	err := xclient.Call(context.Background(), "Mul", a, b)
//	if err != nil {
//		log.Fatalf("failed to call: %v", err)
//	}
//
//	log.Printf("%d * %d = %d", a.A1, a.A2, b.B)
//}
