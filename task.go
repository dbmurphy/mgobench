package mgobench

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// var (
// 	session, err = mgo.Dial("127.0.0.1:27017")
// 	//session.SetMode(mgo.Monotonic, true)
// 	Col      = session.DB("test").C("test")
// 	_   Task = (*InsertTask)(nil)
// )

type TaskResult struct {
	Count     int
	TimeTaken time.Duration
}

type Task interface {
	Run() (*TaskResult, error)
	Label() string
}

type MongoTask struct {
	SM MgoManager
}

type EmptyDoc struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}

// EmptyDocInsertTask inserts empty document {}
type InsertTask struct {
	MongoTask
	Docs []interface{}
	Name string
}

func (t InsertTask) Run() (*TaskResult, error) {
	c, err := t.SM.Coll()
	// fmt.Println(t.)s
	if err != nil {
		return nil, err
	}

	st := time.Now()
	err = c.Insert(t.Docs...)
	if err != nil {
		return nil, err
	}
	r := &TaskResult{
		Count:     len(t.Docs),
		TimeTaken: time.Since(st),
	}
	return r, nil
}

func (t InsertTask) Label() string {
	return t.Name
}
