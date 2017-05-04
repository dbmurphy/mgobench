package cases

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/mgobench"
	"github.com/roshanraj/goRandString/goRand"
)

type FlatT1InsertTask struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	RandomStr string
}

func FlatT1InsertTaskTest(t time.Duration, r *mgobench.ResultWorker, wm mgobench.WorkerManager, mt mgobench.MongoTask) {
	var data = make([]interface{}, 0)
	data = append(data, &FlatT1InsertTask{
		ID:        bson.NewObjectId(),
		RandomStr: goRand.RandString(8),
	})
	ch := mgobench.InsertTask{
		MongoTask: mt,
		Docs:      data,
		Name:      "Oorder",
	}
	fmt.Println("******", ch)
	killTime := time.After(t)
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
