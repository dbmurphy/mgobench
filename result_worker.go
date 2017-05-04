//
package mgobench

import (
	"fmt"
	"sync"
	"time"
)

type istats struct {
	sync.RWMutex
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
	s.RLock()
	defer s.RUnlock()
	avg := 0.0
	cou := 0.0
	//fmt.Println(float64(s.avgTime))
	if s.count != 0 {
		totalTimeInMilli := float64(s.avgTime) / 1000000
		avg = totalTimeInMilli / float64(s.count)
		cou = float64(s.count)
	}

	// avg := (float64(s.avgTime) / float64(s.count)) / 1000

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
func NewResultWorker(n int, rate time.Duration, config *Config) *ResultWorker {
	seedTime, _ := time.ParseDuration("0ms")
	r := &ResultWorker{

		C:        make(chan TaskResult, 100),
		rate:     rate,
		n:        n,
		shutdown: make(chan bool),
		stats: istats{
			avgTime: seedTime,
			count:   0,
		},
	}
	influxdb := NewInfluxClient(config)
	// fmt.Println("first implementation", influxdb)

	for i := 0; i < n; i++ {
		r.wg.Add(1)
		go r.worker(i, influxdb)
	}
	return r
}

type ResultWorker struct {
	sync.RWMutex
	C     chan TaskResult
	rate  time.Duration
	n     int
	stats istats

	shutdown chan bool
	wg       sync.WaitGroup
}

func (r *ResultWorker) Stop() {
	close(r.shutdown)
	r.wg.Wait()
	close(r.C)
}

func (r *ResultWorker) worker(id int, influxdb *Influxdb) {
	defer r.wg.Done()
	ticker := time.NewTicker(r.rate)
	defer ticker.Stop()
	counter := 0

Loop:
	for {
		select {

		case <-ticker.C:
			// send to influxdb
			go sendToInflux(r, id, influxdb)

		case val := <-r.C:

			r.stats.add(val.TimeTaken, val.Count)
			counter++

		case <-r.shutdown:
			// send to influxdb
			sendToInflux(r, id, influxdb)
			break Loop
		}
	}
	fmt.Println("total processed results by worker ", counter)
}

func sendToInflux(r *ResultWorker, id int, influxdb *Influxdb) {
	defer r.stats.reset()
	avg, cou := r.stats.get()
	fmt.Println(avg, "                     ", cou)
	influxdb.InsertData("insertTime", "insert_time", avg)
	influxdb.InsertData("insertCount", "insert_count", cou)

}
