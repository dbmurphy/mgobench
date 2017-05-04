package cases

import (
	"fmt"
	"time"

	mgobench "github.com/mgobench"

	"gopkg.in/mgo.v2/bson"
)

type EmptyDoc struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}

// EmptyTest Func used to return empty data for test
func EmptyDocTest(t time.Duration, r *mgobench.ResultWorker, wm mgobench.WorkerManager, mt mgobench.MongoTask) {

	var data = make([]interface{}, 0)
	data = append(data, &EmptyDoc{})
	ch := mgobench.InsertTask{ //Item
		MongoTask: mt,
		Docs:      data,
		Name:      "Einsert",
	}
	killTime := time.After(t)
	fmt.Println("******", ch)
Loop:
	for {
		select {

		case <-killTime:
			// send to influxdb
			break Loop
			// default:

			wm.Send(ch)
		}
	}

}
