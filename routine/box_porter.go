package routine

import (
	"fmt"
)

func workerPool(workers int) (func(func()), func()) {
	jobChan := make(chan func())
	doneChan := make(chan bool)

	for worker := 0; worker < workers; worker++ {
		go func() {
			for job := range jobChan {
				job()
			}

			doneChan <- true
		}()
	}

	return func(job func()) {
			jobChan <- job
		}, func() {
			close(jobChan)

			for worker := 0; worker < workers; worker++ {
				<-doneChan
			}
		}
}

func moveBox(box int) {
	fmt.Printf("Move box [%02d].\n", box)
}

func PortBox(boxes, porters int) {
	addJob, waitJobs := workerPool(porters)

	for box := 0; box < boxes; box++ {
		addJob(func(box int) func() {
			return func() {
				moveBox(box)
			}
		}(box))
	}
	waitJobs()
}
