package app

import "fmt"

type Worker struct {
	ID      int
	Name    string
	Surname string
}

type WorkerManager struct {
	workers map[int]Worker
}

type TimeAtWork struct {
	EnterTime int
	OutTime   int
}

func NewWorkerManager() *WorkerManager {
	return &WorkerManager{
		workers: make(map[int]Worker),
	}
}

func (wm *WorkerManager) AddWorker(name, surname string) {
	id := len(wm.workers) + 1
	worker := Worker{
		ID:      id,
		Name:    name,
		Surname: surname,
	}
	wm.workers[id] = worker
}

func (wm *WorkerManager) PrintAllWorkers() {
	for _, worker := range wm.workers {
		fmt.Println("Name:%s, Surname:%s", worker.Name, worker.Surname)
	}
}
