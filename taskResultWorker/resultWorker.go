//
package mgobench

import (
	"fmt"
	"sync"
	"time"

	mgobench "github.com/mgobench"
)

type istats struct {
	sync.Mutex
	avgTime time.Duration
	count   int
}

func (s *istats) add(runTime time.Duration, runCount int) {
	s.Lock()
	defer s.Unlock()
	s.avgTime = s.avgTime + runTime
	s.count = s.count + runCount
}
func (s *istats) avg() float64 {
	s.Lock()
	defer s.Unlock()
	avgTime := float64(s.avgTime) / float64(s.count)
	return avgTime
}
func (s *istats) reset() {
	s.Lock()
	defer s.Unlock()
	s.avgTime = 0
	s.count = 0
}

// NewResultWorker returns a resultWorker that will emit from the C
// channel send data to influxdb only 'n' times every 'rate' seconds.
func NewResultWorker(n int, rate time.Duration) *resultWorker {
	seedTime, _ := time.ParseDuration("0ms")
	r := &resultWorker{

		C:        make(chan mgobench.TaskResult),
		rate:     rate,
		n:        n,
		shutdown: make(chan bool),
		stats: istats{
			avgTime: seedTime,
			count:   0,
		},
	}
	r.wg.Add(1)
	go r.worker()
	return r
}

type resultWorker struct {
	sync.RWMutex
	C     chan mgobench.TaskResult
	rate  time.Duration
	n     int
	stats istats

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
			sendToInflux("insertCount", "insert_time", r.stats.avg())
			sendToInflux("insertCount", "insert_time", float64(r.stats.count))
			r.stats.reset()

		case val := <-r.C:
			r.stats.add(val.TimeTaken, val.Count)
			// fmt.Println("Value pushed in channel", val.Count, "  time is ", val.TimeTaken)

		case <-r.shutdown:
			break Loop
		}
	}
}

func sendToInflux(measure string, tag string, val float64) {

	// TODO: locking
	mgobench.InsertData(measure, tag, val)
}
