package routine

import (
	"fmt"
	"sync"
)

type serialPrinter struct {
	minNum int
	maxNum int

	routCount int

	numChan chan int
	wg      *sync.WaitGroup
}

func newSerialPrinter(min, max, count int) *serialPrinter {
	return &serialPrinter{
		minNum: min,
		maxNum: max,

		routCount: count,

		numChan: make(chan int),
		wg:      &sync.WaitGroup{},
	}
}

func (sp serialPrinter) printerRoutine(id int) {
	defer sp.wg.Done()

	for num := range sp.numChan {
		if num <= sp.maxNum && num%sp.routCount == id%sp.routCount {
			fmt.Printf("Routine [%d] --> %02d\n", id, num)

			num++
		}
		sp.numChan <- num
	}

	fmt.Printf("Routine [%d] leaves.\n", id)
}

func (sp serialPrinter) start() {
	for id := 1; id <= sp.routCount; id++ {
		go sp.printerRoutine(id)
		sp.wg.Add(1)
	}
	defer sp.wg.Wait()

	sp.numChan <- sp.minNum
	for num := range sp.numChan {
		if num > sp.maxNum {
			close(sp.numChan)
			break
		}

		sp.numChan <- num
	}
}

func SerialPrint(max, routines int) {
	newSerialPrinter(1, max, routines).start()
}
