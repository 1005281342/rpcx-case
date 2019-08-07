package main

import (
	"github.com/smallnest/rpcx/server"
	"rpcx-case/example"
)

func main() {
	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", ":8972")
}

//func main() {
//	s := server.NewServer()
//	s.RegisterName("C", new(example.C), "")
//	s.Serve("tcp", "8972")
//}
