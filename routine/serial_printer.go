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
	}
}

func (sp serialPrinter) printRoutine(id int) {
	defer sp.wg.Done()

	for num := range sp.numChan {
		if num > sp.maxNum {
			sp.numChan <- num
			break
		}

		if num%sp.routCount == id%sp.routCount {
			fmt.Printf("Routine [%d] --> %02d\n", id, num)

			num++
		}
		sp.numChan <- num
	}

	fmt.Printf("Routine [%d] leaves.\n", id)
}

func (sp *serialPrinter) start() {
	sp.numChan = make(chan int, 1)
	sp.numChan <- sp.minNum

	sp.wg = &sync.WaitGroup{}
	for id := 1; id <= sp.routCount; id++ {
		go sp.printRoutine(id)
		sp.wg.Add(1)
	}

	sp.wg.Wait()
	close(sp.numChan)
}

func SerialPrint(max, routines int) {
	newSerialPrinter(1, max, routines).start()
}
