package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
)

type WorkerService struct {}

type Task struct {
	ID uint64
	Start uint64
	End uint64
}

type Result struct {
	TaskID uint64
	Sum uint64
}

func (w *WorkerService) AssignTask(task *Task, result *Result) error {
	if task == nil {
		return errors.New("task is nil")
	}
	log.Printf("Received task %d: %d-%d\n", task.ID, task.Start, task.End)

	var sum uint64 = 0
	for i := task.Start; i <= task.End; i++ {
		sum += i
	}

	*result = Result{
		TaskID: task.ID,
		Sum: sum,
	}

	return nil
}



func main() {
	if len(os.Args)	!= 2 {
		log.Fatal(("Usage: gor run rpc_server.go <port>"))
	}

	port := os.Args[1]
	address := fmt.Sprintf(":%s", port)

	worker := new(WorkerService)

	err := rpc.Register(worker)
	if err != nil {
		log.Fatal("Failed to register WorkerService:", err)
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Failed to listen on port 8081:", err)
	}
	defer listener.Close()

	log.Printf("RPC server is listening on port %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Failed to accept connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

