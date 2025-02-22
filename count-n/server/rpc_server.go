package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	worker := new(WorkerService)

	err := rpc.Register(worker)
	if err != nil {
		log.Fatal("Failed to register WorkerService:", err)
	}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Failed to listen on port 8081:", err)
	}
	defer listener.Close()

	log.Println("RPC server is listening on port 8081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Failed to accept connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}