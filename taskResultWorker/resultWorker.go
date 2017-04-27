//
package mgobench

import (
	"fmt"
	"sync"
	"time"
)

// NewResultWorker returns a resultWorker that will emit from the C
// channel send data to influxdb only 'n' times every 'rate' seconds.
func NewResultWorker(n int, rate time.Duration) *resultWorker {
	r := &resultWorker{

		C:        make(chan int),
		rate:     rate,
		n:        n,
		shutdown: make(chan bool),
	}
	r.wg.Add(1)
	go r.worker()
	return r
}

type resultWorker struct {
	sync.RWMutex
	C    chan int
	rate time.Duration
	n    int

	shutdown chan bool
	wg       sync.WaitGroup
}

func (r *resultWorker) Stop() {
	close(r.shutdown)
	r.wg.Wait()
	close(r.C)
}

func (r *resultWorker) worker() {
	defer r.wg.Done()
	ticker := time.NewTicker(r.rate)
	defer ticker.Stop()

Loop:
	for {
		select {

		case <-ticker.C:
			fmt.Println("second passed")
			// send to influxdb

		case val := <-r.C:
			fmt.Println("Value pushed in channel", val)

		case <-r.shutdown:
			break Loop
		}
	}
}
