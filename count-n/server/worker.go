package main

import (
	"errors"
	"log"
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
