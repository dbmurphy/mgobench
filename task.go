package mgobench

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	session, err = mgo.Dial("server1.example.com,server2.example.com")
	//session.SetMode(mgo.Monotonic, true)
	Col      = session.DB("test").C("test")
	_   Task = (*InsertTask)(nil)
)

type TaskResult struct {
	Count     int32
	TimeTaken time.Duration
}

type Task interface {
	Run() (TaskResult, error)
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

func (t InsertTask) Run() (TaskResult, error) {
	c, err := t.SM.Coll()

	if err != nil {
		return nil, err
	}

	st := time.Now()
	err = c.Insert(t.Docs...)
	if err != nil {
		return nil, err
	}
	r := TaskResult{
		Count:     len(t.Docs),
		TimeTaken: time.Since(st),
	}
	return r, nil
}

func (t InsertTask) Label() string {
	return t.Name
}

type FlatT1Doc struct {
	ID    bson.ObjectId `bson:"_id,omitempty`
	StrF  string        `bson:"strf"`
	IntF  int64         `bson:"intf"`
	BoolF bool          `bson:"boolf"`
	TimeF time.Time     `bson:"timef"`
}

type FlatT1InsertTask struct {
}
