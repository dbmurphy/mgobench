package mgobench

import (
	"fmt"
	"sync"
)

const (
	WORKERMGR_STARTED WorkerMgrStatus = iota
	WORKERMGR_STOPPED
)

type WorkerMgrStatus int

type WorkerManager interface {
	Stop()
	IsRunning() bool
	Send(t Task) error
	T() chan<- Task
	NumWorker() uint32
	Result() <-chan TaskResult
}

// BufferPool is used to recive task
type workerManager struct {
	numWorker uint32
	tasks     chan Task
	result    chan TaskResult
	status    WorkerMgrStatus
	wait      sync.WaitGroup
	shutdown  chan bool
}

func (w *workerManager) start() error {
	var i uint32
	for i = 0; i < w.numWorker; i++ {
		w.wait.Add(1)
		go worker(w.tasks, w.result, w)
	}
	w.status = WORKERMGR_STARTED
	return nil
}
func (w *workerManager) Result() <-chan TaskResult {

	return w.result
}

func (w *workerManager) Stop() {
	fmt.Println("stoped called")
	fmt.Println(WORKERMGR_STOPPED, "   current status  ", w.status)
	if w.status == WORKERMGR_STOPPED {
		return
	}
	close(w.shutdown)
	close(w.tasks)
	w.wait.Wait()
	w.status = WORKERMGR_STOPPED
	return
}

// IsRunning returns if workers are running
func (w *workerManager) IsRunning() bool {

	return w.status == WORKERMGR_STARTED
}

func (w *workerManager) NumWorker() uint32 {

	return w.numWorker
}

func (w *workerManager) T() chan<- Task {
	return w.tasks
}

func worker(t <-chan Task, ch chan TaskResult, w *workerManager) {

	defer func() {
		w.wait.Done()
		fmt.Println("########## down")
	}()

Loop:
	for {
		select {

		case val := <-t:

			res, err := val.Run()
			if err == nil {

				ch <- *res

			}

		case <-w.shutdown:
			break Loop
		}
	}

	///////////////////////////////////////////////////////////////

}

func (w workerManager) Send(t Task) error {

	w.tasks <- t
	return nil
}

// NewWorkerManager returns workerManager
func NewWorkerManager(n uint32, r chan TaskResult) WorkerManager {
	var wg sync.WaitGroup
	wm := &workerManager{
		numWorker: n,
		tasks:     make(chan Task, 100),
		result:    r,
		wait:      wg,
		shutdown:  make(chan bool),
	}

	wm.start()
	return wm
}
