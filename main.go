package main

import "github.com/maksymshtarkberg/Calculating-Work-Time/app"

func main() {
	workerManager := app.NewWorkerManager()
	workerManager.AddWorker("John", "Doe")

	workerManager.PrintAllWorkers()

}
