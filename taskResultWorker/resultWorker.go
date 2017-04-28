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
	s.count++
}

func (s *istats) get() (float64, float64) {
	s.Lock()
	defer s.Unlock()
	if s.count == 0 {
		s.count = 1
	}
	avg := (float64(s.avgTime) / float64(s.count)) / 1000
	cou := float64(s.count)
	return avg, cou
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

		C:        make(chan mgobench.TaskResult, 100),
		rate:     rate,
		n:        n,
		shutdown: make(chan bool),
		stats: istats{
			avgTime: seedTime,
			count:   0,
		},
	}
	r.wg.Add(1)
	go r.worker(1)
	r.wg.Add(1)
	go r.worker(2)
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

func (r *resultWorker) worker(id int) {
	defer r.wg.Done()
	ticker := time.NewTicker(r.rate)
	defer ticker.Stop()
	counter := 0

Loop:
	for {
		select {

		case <-ticker.C:
			// send to influxdb
			go sendToInflux(r, id)

		case val := <-r.C:
			r.stats.add(val.TimeTaken, val.Count)
			counter++

		case <-r.shutdown:
			// send to influxdb
			sendToInflux(r, id)
			break Loop
		}
	}
	fmt.Println("total processed results by worker ", counter)
}

func sendToInflux(r *resultWorker, id int) {
	avg, cou := r.stats.get()
	// fmt.Println("on ", id, " second passed inserted %s in avg time %s", cou, avg)
	mgobench.InsertData("insertTime", "insert_time", avg)
	mgobench.InsertData("insertCount", "insert_count", cou)
	r.stats.reset()
}
