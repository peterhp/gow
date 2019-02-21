package routine

import (
	"fmt"
	"sync"

	"github.com/peterhp/gow/crypt"
)

type cipherJob struct {
	cipher string
}

type workbench struct {
	jobChan chan cipherJob

	producers int
	consumers int

	jobs     int
	jobsChan chan int

	pwg *sync.WaitGroup
	cwg *sync.WaitGroup
}

func newWorkbench(producers, consumers, jobs int) *workbench {
	return &workbench{
		producers: producers,
		consumers: consumers,

		jobs: jobs,
	}
}

func (wb workbench) produceRoutine(id int) {
	defer wb.pwg.Done()

	for range wb.jobsChan {
		plain := crypt.RandString(10)
		cipher := crypt.Base64Encrypt(plain)
		fmt.Printf("[P%02d]: [%s] --> [%s]\n", id, plain, cipher)

		wb.jobChan <- cipherJob{cipher}
	}

	fmt.Printf("[P%02d] leaves.\n", id)
}

func (wb workbench) consumeRoutine(id int) {
	defer wb.cwg.Done()

	for job := range wb.jobChan {
		cipher := job.cipher
		plain := crypt.Base64Decrypt(cipher)
		fmt.Printf("    [C%02d]: [%s] --> [%s]\n", id, cipher, plain)
	}

	fmt.Printf("    [C%02d] leaves.\n", id)
}

func (wb *workbench) runProducers() {
	wb.pwg = &sync.WaitGroup{}

	for id := 1; id <= wb.producers; id++ {
		go wb.produceRoutine(id)
		wb.pwg.Add(1)
	}
}

func (wb workbench) waitProducers() {
	wb.pwg.Wait()
}

func (wb *workbench) runConsumers() {
	wb.cwg = &sync.WaitGroup{}

	for id := 1; id <= wb.consumers; id++ {
		go wb.consumeRoutine(id)
		wb.cwg.Add(1)
	}
}

func (wb workbench) waitConsumers() {
	wb.cwg.Wait()
}

func (wb workbench) prepareJobs() {
	for job := 0; job < wb.jobs; job++ {
		wb.jobsChan <- job
	}
}

func (wb *workbench) start() {
	wb.jobsChan = make(chan int, wb.jobs)
	wb.jobChan = make(chan cipherJob, 5)

	wb.prepareJobs()
	close(wb.jobsChan)

	wb.runProducers()
	wb.runConsumers()

	wb.waitProducers()
	close(wb.jobChan)
	wb.waitConsumers()
}

func ProduceConsume(producers, consumers, jobs int) {
	newWorkbench(producers, consumers, jobs).start()
}
