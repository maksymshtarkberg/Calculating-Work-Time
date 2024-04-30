package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Worker struct {
	ID         int
	Name       string
	Surname    string
	TimeAtWork *TimeAtWork
}

type TimeAtWork struct {
	EnterTime int
	OutTime   int
}

type WorkerManager struct {
	workers map[int]Worker
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

func HandleAddWorker(wm *WorkerManager) {
	var name string
	var surname string

	fmt.Print("Enter employee name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter employee surname: ")
	surname, _ = reader.ReadString('\n')
	surname = strings.TrimSpace(surname)

	if name == "" || surname == "" {
		fmt.Print("\nName cannot be empty\n")
		return
	}

	wm.AddWorker(name, surname)
}

func (t *TimeAtWork) SetWorkingHours(enterTime, outTime int) {
	t.EnterTime = enterTime
	t.OutTime = outTime
}

func HandleSetWorkingHours(worker *Worker) {
	var enterTime int
	var outTime int

	fmt.Print("Enter the time the employee started work:")
	fmt.Scanln(&enterTime)
	fmt.Print("Enter the time the employee finished work:")
	fmt.Scanln(&outTime)

	worker.TimeAtWork.SetWorkingHours(enterTime, outTime)

}

func (wm *WorkerManager) GetWorkers() map[int]Worker {
	return wm.workers
}

func GetWorker(wm *WorkerManager) {
	fmt.Println("Workers:")
	for _, worker := range wm.workers {
		fmt.Printf("ID: %d, Name: %s, Surname: %s\n", worker.ID, worker.Name, worker.Surname)
		if worker.TimeAtWork != nil {
			fmt.Printf("Enter Time: %d, Out Time: %d\n", worker.TimeAtWork.EnterTime, worker.TimeAtWork.OutTime)
		}
	}
}
