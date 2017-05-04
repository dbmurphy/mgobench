package cases

import (
	"time"

	mgobench "github.com/mgobench"

	"gopkg.in/mgo.v2/bson"
)

type EmptyDoc struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}

// EmptyTest Func used to return empty data for test
func EmptyDocTest(t time.Duration, r *mgobench.ResultWorker, wm mgobench.WorkerManager, mt mgobench.MongoTask) {

	killTime := time.After(t)

Loop:
	for {
		select {

		case <-killTime:
			// send to influxdb
			break Loop
		default:
			var data = make([]interface{}, 0)
			data = append(data, &EmptyDoc{})
			ch := mgobench.InsertTask{ //Item
				MongoTask: mt,
				Docs:      data,
				Name:      "Einsert",
			}
			wm.Send(ch)
		}
	}

}
