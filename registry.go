package mgobench

import (
	"errors"
	"sync"
	"time"
)

// NewFunc will used as type of data
// type NewFunc func(d time.Duration)
type NewFunc func(time.Duration, *ResultWorker, WorkerManager, MongoTask)
type Registry struct {
	sync.RWMutex
	list map[string]NewFunc
}

// Newregistry initialize registry
func Newregistry() *Registry {
	return &Registry{
		list: make(map[string]NewFunc),
	}
}

// Add is add func in registry
func (r *Registry) Add(name string, funcName NewFunc) error {
	r.Lock()
	defer r.Unlock()
	if r.list[name] != nil {
		return errors.New("func alredy registered")
	}
	r.list[name] = funcName
	return nil
}

// Get is used to get func from registry
func (r *Registry) Get(name string) (NewFunc, error) {
	r.Lock()
	defer r.Unlock()
	if r.list[name] == nil {
		return nil, errors.New("func not found")
	}
	return r.list[name], nil
}
