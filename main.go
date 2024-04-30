package main

import "github.com/maksymshtarkberg/Calculating-Work-Time/app"

func main() {
	workerManager := app.NewWorkerManager()
	app.HandleAddWorker(workerManager)

	workers := workerManager.GetWorkers()
	for _, worker := range workers {
		app.HandleSetWorkingHours(&worker)
		worker.TimeAtWork = &app.TimeAtWork{}
	}

	app.GetWorker(workerManager)
}
