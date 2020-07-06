package utils

import "fmt"

type Task struct {
	taskId int
	f      func() error
}

func NewTask(id int, f func() error) *Task {
	return &Task{
		taskId: id,
		f:      f,
	}
}

func (t *Task) execute() {
	t.f()
}

type Pool struct {
	workerNum  int
	EntryChan  chan *Task
	workerChan chan *Task
}

func NewPool(num int) *Pool {
	return &Pool{
		workerNum:  num,
		EntryChan:  make(chan *Task),
		workerChan: make(chan *Task),
	}
}

func (p *Pool) worker(id int) {
	for task := range p.workerChan {
		task.execute()
		fmt.Println("workerId:", id, "taskId:", task.taskId, "is done")
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}

	for task := range p.EntryChan {
		p.workerChan <- task
	}
}
