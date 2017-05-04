package mgobench

import (
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
	sync.RWMutex
	numWorker uint32
	tasks     chan Task
	result    chan TaskResult
	status    WorkerMgrStatus
	wait      *sync.WaitGroup
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
	w.Lock()
	defer w.Unlock()
	if w.status == WORKERMGR_STOPPED {
		return
	}
	close(w.tasks)
	w.wait.Wait()
	w.status = WORKERMGR_STOPPED
	return
}

// IsRunning returns if workers are running
func (w *workerManager) IsRunning() bool {
	w.RLock()
	defer w.RLock()
	return w.status == WORKERMGR_STARTED
}

func (w *workerManager) NumWorker() uint32 {
	w.RLock()
	defer w.RLock()
	return w.numWorker
}

func (w *workerManager) T() chan<- Task {
	return w.tasks
}

func worker(t <-chan Task, ch chan TaskResult, w *workerManager) {
	go func() {
		// fmt.Println("sfsdffd", len(ch))
		defer w.wait.Done()
		for c := range t {

			res, err := c.Run()
			if err == nil {
				defer res.Close()

				ch <- *res
			}

		}
	}()

}

func (w *workerManager) Send(t Task) error {
	w.tasks <- t
	return nil
}

// NewWorkerManager returns workerManager
func NewWorkerManager(n uint32, r chan TaskResult) WorkerManager {
	var wg sync.WaitGroup
	wm := &workerManager{
		numWorker: n,
		tasks:     make(chan Task, 1000000),
		result:    r,
		wait:      &wg,
	}

	wm.start()
	return wm
}
