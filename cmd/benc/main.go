package main

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"

	Mbp "github.com/mgobench"
	cases "github.com/mgobench/cases"
)

var session, err = mgo.Dial("10.5.2.143:27017")

func main() {
	//session.SetMode(mgo.Strong, false)
	newSesCopy := session.Copy()
	wm := Mbp.NewWorkerManager(3)
	// fmt.Println(wm.NumWorker())
	// fmt.Println(wm.IsRunning())

	ccc := Mbp.NewCollectionBindFunc("Oorder", "test")
	mm := Mbp.MgoManager{
		Session: newSesCopy,
		CFn:     ccc,
	}

	mt := Mbp.MongoTask{
		SM: mm,
	}

	var data = make([]interface{}, 0)
	data = append(data, cases.EmptyTest())
	for i := 0; i < 100; i++ {
		it := Mbp.InsertTask{
			MongoTask: mt,
			Docs:      data,
			Name:      "Oorder",
		}
		wm.Send(it)
	}
	tr := wm.Result()
	r := Mbp.NewResultWorker(5, 2*time.Second)
	for d := range tr {
		r.C <- d
		//defer d.Session.Close()
		fmt.Println(d.Count, "time", d.TimeTaken)
	}

	wm.Stop()
	r.Stop()
}
